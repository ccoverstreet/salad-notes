package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/ccoverstreet/salad-notes/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	port := flag.Int("p", 23452, "Port")
	flag.Parse()

	// Setup logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	//With().Caller().Logger()
	//log.Logger = log.

	app, err := app.CreateApp("mynotes")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Unable to create application instance")
		return
	}

	log.Info().Msgf("Starting app on port %d", *port)
	app.Start(*port)
}
