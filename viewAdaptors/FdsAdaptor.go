package viewAdaptors

import (
	"procpeek/tools"
	"sort"
	"strconv"

	"github.com/rivo/tview"
)

func FdsAdaptor(data tools.FDs, view *tview.Table) {

	var tData [][]string

	sort.Slice(data.Files, func(i, j int) bool {
		f, _ := strconv.Atoi(data.Files[i].FDName)
		s, _ := strconv.Atoi(data.Files[j].FDName)
		return f < s
	})

	for _, file := range data.Files {
		tData = append(tData, []string{file.FDName, file.Name})
	}

	TableAdaptor(tData, view)
}
