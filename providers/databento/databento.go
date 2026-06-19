package databento

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/adanrsantos/TradeTUI/types"
)

type model struct {
	cfg   *types.Config
	model databentoModel
}

func New(cfg *types.Config) *model {
	return &model{
		cfg: cfg,
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
		case "tab":
			if m.model.focus == MainFocus {
				m.model.focus = SettingFocus
			} else {
				m.model.focus = MainFocus
			}
		}
	}

	return m, cmd
}

func (m model) View() tea.View {
	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			APIMenu(m.model),
			SettingButton(m.model),
		),
	)
}
