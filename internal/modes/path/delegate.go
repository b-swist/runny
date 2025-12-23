package path

import (
	"github.com/b-swist/runny/internal/app"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func chosenItem(m *list.Model) tea.Cmd {
	it, ok := m.SelectedItem().(*item)
	if !ok {
		return nil
	}

	return func() tea.Msg {
		return app.ChosenItemMsg(it)
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
			if key.Matches(msg, keys.choose) {
				return chosenItem(m)
			}

			it, ok := m.SelectedItem().(*item)
			if !ok {
				return nil
			}

			n := len(it.entry.path)
			if n == 0 {
				return nil
			}

			switch {
			case key.Matches(msg, keys.nextPath):
				if it.index < n-1 {
					it.index++
				} else {
					it.index = 0
				}
				m.SetItem(m.GlobalIndex(), it)

			case key.Matches(msg, keys.prevPath):
				if it.index > 0 {
					it.index--
				} else {
					it.index = n - 1
				}
				m.SetItem(m.GlobalIndex(), it)
			}
		}
		return nil
	}

	return d
}

type delegateKeyMap struct {
	choose   key.Binding
	nextPath key.Binding
	prevPath key.Binding
}

func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{d.choose}
}

func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{d.choose, d.nextPath, d.prevPath}}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
		nextPath: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "next path"),
		),
		prevPath: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("shift+tab", "prev path"),
		),
	}
}
