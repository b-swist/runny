package app

import (
	"fmt"
	"log"

	"github.com/b-swist/runny/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func Run(model tea.Model) error {
	logFile, err := utils.LogPath()
	if err != nil {
		return err
	}

	f, err := tea.LogToFile(logFile, "")
	if err != nil {
		return err
	}
	defer f.Close()

	p := tea.NewProgram(model, tea.WithAltScreen())

	fm, err := p.Run()
	if err != nil {
		return fmt.Errorf("error running program: %w", err)
	}

	m, ok := fm.(Model)
	if !ok {
		return fmt.Errorf("unexpected final model type: %T", fm)
	}

	log.Println("debug:", "model implements the interface")

	if e := m.ChosenEntry(); e != nil {
		if err := e.Launch(); err != nil {
			return fmt.Errorf("failed to run entry: %w", err)
		}
	}

	return nil
}
