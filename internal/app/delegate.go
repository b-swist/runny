package app

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type chosenItemMsg struct {
	item item
}

func chosenItem(m *list.Model) tea.Cmd {
	it, ok := m.SelectedItem().(item)
	if !ok {
		return func() tea.Msg { return nil }
	}
	return func() tea.Msg {
		return chosenItemMsg{item: it}
	}
}

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	d.ShortHelpFunc = keys.ShortHelp
	d.FullHelpFunc = keys.FullHelp

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				return chosenItem(m)
			}
		}
		return nil
	}

	return d
}

type delegateKeyMap struct {
	choose key.Binding
}

func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{d.choose}
}

func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{d.choose}}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("↵", "choose"),
		),
	}
}
