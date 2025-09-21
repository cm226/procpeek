package main

import (
	"flag"
	"procpeek/tools"
	"procpeek/updater"
	"procpeek/viewAdaptors"
	"procpeek/views"
	"time"

	"os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func buildSysCallsView(app *tview.Application, pid *int) (*tview.TextView, *exec.Cmd) {

	straceOut, cmd := tools.Strace(*pid)
	sysCalls := views.SystemCalls(app)
	viewAdaptors.CopyStream(straceOut, sysCalls)
	return sysCalls, cmd
}

func buildFDPages(app *tview.Application, pid *int, viewUpdater *updater.ViewUpdater) *tview.Pages {

	LsofOut := func() []map[rune]string { return tools.Lsof(*pid) }
	var lsofCache = updater.MakeToolCache(LsofOut)
	viewUpdater.AddCache(&lsofCache)

	filesTable := views.Table(app, "Files")
	viewUpdater.AddView(func() { viewAdaptors.FileAdaptorAdaptor(lsofCache, filesTable) })

	pages := tview.NewPages()
	pages.AddPage("Files",
		filesTable,
		true,
		true)

	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'f' {
			pages.SwitchToPage("Files")
		} else if event.Rune() == 's' {
			pages.SwitchToPage("Sockets")
		}
		return event
	})
	return pages
}

func main() {

	pid := flag.Int("p", 42, "The process ID (pid) of the process to peek")
	flag.Parse()

	var updater = updater.CreateNew(time.Millisecond * 1000)

	app := tview.NewApplication()

	sysCalls, cmd := buildSysCallsView(app, pid)
	fdsTable := buildFDPages(app, pid, updater)

	updater.Run(app)
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
