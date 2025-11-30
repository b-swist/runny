package main

import (
	"log"

	"github.com/b-swist/runny/internal/app"
	"github.com/b-swist/runny/internal/launcher"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m, err := app.NewModel()
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	final, err := p.Run()
	if err != nil {
		log.Fatal("Error running program:", err)
	}

	fm, ok := final.(app.Model)
	if !ok {
		log.Fatal(err)
	}

	if fm.Selected() != nil {
		launcher.Launch(fm.Selected())
	}
}
