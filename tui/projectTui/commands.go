package projectTui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/christianhturner/api-workbench/project"
)

type (
	errMsg               struct{ error }
	updateProjectListMsg struct{}
	renameProjectMsg     []list.Item

	// SelectMsg the message to change the view to the selected entry
	SelectMsg struct{ ActiveProjectID uint }
)

func createProjectCmd(name string, db *project.DB) tea.Cmd {
	return func() tea.Msg {
		_, err := db.CreateProject(name)
		if err != nil {
			return errMsg{err}
		}
		return updateProjectListMsg{}
	}
}

func renameProjectCmd(id uint, db *project.DB, name string) tea.Cmd {
	return func() tea.Msg {
		db.RenameProject(id, name)
		projects, err := db.GetAllProjects()
		if err != nil {
			return errMsg{err}
		}
		items := projectsToItems(projects)
		return renameProjectMsg(items)
	}
}

func deleteProjectCmd(id uint, pr *project.DB) tea.Cmd {
	return func() tea.Msg {
		err := pr.DeleteProject(id)
		if err != nil {
			return errMsg{err}
		}
		return updateProjectListMsg{}
	}
}
