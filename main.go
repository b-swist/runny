package main

import (
	"log"

	"github.com/b-swist/runny/internal/app"
	"github.com/b-swist/runny/internal/modes/drun"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	e, _ := drun.Entries()
	p := tea.NewProgram(app.NewModel(e), tea.WithAltScreen())

	fm, err := p.Run()
	if err != nil {
		log.Fatal("Error running program:", err)
	}

	m, ok := fm.(app.Model)
	if !ok {
		return
	}

	if e := m.ChosenEntry(); e != nil {
		if err := e.Launch(); err != nil {
			log.Fatal(err)
		}
	}
}
