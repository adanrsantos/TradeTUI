package databento

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/adanrsantos/TradeTUI/types"
)

type model struct {
	settingCursor int
	mainCursor    int
	query         Query
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
				case HistMenuScreen:
					m.screen = HistoricalMenu[m.mainCursor].Target
				}
			case SettingFocus:
			}
			m.mainCursor = 0
		case "h", "backspace":
			switch m.screen {
			case HistMenuScreen, LiveMenuScreen:
				m.mainCursor = 0
				m.screen = MainMenuScreen
			case HistSymbol, HistSchema, HistStart, HistEnd, HistLimit, LiveSymbol, LiveSchema:
				if parent, ok := Parent[m.screen]; ok {
					m.screen = parent
				}
			}
		case "j", "down":
			switch m.screen {
			case MainMenuScreen:
				if m.mainCursor < len(Menu)-1 {
					m.mainCursor++
				}
			case HistMenuScreen:
				if m.mainCursor < len(HistoricalMenu)-1 {
					m.mainCursor++
				}
			case LiveMenuScreen:
				if m.mainCursor < len(LiveMenu)-1 {
					m.mainCursor++
				}
			case HistSymbol:
				if m.mainCursor < len(SymbolChoices)-1 {
					m.mainCursor++
				}
			case HistSchema:
				if m.mainCursor < len(SchemaChoices)-1 {
					m.mainCursor++
				}
			}
		case "k", "up":
			switch m.screen {
			case MainMenuScreen, HistMenuScreen, LiveMenuScreen, HistSymbol, HistSchema:
				if m.mainCursor > 0 {
					m.mainCursor--
				}
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
