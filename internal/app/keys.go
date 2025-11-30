package app

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	choose key.Binding
}

func newKeyMap() *keyMap {
	return &keyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("↵", "select"),
		),
	}
}
