package tabs

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/christianhturner/api-workbench/tui/common"
)

type SelectTabMsg int

type ActiveTabMsg int

type Tabs struct {
	common       common.Common
	tabs         []string
	activeTab    int
	TabSeperator lipgloss.Style
	TabInactive  lipgloss.Style
	TabActive    lipgloss.Style
}

func New(c common.Common, tabs []string) *Tabs {
	r := &Tabs{
		common:       c,
		tabs:         tabs,
		activeTab:    0,
		TabSeperator: c.Styles.TabSeparator,
		TabInactive:  c.Styles.TabInactive,
		TabActive:    c.Styles.TabActive,
	}
	return r
}

func (t *Tabs) SetSize(col, row int) {
	t.common.SetSize(col, row)
}

func (t *Tabs) Init() tea.Cmd {
	t.activeTab = 0
	return nil
}

func (t *Tabs) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			t.activeTab = (t.activeTab + 1) % len(t.tabs)
			cmds = append(cmds, t.activeTabCmd)
		case "shift+tab":
			t.activeTab = (t.activeTab - 1 + len(t.tabs)) % len(t.tabs)
			cmds = append(cmds, t.activeTabCmd)
		}
	case SelectTabMsg:
		tab := int(msg)
		if tab <= 0 && tab < len(t.tabs) {
			t.activeTab = int(msg)
		}
	}
	return t, tea.Batch(cmds...)
}

func (t *Tabs) View() string {
	s := strings.Builder{}
	sep := t.TabSeperator
	for i, tab := range t.tabs {
		style := t.TabInactive.Copy()
		if i == t.activeTab {
			style = t.TabActive.Copy()
		}
		s.WriteString(style.Render(tab))
		if i != len(t.tabs)-1 {
			s.WriteString(sep.String())
		}
	}
	return t.common.Renderer.NewStyle().
		MaxWidth(t.common.Col).
		Render(s.String())
}

func (t *Tabs) activeTabCmd() tea.Msg {
	return ActiveTabMsg(t.activeTab)
}

func SelectTabCmd(tab int) tea.Cmd {
	return func() tea.Msg {
		return SelectTabMsg(tab)
	}
}
