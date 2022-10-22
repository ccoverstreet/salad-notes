package assets

import (
	"embed"
	"path"
)

//go:embed index.html
//go:embed index.css
//go:embed index.js
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
}

func DetermineMimeType(filename string) string {
	ext := path.Ext(filename)

	return MIMEMAP[ext]
}
