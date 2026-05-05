package ui

import (
	lipgloss "charm.land/lipgloss/v2"
)

var (
	LeftSideStyle = lipgloss.NewStyle().Width(50).Height(16)
	RightSideStyle = lipgloss.NewStyle().Width(25).Height(5)

	MainMenuStyle = lipgloss.NewStyle().Width(50).Height(11).Border(lipgloss.RoundedBorder())
	MainHeaderStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("45")).
		Height(5).
		Bold(true)

	PanelHeaderStyle = lipgloss.NewStyle().Bold(true)

	BoxStyle = lipgloss.NewStyle().Width(25).Border(lipgloss.RoundedBorder())

	ProgressBarStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).MarginTop(1)
)