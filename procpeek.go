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

const FILES_PAGE = "Files"
const SOCKETS_PAGE = "Sockets"
const SYSCALL_PAGE = "SysCalls"

func buildSysCallsView(app *tview.Application, pid *int) (*tview.TextView, *exec.Cmd) {

	straceOut, cmd := tools.Strace(*pid)
	sysCalls := views.SystemCalls(app)
	viewAdaptors.CopyStream(straceOut, sysCalls)
	return sysCalls, cmd
}

func buildFDPages(app *tview.Application, pid *int, viewUpdater *updater.ViewUpdater) (*tview.Table, *tview.Table) {

	LsofOut := func() []map[rune]string { return tools.Lsof(*pid) }
	var lsofCache = updater.MakeToolCache(LsofOut)
	viewUpdater.AddCache(&lsofCache)

	filesTable := views.Table(app, "Files")
	viewUpdater.AddView(func() { viewAdaptors.FileAdaptorAdaptor(lsofCache, filesTable) })
	socketTable := views.Table(app, "Sockets")
	viewUpdater.AddView(func() { viewAdaptors.SocketAdaptorAdaptor(lsofCache, socketTable) })

	return filesTable, socketTable
}

func main() {

	pid := flag.Int("p", 42, "The process ID (pid) of the process to peek")
	flag.Parse()

	var updater = updater.CreateNew(time.Millisecond * 1000)

	app := tview.NewApplication()

	sysCalls, _ := buildSysCallsView(app, pid)
	files, sockets := buildFDPages(app, pid, updater)

	pages := tview.NewPages()
	pages.AddPage(FILES_PAGE,
		files,
		true,
		true)

	pages.AddPage(SOCKETS_PAGE,
		sockets,
		true,
		true)

	pages.AddPage(SYSCALL_PAGE,
		sysCalls,
		true,
		true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'f' {
			pages.SwitchToPage(FILES_PAGE)
		} else if event.Rune() == 's' {
			pages.SwitchToPage(SOCKETS_PAGE)
		} else if event.Rune() == 'y' {
			pages.SwitchToPage(SYSCALL_PAGE)
		} else if event.Rune() == 'q' {
			app.Stop()
		}
		return event
	})

	updater.Run(app)

	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}
