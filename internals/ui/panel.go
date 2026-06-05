package ui

import (
	"fmt"

	"github.com/adanrsantos/TradeTUI/internals/types"

	lipgloss "charm.land/lipgloss/v2"
)

var SettingMainMenuChoices = []types.MenuItem{
	{Label: "Provider", Target: "ProviderView"},
}

func LeftPanel(m *types.Model) string {
	style := LeftPanelStyle

	return style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			MainHeader(m),
			SettingMainMenu(m),
		),
	)
}

func RightPanel(m *types.Model) string {
	style := RightPanelStyle

	return style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			SettingPanel(m),
			SubmitPanel(m),
		),
	)
}

func MainHeader(m *types.Model) string {
	style := PanelHeaderStyle

	return style.Render(AsciiTradeTUI())
}

func SettingMainMenu(m *types.Model) string {
	style := MainMenuStyle
	header := PanelHeaderStyle

	if m.Focus == types.MainMenuFocus {
		style = style.BorderForeground(lipgloss.Color("205"))
	}

	menu := ""
	labels := make([]string, len(SettingMainMenuChoices))
	for i, item := range SettingMainMenuChoices {
		cursor := " "
		if m.MainMenuCursor == i && m.Focus == types.MainMenuFocus {
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
			header.Render("Setting Menu"),
			lipgloss.NewStyle().Height(6).Render(menu),
		),
	)
}

func SettingPanel(m *types.Model) string {
	style := BoxStyle

	if m.Focus == types.SettingFocus {
		style = style.BorderForeground(lipgloss.Color("205"))
	}

	menu := ""
	labels := make([]string, len(SettingMainMenuChoices))
	for i, item := range SettingMainMenuChoices {
		cursor := " "
		if m.SettingCursor == i && m.Focus == types.SettingFocus {
			cursor = ">"
		}
		labels[i] = item.Label

		menu += fmt.Sprintf("%s %s", cursor, labels[i])

		if i < len(labels)-1 {
			menu += "\n"
		}
	}

	return style.Render(menu)
}

func SubmitPanel(m *types.Model) string {
	style := BoxStyle

	if m.Focus == types.SubmitFocus {
		style = style.BorderForeground(lipgloss.Color("205"))
	}

	menu := ""
	choices := SubmitChoices()

	for i, choice := range choices {
		cursor := " "
		if m.SubmitCursor == i && m.Focus == types.SubmitFocus {
			cursor = ">"
		}

		menu += fmt.Sprintf("%s %s", cursor, choice)

		if i < len(choices)-1 {
			menu += "\n"
		}
	}

	if m.Err != "" {
		menu += "\n"
		return style.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				menu,
				lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Render(m.Err),
			),
		)
	}

	return style.Render(menu)
}

func SubmitChoices() []string {
	return []string{"Submit", "Reset"}
}
