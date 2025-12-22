package app

import (
	"fmt"
	"log"

	"github.com/b-swist/runny/internal/modes"
	"github.com/b-swist/runny/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func Run[E modes.Entry](entryFunc func() ([]E, error)) error {
	logFile, err := utils.LogPath()
	if err != nil {
		return err
	}

	f, err := tea.LogToFile(logFile, "")
	if err != nil {
		return err
	}
	defer f.Close()

	entries, err := entryFunc()
	if err != nil {
		log.Printf("warn: error getting entries: %v", err)
	}

	p := tea.NewProgram(newModel(entries), tea.WithAltScreen())

	fm, err := p.Run()
	if err != nil {
		return fmt.Errorf("error running program: %w", err)
	}

	m, ok := fm.(model)
	if !ok {
		return fmt.Errorf("unexpected final model type: %T", fm)
	}

	if e := m.ChosenEntry(); e != nil {
		if err := e.Launch(); err != nil {
			return fmt.Errorf("failed to run entry: %w", err)
		}
	}

	return nil
}
