package xdg

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/MatthiasKunnen/xdg/desktop"
	"github.com/b-swist/runny/internal/utils"
)

type desktopEntry desktop.Entry

func (e *desktopEntry) FilterValue() string { return e.Title() }
func (e *desktopEntry) Title() string       { return e.Name.Default }
func (e *desktopEntry) Action() error       { return e.launch() }

func (e desktopEntry) Description() string {
	if e.Comment.Default != "" {
		return e.Comment.Default
	}
	if e.GenericName.Default != "" {
		return e.GenericName.Default
	}
	return "No description"
}

func Entries() ([]*desktopEntry, error) {
	entries, err := loadAllEntries()
	if err != nil {
		return nil, err
	}
	filtered := filterVisibleEntries(entries)
	sortEntriesByName(filtered)
	return filtered, nil
}

func loadEntry(path string) (*desktopEntry, error) {
	entry, err := desktop.LoadFile(path)
	if err != nil {
		return nil, err
	}
	return (*desktopEntry)(entry), nil
}

func loadAllEntries() ([]*desktopEntry, error) {
	entries, err := desktop.GetDesktopFiles(desktop.GetDesktopFileLocations())
	if err != nil {
		return nil, err
	}

	var (
		result = make([]*desktopEntry, 0, len(entries))
		errs   = make([]error, 0)
	)

	for name, paths := range entries {
		if len(paths) == 0 {
			errs = append(errs, fmt.Errorf("no path associated with %s", name))
			continue
		}
		p := paths[0]

		entry, err := loadEntry(p)
		if err != nil {
			errs = append(errs, fmt.Errorf("could not load %s: %w", p, err))
			continue
		}

		result = append(result, entry)
	}

	return result, errors.Join(errs...)
}

func filterVisibleEntries(entries []*desktopEntry) []*desktopEntry {
	result := make([]*desktopEntry, 0, len(entries))
	desktop := utils.XDGCurrentDesktop()

	for _, e := range entries {

		if !e.isApplication() {
			continue
		}

		if e.isHidden() {
			continue
		}
		if e.isExcluded(desktop) {
			continue
		}

		result = append(result, e)
	}

	return result
}

func (e *desktopEntry) launch() error {
	cmd := stripFieldCodes(e.Exec)
	if e.Terminal {
		if err := utils.LaunchTerm(cmd); err != nil {
			return err
		}
	} else {
		if err := utils.LaunchGui(cmd); err != nil {
			return err
		}
	}

	return nil
}

func (e *desktopEntry) isApplication() bool {
	return e.Type == "Application"
}

func (e *desktopEntry) isHidden() bool {
	return e.NoDisplay || e.Hidden
}

func (e *desktopEntry) isExcluded(desktop []string) bool {
	if len(desktop) == 0 {
		return len(e.OnlyShowIn) > 0
	}

	if len(e.OnlyShowIn) > 0 {
		return !utils.Intersects(e.OnlyShowIn, desktop)
	}

	return utils.Intersects(e.NotShowIn, desktop)
}

func stripFieldCodes(e desktop.ExecValue) []string {
	return e.ToArguments(desktop.FieldCodeProvider{})
}

func sortEntriesByName(entries []*desktopEntry) {
	slices.SortFunc(entries, func(a, b *desktopEntry) int {
		return strings.Compare(
			strings.ToLower(a.Title()),
			strings.ToLower(b.Title()),
		)
	})
}
