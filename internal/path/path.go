package path

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/b-swist/runny/internal/utils"
)

type PathEntry struct {
	name string
	path string
}

func (e *PathEntry) Title() string       { return e.name }
func (e *PathEntry) FilterValue() string { return e.Title() }
func (e *PathEntry) Description() string { return e.path }
func (e *PathEntry) Action() error       { return e.launch() }

func Entries() ([]*PathEntry, error) {
	p, err := path()

	if err != nil {
		return nil, err
	}

	var (
		result = make([]*PathEntry, 0, 2048)
		errs   = make([]error, 0)
		added  = make(map[string]struct{})
	)

	for _, dir := range p {
		if !filepath.IsAbs(dir) {
			continue
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			errs = append(errs, fmt.Errorf("error reading dir %s: %w", dir, err))
		}

		for _, e := range entries {
			if e.IsDir() {
				continue
			}

			info, err := e.Info()
			if err != nil {
				errs = append(errs, fmt.Errorf("error reading info of %s: %w", e, err))
				continue
			}

			if !isExecutable(info) {
				continue
			}

			name := info.Name()

			if _, exists := added[name]; exists {
				continue
			}
			added[name] = struct{}{}

			result = append(result, &PathEntry{
				name,
				filepath.Join(dir, name),
			})
		}
	}

	return result, nil
}

func path() ([]string, error) {
	p, ok := os.LookupEnv("PATH")
	if !ok {
		return nil, fmt.Errorf("PATH not set")
	}

	result := filepath.SplitList(p)
	for i, d := range result {
		result[i] = filepath.Clean(os.ExpandEnv(d))
	}

	return result, nil
}

func (e *PathEntry) launch() error { return utils.LaunchTerm([]string{e.path}) }

func isExecutable(f os.FileInfo) bool {
	return f.Mode()&0111 != 0
}
