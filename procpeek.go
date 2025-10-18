package main

import (
	"flag"
	"procpeek/menu"
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

type Page struct {
	name     string
	ui       tview.Primitive
	shortcut rune
}

type App struct {
	app     *tview.Application
	updater *updater.ViewUpdater
	pages   []*Page
}

func buildSysCallsView(app *tview.Application, pid *int) (*Page, *exec.Cmd) {

	straceOut, cmd := tools.Strace(*pid)
	sysCalls := views.SystemCalls(app)
	viewAdaptors.CopyStream(straceOut, sysCalls)

	page := Page{name: "SysCalls", ui: sysCalls, shortcut: 'y'}
	return &page, cmd
}

func buildFDPages(app *tview.Application, pid *int, viewUpdater *updater.ViewUpdater) []*Page {

	LsofOut := func() []map[rune]string { return tools.Lsof(*pid) }
	var lsofCache = updater.MakeToolCache(LsofOut)
	viewUpdater.AddCache(&lsofCache)

	filesTable := views.Table(app, "Files")
	viewUpdater.AddView(func() { viewAdaptors.FileAdaptorAdaptor(lsofCache, filesTable) })
	socketTable := views.Table(app, "Sockets")
	viewUpdater.AddView(func() { viewAdaptors.SocketAdaptorAdaptor(lsofCache, socketTable) })

	filesPage := Page{name: "Files", ui: filesTable, shortcut: 'f'}
	socketPage := Page{name: "Socket", ui: socketTable, shortcut: 's'}
	return []*Page{&filesPage, &socketPage}
}

func initApp(app *App) {
	pages := tview.NewPages()
	menuItems := [][]string{}

	for _, page := range app.pages {
		menuItems = append(menuItems, []string{string(page.shortcut), page.name})
		pages.AddPage(page.name,
			page.ui,
			true,
			true)
	}

	app.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for _, page := range app.pages {
			if event.Rune() == page.shortcut {
				pages.SwitchToPage(page.name)
			}
		}
		return event
	})

	menuView := menu.NewMenu(menuItems)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(menuView, 1, 1, false)

	app.app.SetRoot(flex, true).SetFocus(pages)

}

func main() {

	pid := flag.Int("p", 42, "The process ID (pid) of the process to peek")
	flag.Parse()

	app := App{
		updater: updater.CreateNew(time.Millisecond * 1000),
		app:     tview.NewApplication(),
		pages:   []*Page{},
	}

	sysCalls, _ := buildSysCallsView(app.app, pid)

	app.pages = append(app.pages, sysCalls)
	app.pages = append(app.pages, buildFDPages(app.app, pid, app.updater)...)

	initApp(&app)

	app.updater.Run(app.app)
	if err := app.app.Run(); err != nil {
		panic(err)
	}
}
