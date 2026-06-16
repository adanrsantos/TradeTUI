package databento

import (
	tea "charm.land/bubbletea/v2"
	"github.com/adanrsantos/TradeTUI/types"
)

type model struct {
	cfg    *types.Config
	cursor int
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
	return m, nil
}

func (m model) View() tea.View {
	return tea.NewView("Databento")
}
