package pandoc

import (
	"os/exec"
)

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func RunPandoc(srcFile string, args []string) ([]byte, error) {
	fullArgs := append([]string{srcFile}, args...)
	cmd := exec.Command("pandoc", fullArgs...)

	return cmd.Output()
}
