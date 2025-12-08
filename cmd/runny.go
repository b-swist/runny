package cmd

import (
	"errors"
	"flag"
	"fmt"

	"github.com/b-swist/runny/internal/app"
	"github.com/b-swist/runny/internal/modes/drun"
	"github.com/b-swist/runny/internal/modes/run"
)

var (
	help = flag.Bool("h", false, "Show help message")
	mode = flag.String("m", "", "Specify a mode to launch")
)

const (
	ModeRun  = "run"
	ModeDrun = "drun"
)

var (
	ErrTooManyArgs = errors.New("too many arguments provided")
	ErrNoArgs      = errors.New("no arguments provided")
	ErrInvalidMode = errors.New("invalid mode provided")
)

func Main() error {
	flag.Parse()

	switch {
	case flag.NArg() > 0, flag.NFlag() > 1:
		return ErrTooManyArgs
	case flag.NFlag() == 0, *help:
		flag.PrintDefaults()
	case *mode != "":
		return handleMode(*mode)
	}

	return nil
}

func handleMode(s string) error {
	switch s {
	case "run":
		return app.Run(run.Entries)
	case "drun":
		return app.Run(drun.Entries)
	default:
		return fmt.Errorf("%w: %s", ErrInvalidMode, s)
	}
}
