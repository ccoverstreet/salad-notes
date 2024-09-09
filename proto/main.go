package main

import (
	"log"

	"github.com/ccoverstreet/salad-notes/app"
	"github.com/ccoverstreet/salad-notes/notesbowl"
)

func main() {
	bowl := notesbowl.CreateNotesBowl("notesdir")
	files, err := bowl.ListDirectory("coding")
	if err != nil {
		panic(err)
	}

	log.Println(files)

	core := app.CreateWebCore()
	core.Start(8080)
}
