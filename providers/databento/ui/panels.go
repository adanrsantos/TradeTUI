package ui

import (
	"charm.land/lipgloss/v2"
	"github.com/adanrsantos/TradeTUI/providers/databento/types"
	"strings"
)

func Dashboard(m *types.DatabentoModel) string {
	var s strings.Builder
	style := DashboardStyle

	s.WriteString(Header(m.Secret))
	s.WriteString("\n\n")
	s.WriteString(MainMenu(m))
	s.WriteString("\n\n")
	s.WriteString(RecentActivity(m))
	s.WriteString("\n\n")
	s.WriteString(Keybinds())

	return style.Render(s.String())
}

func MainMenu(m *types.DatabentoModel) string {
	var s strings.Builder

	switch m.Screen {
	case types.MainMenuScreen:
		s.WriteString(PaddingStyle.Render(RenderMenu(Menu, m.MainCursor)))
	case types.HistMenuScreen:
		s.WriteString(LabelStyle.Render("Historical Request"))
		s.WriteString("\n")
		s.WriteString(PaddingStyle.Render(RenderMenu(HistoricalMenu, m.MainCursor)))
		s.WriteString("\n")
		s.WriteString(PaddingStyle.BorderStyle(lipgloss.NormalBorder()).BorderTop(true).Render(RenderMenu(SubmitChoices, m.SubmitCursor)))
	case types.HistDataset:
		items := make([]types.MenuItem, len(Datasets))

		for i, dataset := range Datasets {
			items[i] = types.MenuItem{
				Label: dataset.Display,
			}
		}

		s.WriteString(LabelStyle.Render("Historical Request"))
		s.WriteString("\n")
		s.WriteString(PaddingStyle.Render(RenderMenu(items, m.MainCursor)))
	case types.HistSymbol:
		items := make([]types.MenuItem, len(FutureSymbols))

		for i, symbol := range FutureSymbols {
			items[i] = types.MenuItem{
				Label: symbol.Display,
			}
		}

		s.WriteString(LabelStyle.Render("Historical Request"))
		s.WriteString("\n")
		s.WriteString(PaddingStyle.Render(RenderMenu(items, m.MainCursor)))
	}

	return s.String()
}

func Header(secret string) string {
	var s strings.Builder
	s.WriteString(LabelStyle.Render("CurrentProvider: "))
	s.WriteString(ValueStyle.Render("Databento"))
	s.WriteString("\n")
	s.WriteString(LabelStyle.Render("API_Key: "))
	if secret != "" {
		s.WriteString(SuccessStyle.Render("Loaded "))
	} else {
		s.WriteString(ErrorStyle.Render("Missing ensure .env contains 'DATABENTO_API_KEY='"))
	}

	return s.String()
}

func RecentActivity(m *types.DatabentoModel) string {
	var s strings.Builder

	s.WriteString(LabelStyle.Render("RecentActivity\n"))

	return s.String()
}

func Keybinds() string {
	var s strings.Builder
	s.WriteString("Move: j/k or \u2191/\u2193\tSelect: l or enter\tBack: h or backspace\nJump: tab\tQuit: q or ctrl+c\tHelp: ?")
	return s.String()
}

func RenderMenu(items []types.MenuItem, cursor int) string {
	var s strings.Builder

	for i, item := range items {
		if cursor == -1 {
			s.WriteString(" ")
			s.WriteString(item.Label)
		} else {
			if i == cursor {
				s.WriteString("> ")
				s.WriteString(HoverStyle.Render(item.Label))
			} else {
				s.WriteString(" ")
				s.WriteString(item.Label)
			}
		}

		if i < len(items)-1 {
			s.WriteString("\n")
		}
	}

	return s.String()
}

func SchemaChoices(symbols []types.Schema) []types.MenuItem {
	items := make([]types.MenuItem, len(symbols))

	for i, symbol := range symbols {
		items[i] = types.MenuItem{
			Label: symbol.Display,
		}
	}

	return items
}

func SymbolChoices(symbols []types.Symbol) []types.MenuItem {
	items := make([]types.MenuItem, len(symbols))

	for i, symbol := range symbols {
		items[i] = types.MenuItem{
			Label: symbol.Display,
		}
	}

	return items
}
