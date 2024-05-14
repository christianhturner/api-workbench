package settings

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	// Up used during nav mode
	Up key.Binding
	// Down used during nav mode
	Down key.Binding
	// Left used during nav mode
	Left key.Binding
	// Right used during nav mode
	Right key.Binding
	// Enter used during nav, edit, and create mode
	Enter key.Binding
	// Quite used during nav mode
	Quit key.Binding

	// Create used during create mode
	Create key.Binding
	// Rename used during create mode
	Rename key.Binding
	// Delete used during create mode
	Delete key.Binding
	// Back used during create mode
	Back key.Binding
}

func DefaultKeyMap(stringContext string) *KeyMap {
	km := new(KeyMap)
	km.Up = key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	)

	km.Down = key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	)

	km.Left = key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("←/h", "move left"),
	)

	km.Right = key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("→/l", "move right"),
	)

	km.Enter = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("&#8617;/return", "Enter"),
	)

	km.Quit = key.NewBinding(
		key.WithKeys("q", "ctrl+c", "exc"),
		key.WithHelp("", "quit application"),
	)

	km.Create = key.NewBinding(
		key.WithKeys("c", "create"),
		key.WithHelp("", fmt.Sprintf("create %s", stringContext)),
	)

	km.Rename = key.NewBinding(
		key.WithKeys("r", "rename"),
		key.WithHelp("", fmt.Sprintf("rename %s", stringContext)),
	)

	km.Delete = key.NewBinding(
		key.WithKeys("d", "delete"),
		key.WithHelp("", fmt.Sprintf("delete %s", stringContext)),
	)

	km.Back = key.NewBinding(
		key.WithKeys("b", "back"),
		key.WithHelp("←", "back"),
	)
	return km
}
