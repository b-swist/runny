package app

import (
	"github.com/b-swist/runny/internal/modes"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var appStyle = lipgloss.NewStyle().Padding(1, 2)

type item struct {
	title, desc string
	entry       modes.Entry
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list        list.Model
	chosenEntry modes.Entry
}

func (m *model) ChosenEntry() modes.Entry { return m.chosenEntry }

func newModel[E modes.Entry](entries []E) model {
	items := make([]list.Item, 0, len(entries))
	for _, e := range entries {
		items = append(items, item{
			title: e.DefaultName(),
			desc:  e.Description(),
			entry: e,
		})
	}

	delegate := newItemDelegate(newDelegateKeyMap())
	modelList := list.New(items, delegate, 0, 0)
	modelList.Title = "runny"

	return model{list: modelList}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return appStyle.Render(m.list.View())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case chosenEntryMsg:
		m.chosenEntry = msg
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
