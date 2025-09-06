package main

import (
	"flag"
	"procpeek/controllers"
	"procpeek/tools"
	"procpeek/views"

	"github.com/rivo/tview"
)

func main() {

	pid := flag.Int("p", 42, "The process ID (pid) of the process to peek")
	flag.Parse()

	straceOut, cmd := tools.Strace(*pid)

	app := tview.NewApplication()

	sysCalls := views.SystemCalls(app)
	controllers.CopyStream(straceOut, sysCalls)

	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(sysCalls, 0, 1, false), 0, 2, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}
