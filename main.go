package main

import (
	"os"

	"github.com/ccoverstreet/salad-notes/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	setupLogging()
	log.Info().Msg("Starting Salad Notes")

	core := app.CreateSaladApp()
	core.Listen(8080)

	//pandoc.RunPandoc("test.md", "test.html", []string{"--mathjax"})
}

func setupLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().Caller().Logger()
}
