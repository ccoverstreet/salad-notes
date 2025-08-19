package main

import (
	"log"

	"github.com/ccoverstreet/Salad-Notes/core"
)

func main() {
	log.Print("Salad Notes starting...")
	core, err := core.CreateCore()
	if err != nil {
		panic(err)
	}

	core.Start()
}
