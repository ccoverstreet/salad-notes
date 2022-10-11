package pandoc

import "os/exec"

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func RunPandoc(srcFile string, destFile string, args []string) error {
	cmd := exec.Command("pandoc", args...)

	return nil
}
