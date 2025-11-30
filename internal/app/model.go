package app

import (
	"github.com/b-swist/runny/internal/desktop"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title, desc string
	entry       *desktop.Entry
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	list     list.Model
	keys     *keyMap
	selected *desktop.Entry
}

func (m Model) Selected() *desktop.Entry { return m.selected }

func NewModel() (Model, error) {
	entries, err := desktop.GetAppEntries()
	if err != nil {
		return Model{}, err
	}

	items := make([]list.Item, 0, len(entries))
	for _, e := range entries {
		items = append(items, item{
			title: desktop.GetDefaultName(e),
			desc:  desktop.GetDescription(e),
			entry: e,
		})
	}

	m := Model{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
		keys: newKeyMap(),
	}
	m.list.Title = "runny"

	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	return m.list.View()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.choose):
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.selected = i.entry
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
