package app

import (
	"github.com/adanrsantos/TradeTUI/internals/providers"
	"github.com/adanrsantos/TradeTUI/internals/types"
	"github.com/adanrsantos/TradeTUI/internals/ui"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type Model struct {
	width  int
	height int

	types.Model

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
		case "j", "down":
			if m.Focus == types.MainMenuFocus && m.MainMenuCursor < len(ui.SettingMainMenuChoices)-1 {
				m.MainMenuCursor++
			}
			if m.Focus == types.SettingFocus && m.MainMenuCursor < len(ui.SettingMainMenuChoices)-1 {
				m.MainMenuCursor++
			}
			if m.Focus == types.SubmitFocus && m.SubmitCursor < len(ui.SubmitChoices())-1 {
				m.SubmitCursor++
			}
		case "k", "up":
			if m.Focus == types.MainMenuFocus && m.MainMenuCursor > 0 {
				m.MainMenuCursor--
			}
			if m.Focus == types.SettingFocus && m.MainMenuCursor > 0 {
				m.SettingCursor--
			}
			if m.Focus == types.SubmitFocus && m.SubmitCursor > 0 {
				m.SubmitCursor--
			}
		case "tab":
			m.MainMenuCursor = 0
			m.SubmitCursor = 0
			m.SettingCursor = 0
			switch m.Focus {
			case types.MainMenuFocus:
				m.Focus = types.SettingFocus
			case types.SettingFocus:
				m.Focus = types.SubmitFocus
			case types.SubmitFocus:
				m.Focus = types.MainMenuFocus
			}
		}
	}

	m.Provider, cmd = m.Provider.Update(msg)

	return m, cmd
}

func (m Model) View() tea.View {
	leftPanelStyle := ui.LeftPanelStyle
	rightPanelStyle := ui.RightPanelStyle

	return tea.NewView(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			leftPanelStyle.Render(ui.LeftPanel(&m.Model)),
			rightPanelStyle.Render(ui.RightPanel(&m.Model)),
		),
	)
}
