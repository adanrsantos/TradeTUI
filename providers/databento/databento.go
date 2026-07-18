package databento

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"fmt"
	"github.com/adanrsantos/TradeTUI/providers/databento/api"
	"github.com/adanrsantos/TradeTUI/providers/databento/types"
	"github.com/adanrsantos/TradeTUI/providers/databento/ui"
	globalTypes "github.com/adanrsantos/TradeTUI/types"
)

type Model struct {
	*types.DatabentoModel
}

func New(cfg *globalTypes.ProviderDetails, secret string) *Model {
	return &Model{
		DatabentoModel: &types.DatabentoModel{
			SettingCursor: 0,
			MainCursor:    0,
			Screen:        types.MainMenuScreen,
			Focus:         types.MainFocus,
			Cfg:           cfg,
			Secret:        secret,
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "tab":
			m.MainCursor = 0
			if m.Focus == types.MainFocus {
				m.Focus = types.SettingFocus
			} else {
				m.Focus = types.MainFocus
			}
		case "enter":
			switch m.Focus {
			case types.MainFocus:
				switch m.Screen {
				case types.MainMenuScreen:
					m.Screen = ui.Menu[m.MainCursor].Target
				case types.HistMenuScreen:
					m.Screen = ui.HistoricalMenu[m.MainCursor].Target
				case types.HistDataset:
					m.Query.Dataset = ui.Datasets[m.MainCursor]
					if parent, ok := ui.Parent[m.Screen]; ok {
						m.Screen = parent
					}
				case types.HistSymbol:
					m.Query.Symbol = ui.Symbols[m.MainCursor]
					if parent, ok := ui.Parent[m.Screen]; ok {
						m.Screen = parent
					}
				case types.HistSchema:
					m.Query.Schema = ui.Schemas[m.MainCursor]
					if parent, ok := ui.Parent[m.Screen]; ok {
						m.Screen = parent
					}
				case types.QuerySubmit:
					err := api.FetchHistory()
					if err != nil {
						fmt.Println(err)
					}
				}
			case types.SettingFocus:
			}
			m.MainCursor = 0
		case "h", "backspace":
			switch m.Screen {
			case types.HistMenuScreen, types.LiveMenuScreen:
				m.MainCursor = 0
				m.Query = types.Query{}
				m.Screen = types.MainMenuScreen
			case types.HistSymbol, types.HistSchema, types.HistStart, types.HistEnd, types.HistLimit, types.LiveSymbol, types.LiveSchema:
				if parent, ok := ui.Parent[m.Screen]; ok {
					m.Screen = parent
				}
			}
		case "j", "down":
			switch m.Screen {
			case types.MainMenuScreen:
				if m.MainCursor < len(ui.Menu)-1 {
					m.MainCursor++
				}
			case types.HistMenuScreen:
				if m.MainCursor < len(ui.HistoricalMenu)-1 {
					m.MainCursor++
				}
			case types.LiveMenuScreen:
				if m.MainCursor < len(ui.LiveMenu)-1 {
					m.MainCursor++
				}
			case types.HistSymbol:
				if m.MainCursor < len(ui.Symbols)-1 {
					m.MainCursor++
				}
			case types.HistSchema:
				if m.MainCursor < len(ui.Schemas)-1 {
					m.MainCursor++
				}
			}
		case "k", "up":
			switch m.Screen {
			case types.MainMenuScreen, types.HistMenuScreen, types.LiveMenuScreen, types.HistSymbol, types.HistSchema:
				if m.MainCursor > 0 {
					m.MainCursor--
				}
			}
		}
	}

	return m, cmd
}

func (m Model) View() tea.View {
	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			ui.MainPanel(m.DatabentoModel),
			ui.SettingButton(m.DatabentoModel),
		),
	)
}
