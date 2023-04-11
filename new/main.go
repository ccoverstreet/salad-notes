package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/ccoverstreet/salad-notes/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Setup logging
	//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).
	//log.Logger = log.
	//With().Caller().Logger()

	app, err := app.CreateApp("tester")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Unable to create application instance")
		return
	}

	log.Info().Msgf("Starting app on port %d", 8080)
	app.Start()
}

func HttpHandlerError(w http.ResponseWriter, err error) {
	log.Info().
		Err(err).
		Msg("Error in HTTP Handler")
}

/*
func ConvertMarkdownToHTMLHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sutil.HttpHandlerError(w, err)
		return
	}

	data := struct {
		Content string `json:"content"`
	}{}

	err = json.Unmarshal(body, &data)

	if err != nil {
		sutil.HttpHandlerError(w, err)
		return
	}

	out, err := pandoc.ConvertMDToHTML(data.Content)
	if err != nil {
		sutil.HttpHandlerError(w, err)
		return
	}

	resp := struct {
		Markdown string `json:"markdown"`
	}{string(out)}

	b, err := json.Marshal(resp)
	if err != nil {
		sutil.HttpHandlerError(w, err)
		return
	}

	w.Write(b)
}
*/
