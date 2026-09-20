package cmd

import (
	"context"
	"os"

	"github.com/b-swist/runny/internal/app"
	"github.com/b-swist/runny/internal/dmenu"
	"github.com/b-swist/runny/internal/path"
	"github.com/b-swist/runny/internal/utils"
	"github.com/b-swist/runny/internal/xdg"
	"github.com/urfave/cli/v3"
)

const version = "v0.3.2"

func Main() error {
	cmd := &cli.Command{
		Name:    "runny",
		Version: version,
		Usage:   "Application launcher in your terminal",
		Action:  defaultAction,
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

func defaultAction(ctx context.Context, cmd *cli.Command) error {
	if utils.IsInteractive() {
		return desktopAction(ctx, cmd)
	} else {
		return dmenuAction(ctx, cmd)
	}
}

func desktopAction(ctx context.Context, cmd *cli.Command) error {
	return action(cmd, xdg.Entries)
}

func pathAction(ctx context.Context, cmd *cli.Command) error {
	return action(cmd, path.Entries)
}

func dmenuAction(ctx context.Context, cmd *cli.Command) error {
	var opts []dmenu.DmenuOption

	if d := cmd.String("delim"); d != "" {
		opts = append(opts, dmenu.WithDelimeter(d))
	}
	if i := cmd.String("input"); i != "" {
		opts = append(opts, dmenu.WithInput(i))
	}

	return action(cmd, func() ([]*dmenu.DmenuEntry, error) {
		return dmenu.Entries(opts...)
	}, app.WithDesctiption(false))
}

func action[I app.Item](cmd *cli.Command, fn func() ([]I, error), opts ...app.ModelOption) error {
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
