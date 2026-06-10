package databento

import tea "charm.land/bubbletea/v2"

type model struct {
	cursor int
}

func New() model {
	return model{}
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
