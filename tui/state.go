package tui

import tea "github.com/charmbracelet/bubbletea"

type sessionState int

const (
	mainMenu sessionState = iota
	createNewApi
	navigateRoute
)

type StateModel struct {
	state sessionState
}

func InitialModel() StateModel {
	return StateModel{
		state: mainMenu,
	}
}

func (m StateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	}
}
