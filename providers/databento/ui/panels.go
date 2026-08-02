package ui

import (
	"charm.land/lipgloss/v2"
	"fmt"
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
		fmt.Fprintf(&s, "%s\n%s\n%s",
			LabelStyle.Render("Historical Request"),
			PaddingStyle.Render(RenderHistoricalMenu(m)),
			PaddingStyle.BorderStyle(lipgloss.NormalBorder()).BorderTop(true).Render(RenderMenu(SubmitChoices, m.SubmitCursor)),
		)
	case types.HistDataset:
		fmt.Fprintf(
			&s, "%s\n%s",
			LabelStyle.Render("Historical Request"),
			PaddingStyle.Render(RenderMenu(
				MenuItems(Datasets, func(d types.Dataset) string {
					return d.Display
				}),
				m.MainCursor,
			)),
		)
	case types.HistSymbol:
		fmt.Fprintf(
			&s, "%s\n%s",
			LabelStyle.Render("Historical Request"),
			PaddingStyle.Render(RenderMenu(
				MenuItems(FutureSymbols, func(s types.Symbol) string {
					return s.Display
				}),
				m.MainCursor,
			)),
		)
	case types.HistSchema:
		fmt.Fprintf(
			&s, "%s\n%s",
			LabelStyle.Render("Historical Request"),
			PaddingStyle.Render(RenderMenu(
				MenuItems(Schemas, func(s types.Schema) string {
					return s.Display
				}),
				m.MainCursor,
			)),
		)
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

func RenderHistoricalMenu(m *types.DatabentoModel) string {
	var s strings.Builder

	for i, item := range HistoricalMenu {
		status := Status(m, item.Field)

		if i == m.MainCursor {
			fmt.Fprintf(&s, "> %s %s", HoverStyle.Render(item.Label), status)
		} else {
			if status == "[Locked]" {
				fmt.Fprintf(&s, LockedStyle.Render(" %s %s"), item.Label, status)
			} else {
				fmt.Fprintf(&s, " %s %s", item.Label, status)
			}
		}

		if i < len(HistoricalMenu)-1 {
			s.WriteString("\n")
		}
	}

	dataset := ""
	if m.Query.Dataset != nil {
		dataset = m.Query.Dataset.Display
	}
	symbol := ""
	if m.Query.Symbol != nil {
		symbol = m.Query.Symbol.Display
	}
	schema := ""
	if m.Query.Schema != nil {
		schema = m.Query.Schema.Display
	}
	start := ""
	if m.Query.StartDate != nil {
		start = "not zero"
	}
	end := ""
	if m.Query.EndDate != nil {
		end = "not zero"
	}
	limit := ""
	if m.Query.Limit != nil {
		limit = fmt.Sprintf("%d", m.Query.Limit)
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		s.String(),
		lipgloss.NewStyle().MarginLeft(1).PaddingLeft(1).BorderStyle(lipgloss.NormalBorder()).BorderLeft(true).Render(
			fmt.Sprintf(
				"Dataset: %v\nSymbol: %v\nSchema: %v\nStart: %v\nEnd: %v\nLimit: %v",
				dataset,
				symbol,
				schema,
				start,
				end,
				limit,
			),
		),
	)
}

func MenuItems[T any](items []T, label func(T) string) []types.MenuItem {
	menu := make([]types.MenuItem, len(items))

	for i, item := range items {
		menu[i] = types.MenuItem{
			Label: label(item),
		}
	}

	return menu
}

func Status(m *types.DatabentoModel, field types.QueryField) string {
	switch field {
	case types.DatasetField:
		if m.Query.Dataset == nil {
			return "[Not Selected]"
		}
		return "[Selected]"
	case types.SymbolField:
		if m.Query.Dataset == nil {
			return "[Locked]"
		}
		if m.Query.Symbol == nil {
			return "[Not Selected]"
		}
		return "[Selected]"
	}

	return ""
}
