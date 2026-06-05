package databento

import (
	"time"

	"github.com/adanrsantos/TradeTUI/internals/providers"
	"github.com/adanrsantos/TradeTUI/internals/providers/databento/types"
	"github.com/adanrsantos/TradeTUI/internals/providers/databento/ui"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type DataBentoModel struct {
	types.Model
	width  int
	height int
}

func New() providers.Provider {
	return DataBentoModel{}
}

func (m DataBentoModel) Name() string {
	return "DataBento"
}

func (m DataBentoModel) Init() tea.Cmd {
	return nil
}

func (m DataBentoModel) Update(msg tea.Msg) (providers.Provider, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "crtl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Focus == ui.MainMenuFocus && m.MainMenuCursor > 0 {
				m.MainMenuCursor--
			}
			if m.Focus == ui.SubmitFocus && m.SubmitCursor > 0 {
				m.SubmitCursor--
			}

		case "down", "j":
			if m.Focus == ui.MainMenuFocus && m.MainMenuCursor < len(ui.MenuChoices(&m.Model))-1 {
				m.MainMenuCursor++
			}

		case "enter":
			switch m.Focus {
			case ui.MainMenuFocus:
				switch m.Screen {
				case ui.MainMenuView:
					if m.MainMenuCursor < len(ui.MainMenuChoices) {
						m.Screen = ui.MainMenuChoices[m.MainMenuCursor].Target
					}
				case ui.TimeFrameMenuView:
					if m.MainMenuCursor < len(ui.TimeFrameChoices) {
						m.Config.TimeFrame = ui.TimeFrameChoices[m.MainMenuCursor]
					}
					m.Screen = ui.MainMenuView
				case ui.SymbolMenuView:
					if m.MainMenuCursor < len(ui.SymbolChoices) {
						m.Config.Symbol = ui.SymbolChoices[m.MainMenuCursor]
					}
					m.Screen = ui.MainMenuView
				case ui.StartDateMenuView:
					if m.MainMenuCursor < len(ui.DatePresets) {
						m.Config.StartDate = ui.DatePresets[m.MainMenuCursor].Value()
					}
					m.Screen = ui.MainMenuView
				case ui.EndDateMenuView:
					if m.MainMenuCursor < len(ui.DatePresets) {
						m.Config.EndDate = ui.DatePresets[m.MainMenuCursor].Value()
					}
					m.Screen = ui.MainMenuView
				case ui.LimitMenuView:
					if m.MainMenuCursor < len(ui.LimitChoices) {
						m.Config.Limit = int(ui.LimitChoices[m.MainMenuCursor])
					}
					m.Screen = ui.MainMenuView
				}
				m.MainMenuCursor = 0
			case ui.SubmitFocus:
				switch m.SubmitCursor {
				case 0:
					if m.Config.TimeFrame != "" && m.Config.Symbol != "" && !m.Config.StartDate.IsZero() && !m.Config.EndDate.IsZero() && m.Config.Limit >= 0 {
						newItem := types.HistoryItem{
							Config:    m.Config,
							Timestamp: time.Now(),
						}

						m.History = append(m.History, newItem)

						m.Config = types.QueryConfig{}
						m.Config.Limit = -1
						m.Err = ""
					} else {
						m.Err = "Missing fields!"
					}
				case 1:
					m.Config = types.QueryConfig{}
					m.Config.Limit = -1
					m.Err = ""
				}
			}

		case "esc", "backspace":
			if m.Screen != ui.MainMenuView {
				m.MainMenuCursor = 0
				m.Screen = ui.MainMenuView
			}

		case "tab":
			m.MainMenuCursor = 0
			m.SubmitCursor = 0
			if m.Focus == ui.MainMenuFocus {
				m.Focus = ui.SubmitFocus
			} else {
				m.Focus = ui.MainMenuFocus
			}
		}
	}

	return m, nil
}

func (m DataBentoModel) View(width, height int) tea.View {
	minLeftWidth := 50
	minRightWidth := 25

	leftWidth := int(float64(m.width) * 0.6)
	rightWidth := m.width - leftWidth

	if leftWidth < minLeftWidth {
		leftWidth = minLeftWidth
	}
	if rightWidth < minRightWidth {
		rightWidth = minRightWidth
	}

	return tea.NewView(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
		),
	)
}
