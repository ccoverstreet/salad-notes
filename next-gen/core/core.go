package core

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ccoverstreet/Salad-Notes/notestorage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Core struct {
	DB     *notestorage.NotesDatabase
	router *chi.Mux
}

func CreateCore() (*Core, error) {
	core := new(Core)

	db, err := notestorage.CreateNotesDatabase("storage.db")
	if err != nil {
		panic(err)
	}

	core.DB = db
	core.router = CreateRouter(core)

	return core, nil
}

func CreateRouter(core *Core) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/allItems", WrapRoute(GetAllItemsHandler, core))
		r.Post("/getItemsByName", WrapRoute(GetItemsByNameHandler, core))
		r.Get("/getItemContent/{itemId}", WrapRoute(GetItemContentHandler, core))
		r.Get("/allTags", WrapRoute(GetAllTagsHandler, core))
		r.Post("/createItem", WrapRoute(CreateItemHandler, core))
		r.Post("/updateItemContent", WrapRoute(UpdateItemContentHandler, core))
		r.Get("/markdownAsHTML/{itemId}", WrapRoute(MarkdownToHTMLHandler, core))
		r.Post("/uploadTest", func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseMultipartForm(0)
			fmt.Println(err)

			fmt.Println(r.FormValue, r.PostForm, r.Form, r.MultipartForm)
			fileHeader := r.MultipartForm.File["binarycontent"][0]
			fmt.Println(fileHeader)
			file, err := fileHeader.Open()
			fmt.Println(err)
			b, err := io.ReadAll(file)
			fmt.Println(err)

			fmt.Println(file)

			os.WriteFile("test.png", b, 0666)

			//fmt.Println(fmt.Sprintf("%s", b))

			//fmt.Println(b)
			//fmt.Println(r.FormValue("binarycontent"))
		})
	})

	return r
}

func (core *Core) Start() {
	http.ListenAndServe(":8080", core.router)
}
