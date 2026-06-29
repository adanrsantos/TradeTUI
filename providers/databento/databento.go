package databento

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/adanrsantos/TradeTUI/types"
)

type model struct {
	model databentoModel
}

func New(cfg *types.ProviderDetails) *model {
	return &model{
		model: databentoModel{
			settingCursor: 0,
			mainCursor:    0,
			focus:         MainFocus,
			cfg:           cfg,
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "/":
			m.model.cfg.APIKey = "hello"
			return m, func() tea.Msg {
				return types.SaveConfigMsg{}
			}
		case ".":
			m.model.cfg.APIKey = "bye"
			return m, func() tea.Msg {
				return types.SaveConfigMsg{}
			}
		case "tab":
			m.model.mainCursor = 0
			if m.model.focus == MainFocus {
				m.model.focus = SettingFocus
			} else {
				m.model.focus = MainFocus
			}
		case "enter":
			switch m.model.focus {
			case MainFocus:
				switch m.model.screen {
				case MainMenuScreen:
					m.model.screen = Menu[m.model.mainCursor].Target
				}
			case SettingFocus:
			}
			m.model.mainCursor = 0
		case "h", "backspace":
			switch m.model.screen {
			case HistoricalScreen, LiveScreen, TestDataScreen:
				m.model.mainCursor = 0
				m.model.screen = MainMenuScreen
			}
		case "j", "down":
			switch m.model.screen {
			case MainMenuScreen:
				if m.model.mainCursor < len(Menu)-1 {
					m.model.mainCursor++
				}
			case HistoricalScreen:
				if m.model.mainCursor < len(HistoricalMenu)-1 {
					m.model.mainCursor++
				}
			case LiveScreen:
			case TestDataScreen:
			}
		case "k", "up":
			switch m.model.screen {
			case MainMenuScreen:
				if m.model.mainCursor > 0 {
					m.model.mainCursor--
				}
			case HistoricalScreen:
				if m.model.mainCursor > 0 {
					m.model.mainCursor--
				}
			case LiveScreen:
			case TestDataScreen:
			}
		}
	}

	return m, cmd
}

func (m model) View() tea.View {
	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			MainMenu(m.model),
			SettingButton(m.model),
		),
	)
}
