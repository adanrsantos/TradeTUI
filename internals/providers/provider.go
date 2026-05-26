package providers

import tea "charm.land/bubbletea/v2"

type Provider interface {
	Name() string
	Init() tea.Cmd
	Update(tea.Msg) (Provider, tea.Cmd)
	View(width, height int) tea.View
}
