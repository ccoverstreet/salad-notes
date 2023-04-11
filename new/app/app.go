package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ccoverstreet/salad-notes/database"
	"github.com/ccoverstreet/salad-notes/pandoc"
	sutil "github.com/ccoverstreet/salad-notes/sutil"
	"github.com/go-chi/chi/v5"
)

type App struct {
	DB database.DatabaseHandle
}

func CreateApp(name string) (*App, error) {
	// Check if database file already exists
	fileData := []byte("{}")
	dbFilename := fmt.Sprintf("%s/saladbowl.json", name)
	shouldSave := false

	if _, err := os.Stat(dbFilename); err == nil {
		// File exists
		fileData, err = os.ReadFile(dbFilename)
		if err != nil {
			return nil, err
		}
	} else if errors.Is(err, os.ErrNotExist) {
		shouldSave = true
		err = os.MkdirAll(name, 0755)
		if err != nil {
			return nil, err
		}

		fileData = []byte(fmt.Sprintf(`{
			"dataDir": "%s",
			"contentMap": {}
		}`, name))
	} else {
		return nil, fmt.Errorf("FATAL: Filesystem is in an inconsistent state - %v", err)
	}

	var db database.SaladDB
	err := json.Unmarshal(fileData, &db)
	if err != nil {
		return nil, err
	}

	if shouldSave {
		log.Println(db, "Saving")
		db.Save()
	}

	return &App{&db}, nil
}

type Handler struct {
	Closure func(http.ResponseWriter, *http.Request)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Closure(w, r)
}

func HTTPHandler(closure func(http.ResponseWriter, *http.Request)) http.Handler {
	return Handler{closure}
}

func (app *App) Start() {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "SaladNotes Backend")
	})

	router.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, "<img src='/123'/>")
		})

		r.Post("/getByName", func(w http.ResponseWriter, r *http.Request) {
			query := struct {
				Name string `json:"name"`
			}{}

			err := sutil.UnmarshalJSONBody(r, &query)
			if err != nil {
				sutil.HttpHandlerError(w, err)
			}

			docs := app.DB.GetItemsByName(query.Name)

			b, err := json.Marshal(docs)
			if err != nil {
				sutil.HttpHandlerError(w, err)
			}

			w.Write(b)
		})

		r.Post("/getByTags", func(w http.ResponseWriter, r *http.Request) {
			query := struct {
				Tags []string `json:"tags"`
			}{}

			err := sutil.UnmarshalJSONBody(r, &query)
			if err != nil {
				sutil.HttpHandlerError(w, err)
			}

			docs := app.DB.GetItemsByTags(query.Tags)

			b, err := json.Marshal(docs)
			if err != nil {
				sutil.HttpHandlerError(w, err)
			}

			w.Write(b)
		})

		r.Get("/uid/{uid}", func(w http.ResponseWriter, r *http.Request) {
			uid := chi.URLParam(r, "uid")

			doc, err := app.DB.GetItemByUID(uid)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			log.Println(doc)

			b, err := app.DB.ReadFile(doc.UID)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			jsonBytes, err := json.Marshal(doc)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			w.Header().Set("Content-Type", database.DocTypeToMime(doc.FileType))
			w.Header().Set("saladnotes-info", string(jsonBytes))
			w.Write(b)
		})

		r.Get("/render/{uid}", func(w http.ResponseWriter, r *http.Request) {
			uid := chi.URLParam(r, "uid")

			doc, err := app.DB.GetItemByUID(uid)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			if doc.FileType != database.MD {
				sutil.HttpHandlerError(w, fmt.Errorf("Invalid document type for rendering"))
				return
			}

			b, err := app.DB.ReadFile(doc.UID)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			md, err := pandoc.ConvertMDToHTML(b)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			w.Header().Set("Content-Type", "text/markdown")
			w.Write(md)
		})

		r.Post("/addItem", func(w http.ResponseWriter, r *http.Request) {
			query := struct {
				Name string   `json:"name"`
				Tags []string `json:"tags"`
			}{}

			err := sutil.UnmarshalJSONBody(r, &query)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			newDoc, err := app.DB.AddItem(query.Name, database.MD, query.Tags, []byte{})
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			b, err := json.Marshal(newDoc)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			w.Header().Set("Content-Type", "text/markdown")
			w.Write(b)
		})

		r.Post("/deleteItem", func(w http.ResponseWriter, r *http.Request) {
			query := struct {
				UID string `json:"uid"`
			}{}

			err := sutil.UnmarshalJSONBody(r, &query)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			err = app.DB.DeleteItem(query.UID)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "{}")
		})

		r.Post("/writeItem", func(w http.ResponseWriter, r *http.Request) {
			uid := r.Header.Get("saladnotes-uid")
			data, err := io.ReadAll(r.Body)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			err = app.DB.WriteItem(uid, data)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}
		})

		r.Post("/updateMeta", func(w http.ResponseWriter, r *http.Request) {
			query := struct {
				UID  string   `json:"uid"`
				Name string   `json:"name"`
				Tags []string `json:"tags"`
			}{}

			err := sutil.UnmarshalJSONBody(r, &query)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			err = app.DB.UpdateItemMeta(query.UID, query.Name, query.Tags)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

			fmt.Fprintf(w, "{}")
		})

		r.Post("/renderNote", func(w http.ResponseWriter, r *http.Request) {
			query := struct {
				UID string `json:"uid"`
			}{}

			err := sutil.UnmarshalJSONBody(r, &query)
			if err != nil {
				sutil.HttpHandlerError(w, err)
				return
			}

		})
	})

	http.ListenAndServe(":9999", router)
}
