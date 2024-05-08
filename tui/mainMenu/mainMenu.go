package mainMenu

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/christianhturner/api-workbench/settings"
)

type mainMenuSelection struct {
	title string
}

type MainMenu struct {
	selected   map[int]struct{}
	selections []mainMenuSelection
	cursor     int
	output     []string
}

func InitialModel() MainMenu {
	return MainMenu{
		selections: []mainMenuSelection{
			{title: "Start API Server"},
			{title: "Create new Route"},
		},
		selected: make(map[int]struct{}),
	}
}

func (m MainMenu) Init() tea.Cmd {
	return tea.SetWindowTitle("API Workbench Main Menu")
}

func (m MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, settings.DefaultKeyMap.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, settings.DefaultKeyMap.Down):
			if m.cursor < len(m.selections)-1 {
				m.cursor++
			}
		case key.Matches(msg, settings.DefaultKeyMap.Enter):
			selection := m.selections[m.cursor]
			printWorld := fmt.Sprintf("hello, %s\n", selection.title)
			m.output = append(m.output, printWorld)

		case key.Matches(msg, settings.DefaultKeyMap.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MainMenu) View() string {
	s := "Make a selection for what command you want to run?\n\n"
	for i, choice := range m.selections {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s | %s\n", cursor, choice.title)
	}
	for i := range m.output {
		s += m.output[i]
	}
	s += "\nPress q to quit.\n"
	return s
}
