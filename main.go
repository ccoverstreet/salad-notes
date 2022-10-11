package main

import (
	"log"

	"github.com/ccoverstreet/salad-notes/pandoc"
)

func main() {
	pandocExists := pandoc.CommandExists("pandoc")
	if !pandocExists {
		log.Fatal("Program 'pandoc'")
	}

	//pandoc.RunPandoc("test.md", "test.html", []string{"--mathjax"})

}
