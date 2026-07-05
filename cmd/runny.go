package cmd

import (
	"context"
	"os"

	"github.com/b-swist/runny/internal/app"
	"github.com/b-swist/runny/internal/dmenu"
	"github.com/b-swist/runny/internal/path"
	"github.com/b-swist/runny/internal/xdg"
	"github.com/urfave/cli/v3"
)

var version = "v0.2.0"

func Main() error {
	cmd := &cli.Command{
		Name:    "runny",
		Version: version,
		Usage:   "Application launcher in your terminal",
		Action:  desktopAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "prompt",
				HideDefault: true,
				Usage:       "change default prompt",
				Aliases:     []string{"p"},
			},
		},
		UseShortOptionHandling: true,
		EnableShellCompletion:  true,
		Commands: []*cli.Command{
			{
				Name:    "desktop",
				Usage:   "Uses system .desktop files as input",
				Aliases: []string{"xdg", "drun"},
				Action:  desktopAction,
			},
			{
				Name:    "path",
				Usage:   "Uses executables from $PATH as input",
				Aliases: []string{"run"},
				Action:  pathAction,
			},
			{
				Name:   "dmenu",
				Usage:  "Uses stdin as input",
				Action: dmenuAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Value:   "\n",
						Name:    "delim",
						Usage:   "specify delimeter",
						Aliases: []string{"d"},
					},
					&cli.StringFlag{
						Name:    "input",
						Usage:   "override stdin input",
						Aliases: []string{"i"},
					},
				},
			},
		},
	}

	return cmd.Run(context.Background(), os.Args)
}

func desktopAction(ctx context.Context, cmd *cli.Command) error {
	return defaultAction(cmd, xdg.Entries)
}

func pathAction(ctx context.Context, cmd *cli.Command) error {
	return defaultAction(cmd, path.Entries)
}

func dmenuAction(ctx context.Context, cmd *cli.Command) error {
	return defaultAction(cmd, func() ([]*dmenu.DmenuEntry, error) {
		return dmenu.Entries(cmd.String("input"), cmd.String("delim"))
	}, app.WithDesctiption(false))
}

func defaultAction[I app.Item](cmd *cli.Command, fn func() ([]I, error), opts ...app.ModelOption) error {
	items, err := fn()
	if err != nil {
		return err
	}

	if cmd.String("prompt") != "" {
		opts = append(opts, app.WithPrompt(cmd.String("prompt")))
	}

	cfg := app.NewModelConfig(opts...)

	model := app.NewModel(items, cfg)
	return app.Run(model)
}
