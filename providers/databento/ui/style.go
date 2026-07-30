package ui

import "charm.land/lipgloss/v2"

var (
	DashboardStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
	MainStyle      = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
	FocusStyle     = MainStyle.BorderForeground(lipgloss.Color("228"))

	LabelStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	ValueStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("46"))
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))

	HoverStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("230"))

	PaddingStyle = lipgloss.NewStyle().PaddingLeft(1)
)
