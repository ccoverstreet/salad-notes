package util

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func SendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {
		log.Error().
			Err(err).
			Interface("data", data).
			Msg("Unable to send JSON response")

		w.WriteHeader(500)

		return
	}

	w.Write(b)
}

type HTTPError struct {
	ErrStr string `json:"err"`
}

func HandleHTTPErrorAndLog(w http.ResponseWriter, statusCode int, err error) {
	log.Error().
		Err(err).
		CallerSkipFrame(1).
		Int("statusCode", statusCode).
		Msg("Error handling HTTP request")

	w.WriteHeader(statusCode)
	SendJSONResponse(w, HTTPError{err.Error()})
}
