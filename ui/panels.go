package ui

import (
	lipgloss "charm.land/lipgloss/v2"
	"fmt"
	"github.com/adanrsantos/TradeTUI/types"
)

var SettingMenuChoices = []types.MenuItem{
	{Label: "Databento", Target: "databento"},
	{Label: "Alpha", Target: "alpha"},
}

func SettingPanel(m *types.Model) string {
	style := DefaultStyle
	menu := ""

	labels := make([]string, len(SettingMenuChoices))
	for i, item := range SettingMenuChoices {
		cursor := " "
		if m.SettingCursor == i {
			cursor = ">"
		}

		labels[i] = item.Label

		menu += fmt.Sprintf("%s %s", cursor, labels[i])

		if i < len(labels)-1 {
			menu += "\n"
		}
	}

	return style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			menu,
		),
	)
}
