package app

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	list   list.Model
	chosen Item
}

type Item interface {
	list.Item
	Action() error
}

func NewModel[I Item](items []I, cfg *modelConfig) *Model {
	modelItems := make([]list.Item, 0, len(items))
	for _, i := range items {
		modelItems = append(modelItems, i)
	}

	delegate := DefaultDelegate(cfg.showDescription)
	modelList := list.New(modelItems, delegate, 0, 0)
	modelList.Title = cfg.prompt

	return &Model{list: modelList}
}

func (m Model) ChosenEntry() Item { return Item(m.chosen) }
func (m Model) Init() tea.Cmd     { return func() tea.Msg { return FocusFilterMsg{} } }

func (m Model) View() tea.View {
	r := AppStyle.Render(m.list.View())
	v := tea.NewView(r)
	v.AltScreen = true
	return v
}

func (m *Model) focusFilter() {
	m.list.SetFilterText("")
	m.list.SetFilterState(list.Filtering)
}

type ChosenItemMsg = Item
type FocusFilterMsg struct{}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := AppStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case ChosenItemMsg:
		m.chosen = msg
		return m, tea.Quit

	case FocusFilterMsg:
		m.focusFilter()
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
