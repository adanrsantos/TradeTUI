package app

import (
	"github.com/adanrsantos/TradeTUI/internals/providers"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	width int
	height int

	Provider providers.Provider
}

func (m Model) Init() tea.Cmd {
	return m.Provider.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, cmd
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.Provider, cmd = m.Provider.Update(msg)

	return m, cmd
}

func (m Model) View() tea.View {
	return m.Provider.View(m.width, m.height)
}
