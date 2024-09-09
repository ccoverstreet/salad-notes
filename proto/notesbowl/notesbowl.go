package notesbowl

import (
	"fmt"
	"os"
)

type NotesBowl struct {
	notesDir string
}

type FileItem struct {
	Name  string `json:"name"`
	Dir   string `json:"dir"`
	IsDir bool   `json:"isDir"`
}

func CreateNotesBowl(directory string) *NotesBowl {
	return &NotesBowl{directory}
}

func (bowl *NotesBowl) ListDirectory(dirname string) ([]FileItem, error) {
	entries, err := os.ReadDir(fmt.Sprintf("%s/%s", bowl.notesDir, dirname))
	if err != nil {
		return []FileItem{}, err
	}

	items := make([]FileItem, len(entries))

	for i := 0; i < len(items); i++ {
		items[i].Name = entries[i].Name()
		items[i].Dir = dirname
		items[i].IsDir = entries[i].IsDir()
	}

	return items, nil
}

func (bowl *NotesBowl) GetFile(filename string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("%s/%s", bowl.notesDir, filename))
}
