package main

import (
	"log"

	"github.com/b-swist/runny/internal/app"
	"github.com/b-swist/runny/internal/launcher"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(app.NewModel(), tea.WithAltScreen())

	fm, err := p.Run()
	if err != nil {
		log.Fatal("Error running program:", err)
	}

	m, ok := fm.(app.Model)
	if !ok {
		return
	}

	if m.ChosenEntry() != nil {
		if err := launcher.Launch(m.ChosenEntry()); err != nil {
			log.Fatal(err)
		}
	}
}
