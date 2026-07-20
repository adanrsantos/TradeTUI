package ui

import (
	"charm.land/lipgloss/v2"
	"fmt"
	"github.com/adanrsantos/TradeTUI/providers/databento/types"
	"strings"
)

func DatasetChoices(datasets []types.Dataset) []types.MenuItem {
	items := make([]types.MenuItem, len(datasets))

	for i, dataset := range datasets {
		items[i] = types.MenuItem{
			Label: dataset.Display,
		}
	}

	return items
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

func RenderMenu(items []types.MenuItem, cursor int) string {
	var s strings.Builder

	for i, item := range items {
		if i == cursor {
			s.WriteString("> ")
		} else {
			s.WriteString(" ")
		}

		s.WriteString(item.Label)

		if i < len(items)-1 {
			s.WriteByte('\n')
		}
	}

	return s.String()
}

func MainPanel(m *types.DatabentoModel) string {
	style := MainStyle

	if m.Focus == types.MainFocus {
		style = FocusStyle
	}

	switch m.Screen {
	case types.HistMenuScreen, types.HistDataset, types.HistSymbol, types.HistSchema, types.HistStart, types.HistEnd, types.HistLimit, types.LiveMenuScreen, types.LiveSymbol, types.LiveSchema:
		return style.Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				MainMenu(m),
				QueryView(m),
			),
		)
	default:
		return style.Render(MainMenu(m))
	}
}

func MainMenu(m *types.DatabentoModel) string {
	style := MainStyle

	var items []types.MenuItem
	switch m.Screen {
	case types.MainMenuScreen:
		items = Menu
	case types.HistMenuScreen:
		return style.Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				RenderMenu(HistoricalMenu, m.MainCursor),
				QuerySubmitView(m),
			),
		)
	case types.HistDataset:
		return style.Render(
			RenderMenu(
				DatasetChoices(Datasets),
				m.MainCursor,
			),
		)
	case types.LiveMenuScreen:
		items = LiveMenu
	case types.HistSymbol, types.LiveSymbol:
		return style.Render(
			RenderMenu(
				SymbolChoices(Symbols),
				m.MainCursor,
			),
		)
	case types.HistSchema, types.LiveSchema:
		return style.Render(
			RenderMenu(
				SchemaChoices(Schemas),
				m.MainCursor,
			),
		)
	}

	return style.Render(RenderMenu(items, m.MainCursor))
}

func QueryView(m *types.DatabentoModel) string {
	style := MainStyle

	info := fmt.Sprintf(
		"Dataset: %s\n"+
			"Symbol: %s\n"+
			"Schema: %s\n"+
			"Start: %s\n"+
			"End: %s\n"+
			"Limit: %d\n",
		m.Query.Dataset.Display,
		m.Query.Symbol.Display,
		m.Query.Schema.Display,
		m.Query.StartDate,
		m.Query.EndDate,
		m.Query.Limit,
	)

	return style.Render(info)
}

func QuerySubmitView(m *types.DatabentoModel) string {
	style := MainStyle

	if m.Focus == types.QuerySubmitFocus {
		style = FocusStyle
	}

	return style.Render(RenderMenu(SubmitChoices, m.SubmitCursor))
}

func SettingButton(m *types.DatabentoModel) string {
	style := MainStyle

	if m.Focus == types.SettingFocus {
		style = FocusStyle
	}

	test := "button"

	return style.Render(test)
}
