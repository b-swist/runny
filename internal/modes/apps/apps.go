package apps

import (
	"errors"
	"fmt"
	"sync"

	xdg "github.com/MatthiasKunnen/xdg/desktop"
	"github.com/b-swist/runny/internal/utils"
)

type AppEntry xdg.Entry

func (e *AppEntry) DefaultName() string { return e.Name.Default }

func (e *AppEntry) Description() string {
	if s := e.Comment; s.Default != "" {
		return s.Default
	}
	if s := e.GenericName; s.Default != "" {
		return s.Default
	}
	return "No description"
}

func (e *AppEntry) Launch() error {
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

func Entries() ([]*AppEntry, error) { return appEntries() }

func allEntries() ([]*AppEntry, error) {
	entries, err := xdg.GetDesktopFiles(xdg.GetDesktopFileLocations())
	if err != nil {
		return nil, err
	}

	count := len(entries)
	var (
		resCh = make(chan *AppEntry, count)
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

func filterEntries(entries []*AppEntry) []*AppEntry {
	result := make([]*AppEntry, 0, len(entries))
	desktop := utils.XdgCurrentDesktop()

	for _, entry := range entries {

		if !entry.isApplication() {
			continue
		}
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

func appEntries() ([]*AppEntry, error) {
	entries, err := allEntries()
	if err != nil {
		return nil, err
	}
	filtered := filterEntries(entries)
	sortEntries(filtered)
	return filtered, nil
}
