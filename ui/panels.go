package ui

import (
	lipgloss "charm.land/lipgloss/v2"
	"fmt"
	"github.com/adanrsantos/TradeTUI/types"
)

type MenuItem struct {
	Label  string
	Target string
}

var SettingMenuChoices = []MenuItem{
	{Label: "Databento", Target: "databento"},
	{Label: "Alpha", Target: "alpha"},
}

func SettingPanel(m *types.ParentModel) string {
	menu := ""

	labels := make([]string, len(SettingMenuChoices))
	for i, item := range SettingMenuChoices {
		cursor := " "
		if m.SettingCursor == i {
			cursor = ">"
		}

		labels[i] = item.Label

		menu += fmt.Sprintf("%s %s", cursor, labels[i])
	}

	return lipgloss.NewStyle().Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			menu,
		),
	)
}
