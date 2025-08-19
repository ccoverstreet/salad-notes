package markdown

import (
	"os/exec"
)

func MarkdownToHTML(markdown []byte) ([]byte, error) {
	cmd := exec.Command("pandoc", "-f", "markdown", "-t", "html", "--mathjax")	

	pipe, err := cmd.StdinPipe()
	if err != nil {
		return []byte{}, err
	}

	pipe.Write(markdown)
	pipe.Close()

	bout, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}

	return bout, nil
}
