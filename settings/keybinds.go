package settings

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Enter key.Binding
	Quit  key.Binding
}

func DefaultKeyMap() *KeyMap {
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

	return km
}
