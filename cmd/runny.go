package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/b-swist/runny/internal/app"
	"github.com/b-swist/runny/internal/modes/apps"
	"github.com/b-swist/runny/internal/modes/path"
)

var version = "v0.1.1"

var (
	help = flag.Bool("h", false, "Show help message")
	ver  = flag.Bool("v", false, "Show program version")
	mode = flag.String("m", "", "Specify a mode to launch")
)

const (
	ModePath = "path"
	ModeApps = "apps"
)

var (
	ErrTooManyArgs = errors.New("too many arguments provided")
	ErrInvalidMode = errors.New("invalid mode provided")
)

func Main() error {
	flag.Parse()

	switch {
	case flag.NArg() > 0, flag.NFlag() > 1:
		return ErrTooManyArgs
	case flag.NFlag() == 0, *help:
		flag.PrintDefaults()
	case *ver:
		fmt.Printf("%v: %v\n", os.Args[0], version)
	case *mode != "":
		return handleMode(*mode)
	}

	return nil
}

func handleMode(mode string) error {
	switch mode {
	case ModePath:
		entries, err := path.Entries()
		if err != nil {
			return err
		}
		items := path.GenerateItems(entries)
		model := app.NewModel(items, path.DefaultDelegate())
		return app.Run(model)

	case ModeApps:
		items, err := apps.AppEntries()
		if err != nil {
			return err
		}
		model := app.NewModel(items, apps.DefaultDelegate())
		return app.Run(model)

	default:
		return fmt.Errorf("%w: %s", ErrInvalidMode, mode)
	}
}
