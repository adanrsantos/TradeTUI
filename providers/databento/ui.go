package databento

import (
	"strings"
)

var Menu = []MenuItem{
	{Label: "Historical", Target: HistoricalScreen},
	{Label: "Live", Target: LiveScreen},
	{Label: "Test Data", Target: TestDataScreen},
}

var HistoricalMenu = []MenuItem{
	{Label: "Symbol", Target: SymbolScreen},
	{Label: "Schema", Target: SchemaScreen},
	{Label: "Start", Target: StartScreen},
	{Label: "End", Target: EndScreen},
	{Label: "Limit", Target: LimitScreen},
}

var Schema = []MenuItem{
	{Label: "ohlcv-1s"},
	{Label: "ohlcv-1m"},
	{Label: "ohlcv-1h"},
	{Label: "ohlcv-1d"},
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

func MainMenu(m model) string {
	style := MainStyle

	if m.focus == MainFocus {
		style = FocusStyle
	}

	var items []MenuItem
	switch m.screen {
	case MainMenuScreen:
		items = Menu
	case HistoricalScreen:
		items = HistoricalMenu
	}

	return style.Render(RenderMenu(items, m.mainCursor))
}

func SettingButton(m model) string {
	style := MainStyle

	if m.focus == SettingFocus {
		style = FocusStyle
	}

	test := "button"

	return style.Render(test)
}
