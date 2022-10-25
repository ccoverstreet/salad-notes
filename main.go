package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ccoverstreet/salad-notes/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	flagPortHTTP := flag.Int("port", 33322, "HTTP port using by clients")
	flag.Parse()

	setupLogging()
	log.Info().Msg("Starting Salad Notes")

	portHTTP := *flagPortHTTP

	core := app.CreateSaladApp(".")
	log.Info().
		Str("URL", fmt.Sprintf("http://localhost:%d", portHTTP)).
		Msg("Started Salad Notes core")
	core.Listen(portHTTP)

	//pandoc.RunPandoc("test.md", "test.html", []string{"--mathjax"})
}

func setupLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().Caller().Logger()
}
