package main

import (
	"flag"
	"procpeek/tools"
	"procpeek/updater"
	"procpeek/viewAdaptors"
	"procpeek/views"
	"time"

	"github.com/rivo/tview"
)

func main() {

	pid := flag.Int("p", 42, "The process ID (pid) of the process to peek")
	flag.Parse()

	var updater = updater.CreateNew(time.Millisecond * 1000)

	straceOut, cmd := tools.Strace(*pid)

	app := tview.NewApplication()

	sysCalls := views.SystemCalls(app)
	viewAdaptors.CopyStream(straceOut, sysCalls)

	fdsTable := views.Table(app, "Open files")
	fds := tools.Fd(*pid)
	updater.AddView(func() { viewAdaptors.FdsAdaptor(fds(), fdsTable) })

	updater.Run()
	flex := tview.NewFlex().
		AddItem(fdsTable, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(sysCalls, 0, 1, false), 0, 2, false)

	if err := app.SetRoot(flex, true).SetFocus(fdsTable).Run(); err != nil {
		panic(err)
	}

	if err := cmd.Cancel(); err != nil {
		panic(err)
	}
}
