package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/christianhturner/api-workbench/project"
	"github.com/christianhturner/api-workbench/tui/common"
	"github.com/christianhturner/api-workbench/tui/projectTui"
)

func StartTea(projectDb project.DB) error {
	common.ProjectDB = &projectDb

	m, _ := projectTui.InitProject()
	common.P = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := common.P.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	return nil
}
