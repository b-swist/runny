package drun

import (
	"errors"
	"fmt"

	xdg "github.com/MatthiasKunnen/xdg/desktop"
	"github.com/b-swist/runny/internal/utils"
)

type Entry = xdg.Entry

func DefaultName(e *Entry) string {
	return e.Name.Default
}

func Description(e *Entry) string {
	if s := e.Comment; s.Default != "" {
		return s.Default
	}
	if s := e.GenericName; s.Default != "" {
		return s.Default
	}
	return "No description"
}

func AllEntries() ([]*Entry, error) {
	entries, err := xdg.GetDesktopFiles(xdg.GetDesktopFileLocations())
	if err != nil {
		return nil, err
	}

	result := make([]*Entry, 0, len(entries))
	errs := make([]error, 0, len(entries))

	for _, paths := range entries {
		path := paths[0]
		entry, err := xdg.LoadFile(path)

		if err != nil {
			errs = append(errs, fmt.Errorf("could not parse %s: %w", path, err))
			continue
		}

		result = append(result, entry)
	}

	return result, errors.Join(errs...)
}

func filterEntries(entries []*Entry) []*Entry {
	result := make([]*Entry, 0, len(entries))
	desktop := utils.XdgCurrentDesktop()

	for _, entry := range entries {
		if !isApplication(entry) {
			continue
		}
		if isHidden(entry) {
			continue
		}
		if isExcluded(entry, desktop) {
			continue
		}

		result = append(result, entry)
	}

	return result
}

func ApplicationEntries() ([]*Entry, error) {
	entries, err := AllEntries()
	if err != nil {
		return nil, err
	}
	filtered := filterEntries(entries)
	sortEntries(filtered)
	return filtered, nil
}
