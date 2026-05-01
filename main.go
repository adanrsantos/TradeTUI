package main

import (
	"fmt"
	"os"

	"github.com/adanrsantos/TradeTUI/types"
	"github.com/adanrsantos/TradeTUI/ui"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

var (
	mainStyle = lipgloss.NewStyle().Width(50).Height(15)
	mainMenuStyle = lipgloss.NewStyle().Width(50).Height(10).Border(lipgloss.RoundedBorder())

	sideStyle = lipgloss.NewStyle().Height(15)

	boxStyle = lipgloss.NewStyle().Width(20).Height(5).Border(lipgloss.RoundedBorder())
)

func (m model) leftHeader() string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("62")).
		Height(5).
		Bold(true).
		Render(ui.AsciiTradeTUI())
}

func (m model) leftMenu() string {
	choices := m.choices()

	menu := ""

	for i, choice := range choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		menu += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	return mainMenuStyle.Render(menu)
}

func (m model) leftPanel() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.leftHeader(),
		m.leftMenu(),
	)
}

func (m model) rightPanel() string {
	box1 := boxStyle.Render(fmt.Sprintf(
		"TimeFrame\n%s",
		m.config.TimeFrame,
	))

	box2 := boxStyle.Render(fmt.Sprintf(
		"Symbol\n%s",
		m.config.Symbol,
	))

	box3 := boxStyle.Render(fmt.Sprintf(
		"Limit\n%d",
		m.config.Limit,
	))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		box1,
		box2,
		box3,
	)
}

func (m model) configView() string {
	return fmt.Sprintf(
		"Query Config\n\nTimeFrame: %s\nSymbol: %s\nStart: %v\nEnd: %v\nLimit: %d",
		m.config.TimeFrame,
		m.config.Symbol,
		m.config.StartDate,
		m.config.EndDate,
		m.config.Limit,
	)
}

type model struct {
	screen ui.Screen
	cursor int
	config types.QueryConfig
}

func (m model) choices() []string {
	switch m.screen {
	case ui.MainMenu:
		return []string{
			"TimeFrame",
			"Symbol",
			"Start Date",
			"End Date",
			"Limit",
		}
	case ui.TimeFrameMenu:
		return []string{
			string(types.OneSecond),
			string(types.OneMinute),
			string(types.FifteenMinute),
			string(types.OneHour),
			string(types.FourHour),
			string(types.Daily),
		}
	case ui.SymbolMenu:
		return []string{
			string(types.NQ),
			string(types.ES),
		}
	}

	return nil
}

func (m model) handleEnter() model {
	switch m.screen {

	case ui.MainMenu:
		switch m.cursor {
		case 0:
			m.screen = ui.TimeFrameMenu
		case 1:
			m.screen = ui.SymbolMenu
		}
		m.cursor = 0

	case ui.TimeFrameMenu:
		options := []types.TimeFrame{
			types.OneSecond,
			types.OneMinute,
			types.FifteenMinute,
			types.OneHour,
			types.FourHour,
			types.Daily,
		}
		m.config.TimeFrame = options[m.cursor]
		m.screen = ui.MainMenu
		m.cursor = 0

	case ui.SymbolMenu:
		options := []types.Symbol{
			types.NQ,
			types.ES,
		}
		m.config.Symbol = options[m.cursor]
		m.screen = ui.MainMenu
		m.cursor = 0
	}

	return m
}

func initialModel() model {
	return model{
		screen: ui.MainMenu,
		cursor: 0,
		config: types.QueryConfig{
			TimeFrame: types.OneMinute,
			Symbol:    types.NQ,
			Limit:     100,
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices())-1 {
				m.cursor++
			}

		case "enter":
			return m.handleEnter(), nil
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	leftPanel := mainStyle.Render(m.leftPanel())
	rightPanel := sideStyle.Render(m.rightPanel())

	return tea.NewView(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			leftPanel,
			rightPanel,
		),
	)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
