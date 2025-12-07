package drun

import (
	"errors"
	"fmt"
	"sync"

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

	count := len(entries)
	var (
		resCh = make(chan *Entry, count)
		errCh = make(chan error, count)
		wg    sync.WaitGroup
	)

	for name, paths := range entries {
		if len(paths) == 0 {
			errCh <- fmt.Errorf("no path associated with %s", name)
			continue
		}
		path := paths[0]

		wg.Add(1)
		go func(p string) {
			defer wg.Done()

			entry, err := loadEntry(p)
			if err != nil {
				errCh <- fmt.Errorf("could not load %s: %w", p, err)
				return
			}

			resCh <- entry
		}(path)
	}

	go func() {
		wg.Wait()
		close(resCh)
		close(errCh)
	}()

	result := utils.Collect(resCh)
	errs := utils.Collect(errCh)

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
