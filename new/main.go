package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ccoverstreet/salad-notes/note"
)

func main() {
	fmt.Println("vim-go")
	notes := map[string]note.Note{}

	for i := 0; i < 5; i += 1 {
		notes[strconv.Itoa(i)] =
			note.Note{strconv.Itoa(i), fmt.Sprintf("# test %d", i), []string{"test"}}
	}

	for _, v := range notes {
		if strings.Contains(v.Content, "2") {
			fmt.Println(v)
		}
	}
	fmt.Println(notes)
}
