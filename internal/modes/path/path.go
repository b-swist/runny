package path

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/b-swist/runny/internal/utils"
)

type pathEntry struct {
	name string
	path []string
}

func (e *pathEntry) Name() string { return e.name }
func (e *pathEntry) Launch(idx int) error {
	return utils.LaunchTerm([]string{e.path[idx]})
}

func Entries() ([]*pathEntry, error) {
	path := Path()

	var (
		collect = make(map[string][]string)
		errs    = make([]error, 0)
	)

	for _, d := range path {
		if !filepath.IsAbs(d) {
			continue
		}

		entries, err := os.ReadDir(d)
		if err != nil {
			errs = append(errs, fmt.Errorf("error reading dir %s: %w", d, err))
			continue
		}

		for _, e := range entries {
			info, err := e.Info()
			if err != nil {
				errs = append(errs, fmt.Errorf("error reading info of %s: %w", e, err))
				continue
			}

			if info.IsDir() {
				continue
			}
			if !isExecutable(info) {
				continue
			}

			name := info.Name()
			collect[name] = append(
				collect[name],
				filepath.Join(d, name),
			)
		}
	}

	result := make([]*pathEntry, 0, len(collect))
	for k, v := range collect {
		result = append(result, newEntry(k, v))
	}

	sortEntries(result)

	return result, errors.Join(errs...)
}
