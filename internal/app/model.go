package app

import (
	drun "github.com/b-swist/runny/internal/drun"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var appStyle = lipgloss.NewStyle().Padding(1, 2)

type item struct {
	title, desc string
	entry       *drun.Entry
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	list        list.Model
	chosenEntry *drun.Entry
}

func (m *Model) ChosenEntry() *drun.Entry { return m.chosenEntry }

func NewModel() Model {

	entries, err := drun.ApplicationEntries()
	if err != nil {
		panic(err)
	}

	items := make([]list.Item, 0, len(entries))
	for _, e := range entries {
		items = append(items, item{
			title: drun.DefaultName(e),
			desc:  drun.Description(e),
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
	return appStyle.Render(m.list.View())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
