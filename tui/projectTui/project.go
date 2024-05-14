package projectTui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/christianhturner/api-workbench/project"
	"github.com/christianhturner/api-workbench/settings"
	"github.com/christianhturner/api-workbench/tui/common"
)

type mode int

const (
	nav mode = iota
	edit
	create
)

type Model struct {
	mode     mode
	list     list.Model
	input    textinput.Model
	quitting bool
}

func InitProject() (tea.Model, tea.Cmd) {
	input := textinput.New()
	input.Prompt = "$ "
	input.Placeholder = "Project name..."
	input.CharLimit = 250
	input.Width = 50

	items, err := newProjectList(common.ProjectDB)
	m := Model{mode: nav, list: list.NewModel(items, list.NewDefaultDelegate(), 8, 8), input: input}
	if common.WindowSize.Height != 0 {
		top, right, bottom, left := lipgloss.NewStyle().Margin(0, 2).GetMargin()
		m.list.SetSize(common.WindowSize.Width-left-right, common.WindowSize.Height-top-bottom-1)
	}
	m.list.Title = "projects"
	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			settings.DefaultKeyMap("project").Create,
			settings.DefaultKeyMap(m.list.Title).Rename,
			settings.DefaultKeyMap(m.list.Title).Delete,
			settings.DefaultKeyMap("").Back,
		}
	}
	return m, func() tea.Msg { return errMsg{err} }
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update handle IO and commands
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		common.WindowSize = msg
		top, right, bottom, left := lipgloss.NewStyle().Margin(0, 2).GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
	case updateProjectListMsg:
		projects, err := common.ProjectDB.GetAllProjects()
		if err != nil {
			log.Fatal(err)
		}
		items := projectsToItems(projects)
		m.list.SetItems(items)
		m.mode = nav
	case renameProjectMsg:
		m.list.SetItems(msg)
		m.mode = nav
	case tea.KeyMsg:
		if m.input.Focused() {
			if key.Matches(msg, settings.DefaultKeyMap("").Enter) {
				if m.mode == create {
					cmds = append(cmds, createProjectCmd(m.input.Value(), common.ProjectDB))
				}
				if m.mode == edit {
					cmds = append(cmds, renameProjectCmd(m.getActiveProjectID(), common.ProjectDB, m.input.Value()))
				}
				m.input.SetValue("")
				m.mode = nav
				m.input.Blur()
			}
			if key.Matches(msg, settings.DefaultKeyMap("").Back) {
				m.input.SetValue("")
				m.mode = nav
				m.input.Blur()
			}
			// only log keypresses for the input field when it's focused
			m.input, cmd = m.input.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			switch {
			case key.Matches(msg, settings.DefaultKeyMap("project").Create):
				m.mode = create
				m.input.Focus()
				cmd = textinput.Blink
			case key.Matches(msg, settings.DefaultKeyMap("").Quit):
				m.quitting = true
				return m, tea.Quit
			// case key.Matches(msg, settings.DefaultKeyMap("").Enter):
			// 	activeProject := m.list.SelectedItem().(project.Project)
			// 	entry := InitEntry(constants.Er, activeProject.ID, constants.P)
			// 	return entry.Update(constants.WindowSize)
			case key.Matches(msg, settings.DefaultKeyMap(m.list.Title).Rename):
				m.mode = edit
				m.input.Focus()
				cmd = textinput.Blink
			case key.Matches(msg, settings.DefaultKeyMap(m.list.Title).Delete):
				items := m.list.Items()
				if len(items) > 0 {
					cmd = deleteProjectCmd(m.getActiveProjectID(), common.ProjectDB)
				}
			default:
				m.list, cmd = m.list.Update(msg)
			}
			cmds = append(cmds, cmd)
		}
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}
	if m.input.Focused() {
		return lipgloss.NewStyle().Margin(0, 2).Render(m.list.View() + "\n" + m.input.View())
	}
	return lipgloss.NewStyle().Margin(0, 2).Render(m.list.View() + "\n")
}

func newProjectList(db *project.DB) ([]list.Item, error) {
	projects, err := db.GetAllProjects()
	if err != nil {
		return nil, fmt.Errorf("cannot get all projects: %w", err)
	}
	return projectsToItems(projects), err
}

func projectsToItems(projects []project.Project) []list.Item {
	items := make([]list.Item, len(projects))
	for i, proj := range projects {
		items[i] = list.Item(proj)
	}
	return items
}

func (m Model) getActiveProjectID() uint {
	items := m.list.Items()
	activeItem := items[m.list.Index()]
	return activeItem.(project.Project).ID
}
