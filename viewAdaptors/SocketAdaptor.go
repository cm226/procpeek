package viewAdaptors

import (
	"procpeek/tools"
	"procpeek/updater"

	"github.com/rivo/tview"
)

func SocketAdaptorAdaptor(
	data updater.ToolCache[[]map[rune]string], view *tview.Table) {

	var tData [][]string

	for _, file := range data.Data {
		if file[tools.FILE_TYPE] == "IPv4" || file[tools.FILE_TYPE] == "IPv6" {
			tData = append(tData, []string{file[tools.FILE_DESCRIPTOR], file[tools.FILE_NAME]})
		}
	}

	TableAdaptor(tData, view)
}
