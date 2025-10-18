package menu

import (
	"fmt"
	"log"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewMenu(menuItems [][]string) *tview.TextView {
	menu := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetTextAlign(tview.AlignLeft)

	formattedItems := []string{}
	for _, item := range menuItems {
		formattedItems = append(formattedItems, makeMenuItem(item))
	}
	_, err := fmt.Fprintf(menu, "%s", strings.Join(formattedItems, " "))

	if err != nil {
		log.Printf("failed to create new menu: %s", err.Error())
	}

	return menu
}

func makeMenuItem(items []string) string {
	key := fmt.Sprintf("[%s::b] <%s>[-:-:-]", tcell.ColorFloralWhite, items[0])
	desc := fmt.Sprintf("[%s:%s:b] %s [-:-:-]",
		tcell.ColorFloralWhite,
		tcell.ColorMediumPurple,
		strings.ToUpper(items[1]))

	return key + desc
}
