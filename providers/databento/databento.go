package databento

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"errors"
	"github.com/adanrsantos/TradeTUI/providers/databento/api"
	"github.com/adanrsantos/TradeTUI/providers/databento/types"
	"github.com/adanrsantos/TradeTUI/providers/databento/ui"
	globalTypes "github.com/adanrsantos/TradeTUI/types"
	"strconv"
	"time"
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
			Input:         ti,
			ErrInput:      nil,
			Mode:          types.NormalMode,
			Query:         types.Query{},
			QueryEstimate: 0,
			ErrQuery:      nil,
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
			switch s {
			case "enter", "l":
				value := m.Input.Value()

				t, err := m.ValidateQuery(value)
				if err != nil {
					m.ErrInput = err
					return m, nil
				}

				switch m.Screen {
				case types.HistStart:
					m.Query.StartDate = t
				case types.HistEnd:
					m.Query.EndDate = t
				}

				m.Input.Reset()
				m.Input.Blur()
				m.Mode = types.NormalMode
				m.ErrInput = nil
				m.GoBack()
			case "esc", "h":
				m.Input.Reset()
				m.Input.Blur()
				m.Mode = types.NormalMode
				m.ErrInput = nil
				m.GoBack()
			case "backspace":
				m.DeleteInputChar()
			default:
				if len(s) == 1 && (s[0] < '0' || s[0] > '9') {
					return m, nil
				}
			}
		}
		m.Input, cmd = m.Input.Update(msg)

		m.ValidateTime()

		return m, cmd
	}

	if m.Mode == types.EditMode && m.Screen == types.HistLimit {
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			s := msg.String()
			switch s {
			case "enter", "l":
				value := m.Input.Value()

				limit, err := strconv.Atoi(value)
				if err != nil {
					m.ErrInput = err
					return m, nil
				}

				m.Query.Limit = &limit

				m.Input.Reset()
				m.Input.Blur()
				m.Mode = types.NormalMode
				m.ErrInput = nil
				m.GoBack()
			case "esc", "h":
				m.Input.Reset()
				m.Input.Blur()
				m.Mode = types.NormalMode
				m.ErrInput = nil
				m.GoBack()
			case "backspace":
			default:
				if len(s) == 1 && (s[0] < '0' || s[0] > '9') {
					return m, nil
				}

				if m.Input.Value() == "0" && s != "backspace" {
					return m, nil
				}
			}
		}
		m.Input, cmd = m.Input.Update(msg)

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
					screen := ui.HistoricalMenu[m.MainCursor].Target
					switch screen {
					case types.HistStart, types.HistEnd:
						m.Mode = types.EditMode
						m.Input.Focus()
						format := ui.TimeFormats[m.Query.Schema.Value]
						m.Input.Placeholder = format.Placeholder
						m.Input.CharLimit = format.CharLimit
					case types.HistLimit:
						m.Mode = types.EditMode
						m.Input.Focus()
						m.Input.Placeholder = "0"
						m.Input.CharLimit = 4
					}
					m.GoForward(ui.HistoricalMenu)
				case types.SubmitFocus:
					action := ui.SubmitChoices[m.SubmitCursor].Action

					switch action {
					case types.SubmitAction:
						if m.Query.Dataset == nil || m.Query.Symbol == nil || m.Query.Schema == nil || m.Query.StartDate == nil || m.Query.EndDate == nil || m.Query.Limit == nil {
							m.ErrQuery = errors.New("Incomplete query. All 5 fields must be filled")
							m.Focus = types.MainFocus
							m.MainCursor = 0
							m.SubmitCursor = -1

							return m, nil
						}
						estimate, err := api.HistoryEstimateCost(m.Query, m.Secret)
						if err != nil {
							m.ErrQuery = err
							m.Focus = types.MainFocus
							m.MainCursor = 0
							m.SubmitCursor = -1

							return m, nil
						}
						m.QueryEstimate = estimate
						m.GoForward(ui.SubmitChoices)
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
			case types.HistLimit:
				m.GoBack()
			case types.HistRequest:
				action := ui.RequestChoices[m.MainCursor].Action
				switch action {
				case types.CancelAction:
					m.GoBack()
				case types.ContinueAction:
					candles, err := api.HistoryRequest(m.Query, m.Secret)
					if err != nil {
						m.ErrQuery = err
						return m, nil
					}
					err = api.SaveCandles(candles)
					if err != nil {
						m.ErrQuery = err
						return m, nil
					}
					m.GoBack()
					m.GoBack()
					m.Query = types.Query{}
				}
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
			case types.HistRequest:
				m.MainCursor = IncreaseCursor(m.MainCursor, len(ui.RequestChoices))
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
	if m.Focus == types.SubmitFocus {
		m.MainCursor = 0
		m.SubmitCursor = -1
	}
	m.CursorStack = append(m.CursorStack, m.MainCursor)
	m.Screen = menu[m.MainCursor].Target
	m.MainCursor = 0
	m.Focus = types.MainFocus
	m.ErrQuery = nil
}

