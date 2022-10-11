package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println(commandExists("pandoc"))
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
