package main

import (
	"fmt"
	"os"

	"github.com/ccoverstreet/salad-notes/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	setupLogging()
	log.Info().Msg("Starting Salad Notes")

	core := app.CreateSaladApp(".")
	port := 33322
	log.Info().
		Str("URL", fmt.Sprintf("http://localhost:%d", port)).
		Msg("Started Salad Notes core")
	core.Listen(33322)

	//pandoc.RunPandoc("test.md", "test.html", []string{"--mathjax"})
}

func setupLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().Caller().Logger()
}
