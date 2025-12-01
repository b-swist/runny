package app

import (
	"github.com/b-swist/runny/internal/desktop"
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
	list   list.Model
	chosen item
}

func (m *Model) ChosenEntry() *desktop.Entry { return m.chosen.entry }

func NewModel() Model {

	entries, err := desktop.GetAppEntries()
	if err != nil {
		panic(err)
	}

	items := make([]list.Item, 0, len(entries))
	for _, e := range entries {
		items = append(items, item{
			title: desktop.GetDefaultName(e),
			desc:  desktop.GetDescription(e),
			entry: e,
		})
	}

	delegate := newItemDelegate(newDelegateKeyMap())
	modelList := list.New(items, delegate, 0, 0)
	modelList.Title = "runny"

	return Model{list: modelList}
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
	case chosenItemMsg:
		m.chosen = msg.item
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
