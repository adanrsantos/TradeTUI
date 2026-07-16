package databento

import (
	"charm.land/lipgloss/v2"
	"fmt"
	"strings"
)

var Parent = map[Screen]Screen{
	HistMenuScreen: MainMenuScreen,
	LiveMenuScreen: MainMenuScreen,

	HistSymbol: HistMenuScreen,
	HistSchema: HistMenuScreen,
	HistStart:  HistMenuScreen,
	HistEnd:    HistMenuScreen,
	HistLimit:  HistMenuScreen,

	LiveSymbol: LiveMenuScreen,
	LiveSchema: LiveSchema,
}

var Menu = []MenuItem{
	{Label: "Historical", Target: HistMenuScreen},
	{Label: "Live", Target: LiveMenuScreen},
}

var HistoricalMenu = []MenuItem{
	{Label: "Symbol", Target: HistSymbol},
	{Label: "Schema", Target: HistSchema},
	{Label: "Start", Target: HistStart},
	{Label: "End", Target: HistEnd},
	{Label: "Limit", Target: HistLimit},
}

var LiveMenu = []MenuItem{
	{Label: "Symbol", Target: LiveSymbol},
	{Label: "Schema", Target: LiveSchema},
}

var SymbolChoices = []MenuItem{
	{Label: string(NQ)},
	{Label: string(ES)},
	{Label: string(YM)},
}

var SchemaChoices = []MenuItem{
	{Label: string(OneSecond)},
	{Label: string(OneMinute)},
	{Label: string(OneHour)},
	{Label: string(Daily)},
}

func RenderMenu(items []MenuItem, cursor int) string {
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

func MainPanel(m model) string {
	style := MainStyle

	if m.focus == MainFocus {
		style = FocusStyle
	}

	return style.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			MainMenu(m),
			QueryView(m),
		),
	)
}

func MainMenu(m model) string {
	style := MainStyle

	var items []MenuItem
	switch m.screen {
	case MainMenuScreen:
		items = Menu
	case HistMenuScreen:
		items = HistoricalMenu
	case LiveMenuScreen:
		items = LiveMenu
	case HistSymbol, LiveSymbol:
		items = SymbolChoices
	case HistSchema, LiveSchema:
		items = SchemaChoices
	}

	return style.Render(RenderMenu(items, m.mainCursor))
}

func QueryView(m model) string {
	style := MainStyle

	info := fmt.Sprintf(
		"Symbol: %s\n"+
			"Schema: %s\n"+
			"Start: %s\n"+
			"End: %s\n"+
			"Limit: %d\n",
		m.query.Symbol,
		m.query.Schema,
		m.query.StartDate,
		m.query.EndDate,
		m.query.Limit,
	)

	return style.Render(info)
}

func SettingButton(m model) string {
	style := MainStyle

	if m.focus == SettingFocus {
		style = FocusStyle
	}

	test := "button"

	return style.Render(test)
}
