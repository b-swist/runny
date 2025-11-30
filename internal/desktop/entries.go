package desktop

import (
	xdg "github.com/MatthiasKunnen/xdg/desktop"
	"log"
)

type Entry = xdg.Entry

func getAllEntries() ([]*Entry, error) {
	entries, err := xdg.GetDesktopFiles(xdg.GetDesktopFileLocations())
	if err != nil {
		return nil, err
	}

	result := make([]*Entry, 0, len(entries))
	for _, paths := range entries {
		path := paths[0]
		entry, err := xdg.LoadFile(path)

		if err != nil {
			log.Printf("Could not parse %v: %v", path, err)
			continue
		}

		result = append(result, entry)
	}

	return result, nil
}

func filterAppEntries(entries []*Entry) []*Entry {
	result := make([]*Entry, 0, len(entries))
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

func GetAppEntries() ([]*Entry, error) {
	entries, err := getAllEntries()
	if err != nil {
		return nil, err
	}
	filtered := filterAppEntries(entries)
	sortEntries(filtered)
	return filtered, nil
}

func GetFinalExec(e *Entry) []string {
	return e.Exec.ToArguments(xdg.FieldCodeProvider{})
}

func GetDefaultName(e *Entry) string {
	return e.Name.Default
}

func GetDescription(e *Entry) string {
	if s := e.Comment; s.Default != "" {
		return s.Default
	}
	if s := e.GenericName; s.Default != "" {
		return s.Default
	}
	return "No description"
}
