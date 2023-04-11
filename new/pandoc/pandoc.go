package pandoc

import (
	"os/exec"
)

func ConvertMDToHTML(md []byte) ([]byte, error) {
	cmd := exec.Command("pandoc", "--from=markdown", "--to=html", "--mathjax")
	writer, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	writer.Write(md)
	writer.Close()

	return cmd.Output()
}
