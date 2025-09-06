package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	//if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
	//panic(err)
	//}
	pid := flag.Int("p", 42, "The process ID (pid) of the process to peek")
	flag.Parse()

	cmd := exec.Command("strace", "-p", fmt.Sprint(*pid))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()

	//_, _, cmd := tools.Strace(*pid)

	//buf := bufio.NewReader(stream)
	//errBuf := bufio.NewReader(stderr)

	//for {
	//line, error := buf.ReadString('\n')
	//if error != nil {
	//fmt.Println("error " + error.Error())
	//}

	//errLine, error := errBuf.ReadString('\n')
	//fmt.Println(errLine)
	//fmt.Println(line)
	//}
	cmd.Wait()

}
