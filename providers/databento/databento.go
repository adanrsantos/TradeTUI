package databento

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	//"fmt"
	//"github.com/adanrsantos/TradeTUI/providers/databento/api"
	"charm.land/bubbles/v2/textinput"
	"github.com/adanrsantos/TradeTUI/providers/databento/types"
	"github.com/adanrsantos/TradeTUI/providers/databento/ui"
	globalTypes "github.com/adanrsantos/TradeTUI/types"
)

type Model struct {
	*types.DatabentoModel
}

func New(cfg *globalTypes.ProviderDetails, secret string) *Model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Prompt = ""
	ti.SetWidth(20)

	return &Model{
		DatabentoModel: &types.DatabentoModel{
			MainCursor:    0,
			SubmitCursor:  -1,
			SettingCursor: -1,
			TimeInput:     ti,
			Mode:          types.NormalMode,
			Screen:        types.MainMenuScreen,
			Focus:         types.MainFocus,
			Cfg:           cfg,
			Secret:        secret,
		},
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.Mode == types.EditMode && (m.Screen == types.HistStart || m.Screen == types.HistEnd) {
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			s := msg.String()
			switch msg.String() {
			case "enter":
				value := m.TimeInput.Value()
				switch m.Screen {
				case types.HistStart:
					m.Query.StartDate = &types.TimePreset{
						Display: value,
						Value:   types.TimeValue(value),
					}
				case types.HistEnd:
					m.Query.EndDate = &types.TimePreset{
						Display: value,
						Value:   types.TimeValue(value),
					}
				}

				m.TimeInput.Reset()
				m.TimeInput.Blur()
				m.Mode = types.NormalMode
				m.GoBack()
			case "esc":
				m.TimeInput.Blur()
				m.Mode = types.NormalMode
			case "backspace":
				value := m.TimeInput.Value()
				if len(value) == 5 && value[4:] == "-" {
					m.TimeInput.SetValue(value[:4])
					m.TimeInput.CursorEnd()
				} else if len(value) == 8 && value[7:] == "-" {
					m.TimeInput.SetValue(value[:7])
					m.TimeInput.CursorEnd()
				}
			default:
				if len(s) == 1 && (s[0] < '0' || s[0] > '9') {
					return m, nil
				}
			}
		}
		m.TimeInput, cmd = m.TimeInput.Update(msg)

		value := m.TimeInput.Value()

		if len(value) == 4 && value[4:] != "-" {
			m.TimeInput.SetValue(value + "-")
			m.TimeInput.CursorEnd()
		}

		if len(value) == 7 && value[7:] != "-" {
			m.TimeInput.SetValue(value + "-")
			m.TimeInput.CursorEnd()
		}
		return m, cmd
	}
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "tab":
			switch m.Screen {
			case types.HistMenuScreen:
				if m.Focus == types.MainFocus {
					m.Focus = types.SubmitFocus
					m.MainCursor = -1
					m.SubmitCursor = 0
				} else {
					m.Focus = types.MainFocus
					m.MainCursor = 0
					m.SubmitCursor = -1
				}
			}
		case "l", "enter":
			switch m.Screen {
			case types.MainMenuScreen:
				m.GoForward(ui.Menu)
			case types.HistMenuScreen:
				switch m.Focus {
				case types.MainFocus:
					if m.MainCursor == int(types.DatasetField) {
						m.GoForward(ui.HistoricalMenu)
						return m, cmd
					} else {
						switch m.MainCursor {
						case int(types.SymbolField), int(types.SchemaField):
							if m.Query.Dataset == nil {
								return m, cmd
							}
						case int(types.StartField), int(types.EndField):
							if m.Query.Schema == nil {
								return m, cmd
							}
						}
					}
					m.GoForward(ui.HistoricalMenu)
				case types.SubmitFocus:
					action := ui.SubmitChoices[m.SubmitCursor].Action

					switch action {
					case types.SubmitAction:
						m.Screen = types.HistRequest
						// call function
					case types.ResetAction:
						m.Query = types.Query{}
					}
					m.Focus = types.MainFocus
					m.MainCursor = 0
					m.SubmitCursor = -1
				}
			case types.HistDataset:
				m.Query.Dataset = &ui.Datasets[m.MainCursor]
				m.GoBack()
			case types.HistSymbol:
				m.Query.Symbol = &m.Query.Dataset.Symbols[m.MainCursor]
				m.GoBack()
			case types.HistSchema:
				m.Query.Schema = &ui.Schemas[m.MainCursor]
				m.GoBack()
			case types.HistStart:
				presets := m.Query.Schema.CompatiblePresets
				if m.MainCursor == len(presets)-1 {
					m.Mode = types.EditMode
					m.TimeInput.Focus()
					m.TimeInput.Placeholder = m.Query.Schema.TimeRange
					m.TimeInput.CharLimit = len(m.Query.Schema.TimeRange)
				} else {
					m.Query.StartDate = &presets[m.MainCursor]
					m.GoBack()
				}
			case types.HistEnd:
				presets := m.Query.Schema.CompatiblePresets
				if m.MainCursor == len(presets)-1 {
					m.Mode = types.EditMode
					m.TimeInput.Focus()
					m.TimeInput.Placeholder = m.Query.Schema.TimeRange
					m.TimeInput.CharLimit = len(m.Query.Schema.TimeRange)
				} else {
					m.Query.EndDate = &presets[m.MainCursor]
					m.GoBack()
				}
			case types.HistLimit:
				m.GoBack()
			default:
				m.GoBack()
			}
		case "h", "backspace":
			switch m.Screen {
			case types.HistMenuScreen:
				m.SubmitCursor = -1
			}
			m.GoBack()
		case "j", "down":
			switch m.Screen {
			case types.MainMenuScreen:
				m.MainCursor = IncreaseCursor(m.MainCursor, len(ui.Menu))
			case types.HistMenuScreen:
				switch m.Focus {
				case types.MainFocus:
					m.MainCursor = IncreaseCursor(m.MainCursor, len(ui.HistoricalMenu))
				case types.SubmitFocus:
					m.SubmitCursor = IncreaseCursor(m.SubmitCursor, len(ui.SubmitChoices))
				}
			case types.HistSymbol:
				m.MainCursor = IncreaseCursor(m.MainCursor, len(ui.FutureSymbols))
			case types.HistSchema:
				m.MainCursor = IncreaseCursor(m.MainCursor, len(ui.Schemas))
			case types.HistStart:
				m.MainCursor = IncreaseCursor(m.MainCursor, len(m.Query.Schema.CompatiblePresets))
			case types.HistEnd:
				m.MainCursor = IncreaseCursor(m.MainCursor, len(m.Query.Schema.CompatiblePresets))
			case types.HistRequest:
				// request screen logic
			}
		case "k", "up":
			switch m.Focus {
			case types.MainFocus:
				if m.MainCursor > 0 {
					m.MainCursor--
				}
			case types.SubmitFocus:
				if m.SubmitCursor > 0 {
					m.SubmitCursor--
				}
			}
		}
	}

	return m, cmd
}

func (m *Model) View() tea.View {
	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			ui.Dashboard(m.DatabentoModel),
		),
	)
}

func (m *Model) GoBack() {
	if parent, ok := ui.Parent[m.Screen]; ok {
		m.Screen = parent
	}

	if len(m.CursorStack) == 0 {
		return
	}

	if m.Focus == types.SubmitFocus {
		m.Focus = types.MainFocus
		m.SubmitCursor = -1
	}

	last := len(m.CursorStack) - 1
	m.MainCursor = m.CursorStack[last]
	m.CursorStack = m.CursorStack[:last]
}

func (m *Model) GoForward(menu []types.MenuItem) {
	m.CursorStack = append(m.CursorStack, m.MainCursor)
	m.Screen = menu[m.MainCursor].Target
	m.MainCursor = 0
}

func IncreaseCursor(cursor int, max int) int {
	if cursor < max-1 {
		return cursor + 1
	}
	return cursor
}
