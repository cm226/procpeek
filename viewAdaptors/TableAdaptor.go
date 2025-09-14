package viewAdaptors

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TableAdaptor(data [][]string, view *tview.Table) {

	view.Clear()
	color := tcell.ColorYellow

	for i, row := range data {
		for j, col := range row {
			view.SetCell(i, j,
				tview.NewTableCell(col).
					SetTextColor(color).
					SetAlign(tview.AlignCenter))
		}
	}
	view.SetSelectable(true, false)
}
