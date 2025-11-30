package main

import (
	"github.com/MatthiasKunnen/xdg/desktop"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

type item struct {
	title, desc string
	entry       *desktop.Entry
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

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

type model struct {
	list     list.Model
	keys     *keyMap
	selected *desktop.Entry
}

func newModel() (*model, error) {
	entries, err := getAppEntries()
	if err != nil {
		return nil, err
	}

	items := make([]list.Item, 0, len(entries))
	for _, e := range entries {
		items = append(items, item{
			title: getDefaultName(e),
			desc:  getDescription(e),
			entry: e,
		})
	}

	m := &model{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
		keys: newKeyMap(),
	}
	m.list.Title = "runny"

	return m, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return m.list.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func main() {
	m, err := newModel()
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	final, err := p.Run()
	if err != nil {
		log.Fatal("Error running program:", err)
	}

	fm, ok := final.(model)
	if !ok {
		log.Fatal(err)
	}

	if fm.selected != nil {
		runApp(fm.selected)
	}
}
