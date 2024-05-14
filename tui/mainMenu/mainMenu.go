package mainMenu

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/christianhturner/api-workbench/project"
	"github.com/christianhturner/api-workbench/settings"
)

type mainMenuSelection struct {
	title string
}

type (
	updateProjectListMsg struct{}
	errMsg               struct{ error }
)

type MainMenu struct {
	selected   map[int]struct{}
	selections []mainMenuSelection
	input      textinput.Model
	output     []string
	cursor     int
}

func InitialModel() MainMenu {
	input := textinput.New()
	input.Prompt = "$"
	input.Placeholder = "Project Name: "
	input.CharLimit = 250
	input.Width = 50
	return MainMenu{
		input: input,
		selections: []mainMenuSelection{
			{title: "Start API Server"},
			{title: "Create new Project"},
		},
		selected: make(map[int]struct{}),
	}
}

func createNewProjectCmd(name string) tea.Cmd {
	return func() tea.Msg {
		_, err := project.New(name)
		if err != nil {
			log.Panic("Error creating project %v", err)
			return errMsg{err}
		}
		return updateProjectListMsg{}
	}
}

func (m MainMenu) Init() tea.Cmd {
	return tea.SetWindowTitle("API Workbench Main Menu")
}

func (m MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	var cmds []tea.Cmd
	km := settings.DefaultKeyMap("")
	switch msg := msg.(type) {
	case updateProjectListMsg:
		{
			fmt.Print("Projects Updated")
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, km.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, km.Down):
			if m.cursor < len(m.selections)-1 {
				m.cursor++
			}
		case key.Matches(msg, km.Enter):
			selection := m.selections[m.cursor]
			if selection.title == "Create new Project" {
				m.input.Focus()
				if m.input.Focused() {
					if key.Matches(msg, km.Enter) {
						cmds = append(cmds, createNewProjectCmd(m.input.Value()))
						m.input.SetValue("")
					}
				}
				// projectName := ""
				// fmt.Print("Enter name for the new project: ")
				// fmt.Scanln(&projectName)
				// newProject, err := project.New(projectName
				// if err != nil {
				// 	log.Printf("Error creating project: %s\n", err)
				// 	printError := fmt.Sprintf("Error creating project: %s\n", err)
				// 	m.output = append(m.output, printError)
				// } else {
				// 	printSuccess := fmt.Sprintf("Project '%s' created successfully.\n", newProject)
				// 	m.output = append(m.output, printSuccess)
				// }
			} else {
				printWorld := fmt.Sprintf("hello, %s\n", selection.title)
				m.output = append(m.output, printWorld)
			}

		case key.Matches(msg, km.Quit):
			return m, tea.Quit
		}
	}
	return m, tea.Batch(cmds...)
}

func (m MainMenu) View() string {
	s := "Make a selection for what command you want to run?\n\n"
	for i, choice := range m.selections {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s | %s\n", cursor, choice.title)
		m.input.Focus()
	}
	for i := range m.output {
		s += m.output[i]
	}
	s += "\nPress q to quit.\n"
	return s
}
