package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}
