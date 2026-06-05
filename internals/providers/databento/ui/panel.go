package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/adanrsantos/TradeTUI/internals/providers/databento/types"

	"charm.land/bubbles/v2/progress"
	lipgloss "charm.land/lipgloss/v2"
)

var MainMenuChoices = []types.MenuItem{
	{Label: "TimeFrame", Target: TimeFrameMenuView},
	{Label: "Symbol", Target: SymbolMenuView},
	{Label: "Start Date", Target: StartDateMenuView},
	{Label: "End Date", Target: EndDateMenuView},
	{Label: "Limit", Target: LimitMenuView},
}

var TimeFrameChoices = []types.TimeFrame{
	types.OneSecond,
	types.OneMinute,
	types.FifteenMinute,
	types.OneHour,
	types.FourHour,
	types.Daily,
}

var SymbolChoices = []types.Symbol{
	types.NQ,
	types.ES,
}

var DatePresets = []types.DatePreset{
	{Label: "Today", Value: func() time.Time { return time.Now() }},
	{Label: "Yesterday", Value: func() time.Time { return time.Now().AddDate(0, 0, -1) }},
	{Label: "1 Week Ago", Value: func() time.Time { return time.Now().AddDate(0, 0, -7) }},
	{Label: "1 Month Ago", Value: func() time.Time { return time.Now().AddDate(0, -1, 0) }},
	{Label: "6 Months Ago", Value: func() time.Time { return time.Now().AddDate(0, -6, 0) }},
	{Label: "1 Year Ago", Value: func() time.Time { return time.Now().AddDate(-1, 0, 0) }},
}

var LimitChoices = []types.Limit{
	types.Zero,
	types.Ten,
	types.OneHundred,
	types.OneThousand,
}

func MainMenuPanel(m *types.Model, width int) string {
	availableWidth := width - 2

	style := MainMenuStyle.Width(availableWidth)
	header := PanelHeaderStyle

	if m.Focus == MainMenuFocus {
		style = style.BorderForeground(lipgloss.Color("205"))
	}

	menu := ""
	choices := MenuChoices(m)

	for i, choice := range choices {
		cursor := " "
		if m.MainMenuCursor == i && m.Focus == MainMenuFocus {
			cursor = ">"
		}

		menu += fmt.Sprintf("%s %s", cursor, choice)

		if i < len(choices)-1 {
			menu += "\n"
		}
	}

	return style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			header.Render("DataBento API Helper"),
			lipgloss.NewStyle().Height(6).Render(menu),
			ProgressBar(m),
		),
	)
}

func MenuChoices(m *types.Model) []string {
	switch m.Screen {
	case MainMenuView:
		labels := make([]string, len(MainMenuChoices))
		for i, item := range MainMenuChoices {
			labels[i] = item.Label
		}
		return labels
	case TimeFrameMenuView:
		choices := make([]string, len(TimeFrameChoices))
		for i, tf := range TimeFrameChoices {
			choices[i] = string(tf)
		}
		return choices
	case SymbolMenuView:
		choices := make([]string, len(SymbolChoices))
		for i, sym := range SymbolChoices {
			choices[i] = string(sym)
		}
		return choices
	case StartDateMenuView:
		choices := make([]string, len(DatePresets))
		for i, t := range DatePresets {
			choices[i] = t.Label
		}
		return choices
	case EndDateMenuView:
		choices := make([]string, len(DatePresets))
		for i, t := range DatePresets {
			choices[i] = t.Label
		}
		return choices
	case LimitMenuView:
		choices := make([]string, len(LimitChoices))
		for i, l := range LimitChoices {
			choices[i] = fmt.Sprintf("%d", l)
		}
		return choices
	}

	return nil
}

func ProgressBar(m *types.Model) string {
	var completed float64

	if m.Config.TimeFrame != "" {
		completed++
	}
	if m.Config.Symbol != "" {
		completed++
	}
	if !m.Config.StartDate.IsZero() {
		completed++
	}
	if !m.Config.EndDate.IsZero() {
		completed++
	}
	if m.Config.Limit >= 0 {
		completed++
	}

	percent := completed / 5.0

	style := ProgressBarStyle

	prog := progress.New(
		progress.WithScaled(true),
		progress.WithColors(lipgloss.Color("12"), lipgloss.Color("11")),
	)

	return style.Render(prog.ViewAs(percent))
}

func HistoryPanel(m *types.Model) string {
	style := BoxStyle
	header := PanelHeaderStyle

	var history []string

	if len(m.History) == 0 {
		history = append(history, lipgloss.NewStyle().Faint(true).Render("History is empty..."))
	} else {
		for i := len(m.History) - 1; i >= 0; i-- {
			item := m.History[i]

			line := fmt.Sprintf("%s",
				item.Timestamp.Format("01/02/2006 15:04"),
			)

			history = append(history, "- "+line)

			if len(history) >= 5 {
				break
			}
		}
	}

	return style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			header.Render("History"),
			strings.Join(history, "\n"),
		),
	)
}

func QueryViewPanel(m *types.Model) string {
	style := BoxStyle
	header := PanelHeaderStyle

	start := ""
	end := ""
	limit := ""

	if !m.Config.StartDate.IsZero() {
		start = m.Config.StartDate.Format("2006-01-02")
	}

	if !m.Config.EndDate.IsZero() {
		end = m.Config.EndDate.Format("2006-01-02")
	}

	if m.Config.Limit >= 0 {
		limit = fmt.Sprintf("%d", m.Config.Limit)
	}

	info := lipgloss.NewStyle().PaddingLeft(2).Render(
		fmt.Sprintf(
			"TimeFrame: %s\n"+
				"Symbol: %s\n"+
				"Start: %s\n"+
				"End: %s\n"+
				"Limit: %s",
			m.Config.TimeFrame,
			m.Config.Symbol,
			start,
			end,
			limit,
		),
	)

	return style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			header.Render("API Request"),
			info,
		),
	)
}
