package main

import (
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
			return nil, err
		}
		result = append(result, entry)
	}

	return result, nil
}
