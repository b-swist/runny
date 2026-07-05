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

func Entries(input string, delim string) ([]*DmenuEntry, error) {
	var options []string

	if input != "" {
		options = strings.Split(strings.TrimSpace(input), delim)
	} else {
		data, err := utils.ReadStdin()
		if err != nil {
			return nil, err
		}
		options = strings.Split(strings.TrimSpace(string(data)), delim)
	}

	entries := make([]*DmenuEntry, 0, len(options))

	for _, e := range options {
		entries = append(entries, &DmenuEntry{e})
	}

	return entries, nil
}
