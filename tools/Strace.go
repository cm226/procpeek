package tools

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func Strace(pid int) (io.ReadCloser, *exec.Cmd) {

	cmd := exec.Command("strace", "-p", fmt.Sprint(pid))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stderr = cmd.Stdout // strace uses strErr for printing trace info

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return stdout, cmd
}
