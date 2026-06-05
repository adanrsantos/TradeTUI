package ui

import (
	lipgloss "charm.land/lipgloss/v2"
)

var (
	LeftPanelStyle  = lipgloss.NewStyle().Height(16)
	RightPanelStyle = lipgloss.NewStyle().Width(25).Height(5)

	MainHeaderStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("45")).Height(5).Bold(true)
	PanelHeaderStyle = lipgloss.NewStyle().Bold(true)

	MainMenuStyle = lipgloss.NewStyle().Width(50).Height(11).Border(lipgloss.RoundedBorder())
	BoxStyle      = lipgloss.NewStyle().Width(25).Border(lipgloss.RoundedBorder())
)
