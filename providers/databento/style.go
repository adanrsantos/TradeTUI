package databento

import "charm.land/lipgloss/v2"

var (
	MainStyle  = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
	FocusStyle = MainStyle.BorderForeground(lipgloss.Color("228"))
)

