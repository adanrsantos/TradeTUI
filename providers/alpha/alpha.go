package alpha

import (
	tea "charm.land/bubbletea/v2"
	"github.com/adanrsantos/TradeTUI/types"
)

type model struct {
	model alphaModel
}

func New(cfg *types.ProviderDetails) *model {
	return &model{
		model: alphaModel{
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
	return m, nil
}

func (m model) View() tea.View {
	return tea.NewView("Alpha")
}
