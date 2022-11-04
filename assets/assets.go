package assets

import (
	"embed"
	"path"
)

//go:embed *
var embedFS embed.FS

type FileData struct {
	Bytes    []byte
	MimeType string
}

func GetAsset(filename string) (FileData, error) {
	data, err := embedFS.ReadFile(filename)

	return FileData{data, DetermineMimeType(filename)}, err
}

var MIMEMAP map[string]string = map[string]string{
	".js":  "text/javascript",
	".css": "text/css",
	".svg": "image/svg+xml",
	".png": "image/png",
}

func DetermineMimeType(filename string) string {
	ext := path.Ext(filename)

	return MIMEMAP[ext]
}
