package databento

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/adanrsantos/TradeTUI/types"
)

type model struct {
	settingCursor int
	mainCursor    int
	screen        Screen
	focus         Focus
	cfg           *types.ProviderDetails
}

func New(cfg *types.ProviderDetails) *model {
	return &model{
		settingCursor: 0,
		mainCursor:    0,
		screen:        MainMenuScreen,
		focus:         MainFocus,
		cfg:           cfg,
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
			m.cfg.APIKey = "hello"
			return m, func() tea.Msg {
				return types.SaveConfigMsg{}
			}
		case ".":
			m.cfg.APIKey = "bye"
			return m, func() tea.Msg {
				return types.SaveConfigMsg{}
			}
		case "tab":
			m.mainCursor = 0
			if m.focus == MainFocus {
				m.focus = SettingFocus
			} else {
				m.focus = MainFocus
			}
		case "enter":
			switch m.focus {
			case MainFocus:
				switch m.screen {
				case MainMenuScreen:
					m.screen = Menu[m.mainCursor].Target
				}
			case SettingFocus:
			}
			m.mainCursor = 0
		case "h", "backspace":
			switch m.screen {
			case HistoricalScreen, LiveScreen, TestDataScreen:
				m.mainCursor = 0
				m.screen = MainMenuScreen
			}
		case "j", "down":
			switch m.screen {
			case MainMenuScreen:
				if m.mainCursor < len(Menu)-1 {
					m.mainCursor++
				}
			case HistoricalScreen:
				if m.mainCursor < len(HistoricalMenu)-1 {
					m.mainCursor++
				}
			case LiveScreen:
			case TestDataScreen:
			}
		case "k", "up":
			switch m.screen {
			case MainMenuScreen:
				if m.mainCursor > 0 {
					m.mainCursor--
				}
			case HistoricalScreen:
				if m.mainCursor > 0 {
					m.mainCursor--
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
			MainMenu(m),
			SettingButton(m),
		),
	)
}