func IncreaseCursor(cursor int, max int) int {
	if cursor < max-1 {
		return cursor + 1
	}
	return cursor
}

func (m *Model) ValidateQuery(value string) (*time.Time, error) {
	format := ui.TimeFormats[m.Query.Schema.Value]

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return nil, err
	}

	t, err := time.ParseInLocation(
		format.Layout,
		value,
		location,
	)
	if err != nil {
		return nil, err
	}

	switch m.Screen {
	case types.HistStart:
		if m.Query.EndDate != nil && t.After(*m.Query.EndDate) {
			return nil, errors.New("Start time cannot be after end time")
		}
		if m.Query.EndDate != nil && t.Equal(*m.Query.EndDate) {
			return nil, errors.New("Start time cannot be the same as end time")
		}
	case types.HistEnd:
		if m.Query.StartDate != nil && t.Before(*m.Query.StartDate) {
			return nil, errors.New("End time cannot be before start time")
		}
		if m.Query.StartDate != nil && t.Equal(*m.Query.StartDate) {
			return nil, errors.New("End time cannot be the same as start time")
		}
	}

	return &t, nil
}

func (m *Model) ValidateTime() {
	value := m.Input.Value()
	switch len(value) {
	case 1:
		if value[0] != '1' && value[0] != '2' {
			m.Input.SetValue("")
		}
	case 2:
		if value[0] == '1' {
			if value[1] != '9' {
				m.Input.SetValue(value[:1])
			}
		} else {
			if value[1] != '0' {
				m.Input.SetValue(value[:1])
			}
		}
	case 3:
		if value[1] == '0' {
			if value[2] != '0' && value[2] != '1' && value[2] != '2' {
				m.Input.SetValue(value[:2])
			}
		}
	case 4:
		if value[1] == '0' && value[2] == '2' {
			if value[3] > '6' {
				m.Input.SetValue(value[:3])
			}
		}
		if len(m.Input.Value()) == 4 {
			m.Input.SetValue(value + "-")
		}
	case 6:
		if value[5] != '0' && value[5] != '1' {
			m.Input.SetValue(value[:5])
		}
	case 7:
		if value[5] == '1' {
			if value[6] != '0' && value[6] != '1' && value[6] != '2' {
				m.Input.SetValue(value[:6])
			}
		} else {
			if value[6] == '0' {
				m.Input.SetValue(value[:6])
			}
		}
		if len(m.Input.Value()) == 7 {
			m.Input.SetValue(value + "-")
		}
	case 9:
		if value[8] > '3' {
			m.Input.SetValue(value[:8])
		}
	case 10:
		if value[8] == '3' {
			if value[9] > '1' {
				m.Input.SetValue(value[:9])
			}
		} else if value[8] == '0' {
			if value[9] == '0' {
				m.Input.SetValue(value[:9])
			}
		}
		if len(m.Input.Value()) == 10 {
			input := ui.TimeFormats[m.Query.Schema.Value]
			if input.CharLimit > 10 {
				m.Input.SetValue(value + "T")
			}
		}
	case 12:
		if value[11] > '2' {
			m.Input.SetValue(value[:11])
		}
	case 13:
		if value[11] == '2' {
			if value[12] > '3' {
				m.Input.SetValue(value[:12])
			}
		}
		if len(m.Input.Value()) == 13 {
			input := ui.TimeFormats[m.Query.Schema.Value]
			if input.CharLimit > 13 {
				m.Input.SetValue(value + ":")
			}
		}
	case 15:
		if value[14] > '5' {
			m.Input.SetValue(value[:14])
		}
	case 16:
		input := ui.TimeFormats[m.Query.Schema.Value]
		if input.CharLimit > 16 {
			m.Input.SetValue(value + ":")
		}
	case 18:
		if value[17] > '5' {
			m.Input.SetValue(value[:17])
		}
	}
	m.Input.CursorEnd()
}

func (m *Model) DeleteInputChar() {
	value := m.Input.Value()
	if len(value) == 5 && value[4:] == "-" {
		m.Input.SetValue(value[:4])
	} else if len(value) == 8 && value[7:] == "-" {
		m.Input.SetValue(value[:7])
	} else if len(value) == 11 && value[10:] == "T" {
		m.Input.SetValue(value[:10])
	} else if len(value) == 14 && value[13:] == ":" {
		m.Input.SetValue(value[:13])
	} else if len(value) == 17 && value[16:] == ":" {
		m.Input.SetValue(value[:16])
	}
	m.Input.CursorEnd()
}
