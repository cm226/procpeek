package tools

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func Strace(pid int) (io.ReadCloser, io.ReadCloser, *exec.Cmd) {

	cmd := exec.Command("strace", "-p", fmt.Sprint(pid))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return stdout, stderr, cmd
}
