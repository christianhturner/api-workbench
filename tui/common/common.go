package common

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/christianhturner/api-workbench/project"
	"github.com/christianhturner/api-workbench/settings"
	"github.com/christianhturner/api-workbench/tui/styles"
	"github.com/muesli/termenv"
)

var (
	// The current running program
	P *tea.Program
	// Project DB utils shared across TUI
	ProjectDB *project.DB
	// WindowSize store for the terminal window
	WindowSize tea.WindowSizeMsg
)

type Common struct {
	ctx      context.Context
	Col, Row int
	Styles   *styles.Styles
	KeyMap   *settings.KeyMap
	// zones *zone.manager ---- see github.com/lrstanley/bubblezone if mouse support is desired
	Renderer *lipgloss.Renderer
	Output   *termenv.Output
	Logger   *log.Logger
}

func NewCommon(ctx context.Context, out *lipgloss.Renderer, col, row int) Common {
	if ctx == nil {
		ctx = context.TODO()
	}

	return Common{
		ctx:      ctx,
		Col:      col,
		Row:      row,
		Renderer: out,
		Output:   out.Output(),
		Styles:   styles.DefaultStyles(out),
		KeyMap:   settings.DefaultKeyMap(""),
		Logger:   log.FromContext(ctx).WithPrefix("debug"),
	}
}

// Sets value within the context(ctx)
func (c *Common) SetValue(key, value interface{}) {
	c.ctx = context.WithValue(c.ctx, key, value)
}

func (c *Common) SetSize(col, row int) {
	c.Col = col
	c.Row = row
}

func (c *Common) Context() context.Context {
	return c.ctx
}
