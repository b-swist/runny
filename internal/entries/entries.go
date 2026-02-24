package entries

import (
	"errors"
	"fmt"

	xdg "github.com/MatthiasKunnen/xdg/desktop"
	"github.com/b-swist/runny/internal/utils"
)

func allEntries() ([]*xdg.Entry, error) {
	entries, err := xdg.GetDesktopFiles(xdg.GetDesktopFileLocations())
	if err != nil {
		return nil, err
	}

	var (
		result = make([]*xdg.Entry, 0, len(entries))
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

type Entry xdg.Entry

func AppEntries() ([]*Entry, error) {
	entries, err := allEntries()
	if err != nil {
		return nil, err
	}
	filtered := filterEntries(entries)
	sortEntries(filtered)
	return filtered, nil
}

func filterEntries(entries []*xdg.Entry) []*Entry {
	result := make([]*Entry, 0, len(entries))
	desktop := utils.XdgCurrentDesktop()

	for _, e := range entries {

		if !isApplication(e) {
			continue
		}

		entry := (*Entry)(e)

		if entry.isHidden() {
			continue
		}
		if entry.isExcluded(desktop) {
			continue
		}

		result = append(result, entry)
	}

	return result
}

func (e Entry) DefaultName() string { return e.Name.Default }
func (e Entry) Title() string       { return e.DefaultName() }
func (e Entry) FilterValue() string { return e.Title() }

func (e Entry) Description() string {
	if s := e.Comment; s.Default != "" {
		return s.Default
	}
	if s := e.GenericName; s.Default != "" {
		return s.Default
	}
	return "No description"
}

func (e Entry) Launch() error {
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
