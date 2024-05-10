package styles

import "github.com/charmbracelet/lipgloss"

// Reference: https://github.com/charmbracelet/soft-serve/blob/main/pkg/ui/styles/styles.go

type Styles struct {
	ActiveBorderColor   lipgloss.Color
	InactiveBorderColor lipgloss.Color

	App               lipgloss.Style
	ProjectName       lipgloss.Style
	TopLevelNormalTab lipgloss.Style
	TopLevelActiveTab lipgloss.Style

	MenuItem       lipgloss.Style
	MenuLastUpdate lipgloss.Style

	ProjectSelector struct {
		Normal struct {
			Base    lipgloss.Style
			Title   lipgloss.Style
			Desc    lipgloss.Style
			Command lipgloss.Style
			Updated lipgloss.Style
		}
		Active struct {
			Base    lipgloss.Style
			Title   lipgloss.Style
			Desc    lipgloss.Style
			Command lipgloss.Style
			Updated lipgloss.Style
		}
	}

	Spinner          lipgloss.Style
	SpinnerContainer lipgloss.Style

	Tabs         lipgloss.Style
	TabInactive  lipgloss.Style
	TabActive    lipgloss.Style
	TabSeparator lipgloss.Style

	StatusBar      lipgloss.Style
	StatusBarKey   lipgloss.Style
	StatusBarValue lipgloss.Style
	StatusBarInfo  lipgloss.Style
	StatusBarHelp  lipgloss.Style

	// TODO: More context as project develops
}

func DefaultStyles(r *lipgloss.Renderer) *Styles {
	purple := lipgloss.Color("#9580ff")
	pink := lipgloss.Color("#ff80bf")
	orange := lipgloss.Color("#ffca80")
	yellow := lipgloss.Color("#ffff80")
	green := lipgloss.Color("#8aff80")
	blue := lipgloss.Color("#80ffea")
	dark := lipgloss.Color("#21222B")
	light := lipgloss.Color("##F6F6F0")

	s := new(Styles)

	s.ActiveBorderColor = pink
	s.InactiveBorderColor = yellow

	s.App = r.NewStyle().
		Margin(1, 2)

	s.ProjectName = r.NewStyle().
		Height(1).
		MarginLeft(1).
		MarginBottom(1).
		Padding(0, 1).
		Background(dark).
		Foreground(light).
		Bold(true)

	s.TopLevelNormalTab = r.NewStyle().
		MarginRight(2)

	s.TopLevelActiveTab = s.TopLevelNormalTab.Copy().
		Foreground(purple)

	s.ProjectSelector.Normal.Base = r.NewStyle().
		PaddingLeft(1).
		Border(lipgloss.Border{Left: " "}, false, false, false, true).
		Height(3)

	s.ProjectSelector.Normal.Title = r.NewStyle().Bold(true)

	s.ProjectSelector.Normal.Desc = r.NewStyle().
		Foreground(orange)

	s.ProjectSelector.Normal.Command = r.NewStyle().
		Foreground(blue)

	s.ProjectSelector.Normal.Updated = r.NewStyle().
		Foreground(green)

	s.ProjectSelector.Active.Base = s.ProjectSelector.Normal.Base.Copy().
		BorderStyle(lipgloss.Border{Left: "|"}).
		BorderForeground(pink)

	s.ProjectSelector.Active.Title = s.ProjectSelector.Normal.Title.Copy()

	s.ProjectSelector.Active.Desc = s.ProjectSelector.Normal.Desc.Copy()

	s.ProjectSelector.Active.Command = s.ProjectSelector.Normal.Command.Copy()

	s.ProjectSelector.Active.Updated = s.ProjectSelector.Normal.Updated.Copy()

	s.MenuItem = r.NewStyle().
		PaddingLeft(1).
		Border(lipgloss.Border{
			Left: " ",
		}, false, false, false, true).
		Height(3)

	s.MenuLastUpdate = r.NewStyle().
		Foreground(pink).
		Align(lipgloss.Right)

	s.Tabs = r.NewStyle().
		Height(1)

	s.TabInactive = r.NewStyle()

	s.TabActive = r.NewStyle().
		Underline(true).
		Foreground(orange)

	s.TabSeparator = r.NewStyle().
		SetString("|").
		Padding(0, 1).
		Foreground(blue)

	s.Spinner = r.NewStyle().
		MarginTop(1).
		MarginLeft(2).
		Foreground(green)

	s.SpinnerContainer = r.NewStyle()

	return s
}
