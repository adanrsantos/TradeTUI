package databento

import (
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
	case HistMenuScreen:
		items = HistoricalMenu
	case LiveMenuScreen:
		items = LiveMenu
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
