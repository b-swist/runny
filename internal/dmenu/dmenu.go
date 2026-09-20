package dmenu

import (
	"fmt"
	"strings"

	"github.com/b-swist/runny/internal/utils"
)

type DmenuEntry struct {
	name string
}

func (e *DmenuEntry) Title() string       { return e.name }
func (e *DmenuEntry) FilterValue() string { return e.Title() }
func (e *DmenuEntry) Description() string { return "" }

func (e *DmenuEntry) Action() error {
	_, err := fmt.Println(e.Title())
	return err
}

func Entries(opts ...DmenuOption) ([]*DmenuEntry, error) {
	cfg := &dmenuConfig{delimeter: "\n"}
	for _, opt := range opts {
		opt.Apply(cfg)
	}

	if cfg.input == "" {
		data, err := utils.ReadStdin()
		if err != nil {
			return nil, err
		}
		cfg.input = strings.TrimSpace(string(data))
	}

	split := strings.Split(cfg.input, cfg.delimeter)
	entries := make([]*DmenuEntry, 0, len(split))

	for _, e := range split {
		entries = append(entries, &DmenuEntry{e})
	}

	return entries, nil
}
