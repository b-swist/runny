package app

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

func chosenItem(m *list.Model) tea.Cmd {
	entry, ok := m.SelectedItem().(Item)
	if !ok {
		return nil
	}

	return func() tea.Msg {
		return ChosenItemMsg(entry)
	}
}

func DefaultDelegate() list.DefaultDelegate {
	return newItemDelegate(newDelegateKeyMap())
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
			key.WithHelp("enter", "choose"),
		),
	}
}
