package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/rs/zerolog/log"
)

func HttpHandlerError(w http.ResponseWriter, err error) {
	_, file, line, _ := runtime.Caller(1)
	out := fmt.Sprintf("%s:%d", file, line)
	log.Error().
		Err(err).
		Str("caller", out).
		Msg("Error in HTTP Handler")

	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "%v:%s:%d", err, file, line)
}

func UnmarshalJSONBody(r *http.Request, out any) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, out)
}
