package app

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/b-swist/runny/internal/utils"
)

func Run(model tea.Model) error {
	logFile, err := utils.LogFile()
	if err != nil {
		return err
	}

	f, err := tea.LogToFile(logFile, "")
	if err != nil {
		return err
	}
	defer f.Close()

	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return fmt.Errorf("failed to open controlling terminal: %w", err)
	}
	defer tty.Close()

	p := tea.NewProgram(
		model,
		tea.WithOutput(tty),
	)

	fm, err := p.Run()
	if err != nil {
		return fmt.Errorf("error running program: %w", err)
	}

	m, ok := fm.(Model)
	if !ok {
		return fmt.Errorf("unexpected final model type: %T", fm)
	}

	if e := m.ChosenEntry(); e != nil {
		if err := e.Action(); err != nil {
			return fmt.Errorf("failed to run entry: %w", err)
		}
	}

	return nil
}
