package viewAdaptors

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TableAdaptor(data [][]string, view *tview.Table) {

	view.Clear()
	keyColor := tcell.ColorBlue
	rowColor := tcell.ColorNone

	for i, row := range data {
		for j, col := range row {
			color := rowColor
			if j == 0 {
				color = keyColor
			}
			view.SetCell(i, j,
				tview.NewTableCell(col).
					SetTextColor(color).
					SetAlign(tview.AlignLeft))
		}
	}
	view.SetSelectable(true, false)
	view.SetSelectedStyle(tcell.StyleDefault.Background(keyColor))
}
