package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/b-swist/runny/internal/app"
	"github.com/b-swist/runny/internal/entries"
)

var version = "v0.2.0"

var (
	help = flag.Bool("h", false, "Show help message")
	ver  = flag.Bool("v", false, "Show program version")
)

var ErrTooManyArgs = errors.New("too many arguments provided")

func Main() error {
	flag.Parse()

	switch {
	case flag.NArg() > 0, flag.NFlag() > 1:
		return ErrTooManyArgs
	case *help:
		flag.PrintDefaults()
	case *ver:
		fmt.Printf("%v: %v\n", os.Args[0], version)
	}

	items, err := entries.AppEntries()
	if err != nil {
		return err
	}
	model := app.NewModel(items, app.DefaultDelegate())
	return app.Run(model)
}
