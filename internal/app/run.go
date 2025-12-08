package app

import (
	"fmt"

	"github.com/b-swist/runny/internal/modes"
	tea "github.com/charmbracelet/bubbletea"
)

var logFile = "runny.log"

func Run[E modes.Entry](entryFunc func() ([]E, error)) error {
	f, err := tea.LogToFile(logFile, "")
	if err != nil {
		return fmt.Errorf("open log file %s: %w", logFile, err)
	}
	defer f.Close()

	entries, err := entryFunc()
	if err != nil {
		_ = fmt.Errorf("warn: error getting entries: %w", err)
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
