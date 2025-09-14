package views

import "github.com/rivo/tview"

func Table(app *tview.Application, name string) *tview.Table {
	view := tview.NewTable()
	view.SetBorder(true).SetTitle(name)

	return view

}
