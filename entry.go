package main

import (
	"log"
	"slices"
	"strings"

	"github.com/MatthiasKunnen/xdg/desktop"
)

func getEntries() ([]*desktop.Entry, error) {
	entries, err := desktop.GetDesktopFiles(desktop.GetDesktopFileLocations())
	if err != nil {
		return nil, err
	}

	result := make([]*desktop.Entry, 0, len(entries))
	for _, paths := range entries {
		path := paths[0]
		entry, err := desktop.LoadFile(path)

		if err != nil {
			log.Printf("Could not parse %v: %v", path, err)
			continue
		}

		result = append(result, entry)
	}

	return result, nil
}

func filterAppEntries(entries []*desktop.Entry) []*desktop.Entry {
	result := make([]*desktop.Entry, 0, len(entries))
	xdgCurrentDesktop := getXdgCurrentDesktop()

	for _, entry := range entries {
		if entry.Type != "Application" {
			continue
		}

		if entry.NoDisplay || entry.Hidden {
			continue
		}

		if xdgCurrentDesktop != nil {
			if len(entry.OnlyShowIn) != 0 && intersects(xdgCurrentDesktop, entry.OnlyShowIn) {
				continue
			} else if intersects(xdgCurrentDesktop, entry.NotShowIn) {
				continue
			}
		}

		result = append(result, entry)
	}

	return result
}

func getAppEntries() ([]*desktop.Entry, error) {
	entries, err := getEntries()
	if err != nil {
		return nil, err
	}
	filtered := filterAppEntries(entries)
	sortEntries(filtered)
	return filtered, nil
}

func getFinalExec(e *desktop.Entry) []string {
	return e.Exec.ToArguments(desktop.FieldCodeProvider{})
}

func getDefaultName(e *desktop.Entry) string {
	return e.Name.Default
}

func getDescription(e *desktop.Entry) string {
	if s := e.Comment; s.Default != "" {
		return s.Default
	}
	if s := e.GenericName; s.Default != "" {
		return s.Default
	}
	return "No description"
}

func sortEntries(entries []*desktop.Entry) {
	slices.SortFunc(entries, func(a, b *desktop.Entry) int {
		return strings.Compare(
			strings.ToLower(getDefaultName(a)),
			strings.ToLower(getDefaultName(b)),
		)
	})
}
