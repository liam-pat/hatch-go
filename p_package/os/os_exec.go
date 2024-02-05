package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, _ := exec.LookPath("ls")

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	_ = syscall.Exec(binary, args, env)
}
