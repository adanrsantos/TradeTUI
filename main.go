package main

import (
	"fmt"
	"os"

	"github.com/adanrsantos/TradeTUI/providers/alpha"
	"github.com/adanrsantos/TradeTUI/providers/databento"
	"github.com/adanrsantos/TradeTUI/types"
	"github.com/adanrsantos/TradeTUI/ui"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	types.Model
	ActiveProvider     tea.Model
	AvailableProviders map[string]tea.Model
}

func New() model {
	databentoModel := databento.New()
	alphaModel := alpha.New()

	return model{
		Model: types.Model{
			Screen: types.SettingScreen,
		},
		ActiveProvider: databentoModel,
		AvailableProviders: map[string]tea.Model{
			"databento": databentoModel,
			"alpha":     alphaModel,
		},
	}
}

func (m model) Init() tea.Cmd {
	if m.ActiveProvider != nil {
		return m.ActiveProvider.Init()
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.Screen == types.ProviderScreen {
				m.Screen = types.SettingScreen
			}
		case "enter":
			if m.Screen == types.SettingScreen {
				selected := ui.SettingMenuChoices[m.SettingCursor].Target

				if provider, ok := m.AvailableProviders[selected]; ok {
					m.ActiveProvider = provider

					m.Screen = types.ProviderScreen

					return m, m.ActiveProvider.Init()
				}
			}
		case "j", "down":
			if m.Screen == types.SettingScreen && m.SettingCursor < len(ui.SettingMenuChoices)-1 {
				m.SettingCursor++
			}
		case "k", "up":
			if m.Screen == types.SettingScreen && m.SettingCursor > 0 {
				m.SettingCursor--
			}
		}
	}

	if m.ActiveProvider != nil {
		var cmd tea.Cmd
		var updated tea.Model

		updated, cmd = m.ActiveProvider.Update(msg)
		m.ActiveProvider = updated

		return m, cmd
	}

	return m, cmd
}

func (m model) View() tea.View {
	if m.ActiveProvider == nil {
		return tea.NewView("Loading providers...")
	}
	if m.Screen == types.SettingScreen {
		return tea.NewView(ui.SettingPanel(&m.Model))
	}

	return m.ActiveProvider.View()
}

func main() {
	model := New()

	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
