package markdown

import (
	"testing"
	"fmt"
)

func TestConversion(t *testing.T) {
	markdown := "hello\n"

	b, err := MarkdownToHTML(markdown)	
	fmt.Println(string(b), err)
}
