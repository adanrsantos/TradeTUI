package databento

import "github.com/adanrsantos/TradeTUI/types"

var TimeFrame = []types.MenuItem{
	{Label: "1sec", Target: "1sec"},
	{Label: "15sec", Target: "15sec"},
	{Label: "1min", Target: "1min"},
}

func APIMenu(m databentoModel) string {
	style := MainStyle

	if m.focus == MainFocus {
		style = FocusStyle
	}

	test := "Hello\nWorld\nWhat is up"

	return style.Render(test)
}

func SettingButton(m databentoModel) string {
	style := MainStyle

	if m.focus == SettingFocus {
		style = FocusStyle
	}

	test := "button"

	return style.Render(test)
}
