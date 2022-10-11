package main

import (
	"log"

	"github.com/ccoverstreet/salad-notes/dirwatch"
	"github.com/fsnotify/fsnotify"
)

func main() {
	dw, err := dirwatch.CreateDirWatcher("./")
	if err != nil {
		panic(err)
	}

	dw.SetCallback(testCallback)

	dw.Listen()

	<-make(chan struct{})
}

func testCallback(e fsnotify.Event) {
	log.Println(e)
}
