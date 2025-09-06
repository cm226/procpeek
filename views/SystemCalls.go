package views

import "github.com/rivo/tview"

func SystemCalls(app *tview.Application) *tview.TextView {
	view := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		ScrollToEnd().
		SetWrap(false).
		SetChangedFunc(func() { app.Draw() })

	view.SetBorder(true).SetTitle("System Calls")

	return view

}
