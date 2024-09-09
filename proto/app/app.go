package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ccoverstreet/salad-notes/notesbowl"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type WebCore struct {
	router *chi.Mux
	bowl   *notesbowl.NotesBowl
}

func CreateWebCore() *WebCore {
	app := &WebCore{}

	app.router = CreateRouter(app)
	app.bowl = notesbowl.CreateNotesBowl("notesdir")

	return app
}

func CreateRouter(core *WebCore) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", rootHandler)
	r.Route("/api", func(r chi.Router) {
		r.Get("/", rootHandler)
		r.Post("/listDirectory", wrapRoute(listDirectoryHandler, core))
		r.Post("/getFile", wrapRoute(getFileHandler, core))
	})

	return r
}

func (core *WebCore) Start(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), core.router)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Salad Notes")
}

func wrapRoute(route func(http.ResponseWriter, *http.Request, *WebCore), core *WebCore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		route(w, r, core)
	}
}

func ReadRequest(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {

	}

	err = json.Unmarshal(body, &v)

	return err
}

func listDirectoryHandler(w http.ResponseWriter, r *http.Request, core *WebCore) {
	v := struct {
		Directory string `json:"directory"`
	}{}
	err := ReadRequest(r, &v)
	log.Printf("%v", err)

	items, err := core.bowl.ListDirectory(v.Directory)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	b, err := json.Marshal(items)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	w.Write(b)
}

func getFileHandler(w http.ResponseWriter, r *http.Request, core *WebCore) {
	v := struct {
		Filename string `json:"filename"`
	}{}
	err := ReadRequest(r, &v)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	b, err := core.bowl.GetFile(v.Filename)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	w.Write(b)
}
