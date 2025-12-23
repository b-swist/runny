package apps

import (
	"errors"
	"fmt"
	"sync"

	xdg "github.com/MatthiasKunnen/xdg/desktop"
	"github.com/b-swist/runny/internal/utils"
)

func allEntries() ([]*xdg.Entry, error) {
	entries, err := xdg.GetDesktopFiles(xdg.GetDesktopFileLocations())
	if err != nil {
		return nil, err
	}

	count := len(entries)
	var (
		resCh = make(chan *xdg.Entry, count)
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

type AppEntry xdg.Entry

func AppEntries() ([]*AppEntry, error) {
	entries, err := allEntries()
	if err != nil {
		return nil, err
	}
	filtered := filterEntries(entries)
	sortEntries(filtered)
	return filtered, nil
}

func filterEntries(entries []*xdg.Entry) []*AppEntry {
	result := make([]*AppEntry, 0, len(entries))
	desktop := utils.XdgCurrentDesktop()

	for _, entry := range entries {

		if !isApplication(entry) {
			continue
		}

		app := (*AppEntry)(entry)

		if app.isHidden() {
			continue
		}
		if app.isExcluded(desktop) {
			continue
		}

		result = append(result, app)
	}

	return result
}

func (e AppEntry) DefaultName() string { return e.Name.Default }
func (e AppEntry) Title() string       { return e.DefaultName() }
func (e AppEntry) FilterValue() string { return e.Title() }

func (e AppEntry) Description() string {
	if s := e.Comment; s.Default != "" {
		return s.Default
	}
	if s := e.GenericName; s.Default != "" {
		return s.Default
	}
	return "No description"
}

func (e AppEntry) Launch() error {
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
