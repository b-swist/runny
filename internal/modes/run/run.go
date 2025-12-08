package run

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/b-swist/runny/internal/launcher"
	"github.com/b-swist/runny/internal/utils"
)

type RunEntry struct {
	name, path string
}

func (e *RunEntry) DefaultName() string { return e.name }
func (e *RunEntry) Description() string { return e.path }
func (e *RunEntry) Launch() error       { return launcher.LaunchTerm([]string{e.path}) }

func Entries() ([]*RunEntry, error) {
	path := Path()

	var (
		resCh = make(chan *RunEntry, len(path))
		errCh = make(chan error, len(path))
		wg    sync.WaitGroup
	)

	for _, d := range path {
		if !filepath.IsAbs(d) {
			continue
		}

		entries, err := os.ReadDir(d)
		if err != nil {
			errCh <- fmt.Errorf("error reading dir %s: %w", d, err)
			continue
		}

		for _, entry := range entries {
			wg.Add(1)
			go func(e os.DirEntry) {
				defer wg.Done()

				info, err := e.Info()
				if err != nil {
					errCh <- fmt.Errorf("error reading info of %s: %w", e, err)
					return
				}

				if info.IsDir() {
					return
				}
				if !isExecutable(info) {
					return
				}

				resCh <- newEntry(
					info.Name(),
					filepath.Join(d, info.Name()),
				)
			}(entry)
		}
	}

	go func() {
		wg.Wait()
		close(resCh)
		close(errCh)
	}()

	var (
		result = utils.Collect(resCh)
		errs   = utils.Collect(errCh)
	)

	return result, errors.Join(errs...)
}
