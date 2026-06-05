package ui

import (
	lipgloss "charm.land/lipgloss/v2"
)

var (
	MainMenuStyle = lipgloss.NewStyle().Width(50).Height(11).Border(lipgloss.RoundedBorder())
	BoxStyle      = lipgloss.NewStyle().Width(25).Border(lipgloss.RoundedBorder())

	MainHeaderStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("45")).Height(5).Bold(true)
	PanelHeaderStyle = lipgloss.NewStyle().Bold(true)

	ProgressBarStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).MarginTop(1)
)
