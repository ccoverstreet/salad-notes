package util

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

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

func UnmarshalJSONBody(r *http.Request, out interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &out)
	return err
}

type FileInfo struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
}

// Returns a JSON friendly array of fileInfo
// for the target directory
func ListDir(dirName string) ([]FileInfo, error) {
	files, err := os.ReadDir(dirName)
	if err != nil {
		return []FileInfo{}, err
	}

	listOfFileInfo := make([]FileInfo, len(files))

	for i, v := range files {
		listOfFileInfo[i] = FileInfo{v.Name(), v.IsDir()}
	}

	return listOfFileInfo, nil
}
