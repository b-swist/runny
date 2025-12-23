package path

import (
	"github.com/b-swist/runny/internal/utils"
)

type item struct {
	entry *pathEntry
	index int
}

func newItem(e *pathEntry) *item {
	return &item{
		entry: e,
		index: 0,
	}
}

func GenerateItems(entries []*pathEntry) []*item {
	result := make([]*item, len(entries))

	for i, v := range entries {
		result[i] = newItem(v)
	}

	return result
}

func (i item) Title() string       { return i.entry.Name() }
func (i item) FilterValue() string { return i.Title() }
func (i item) Description() string { return i.entry.path[i.index] }

func (i item) Launch() error {
	cmd := []string{i.Description()}
	return utils.LaunchTerm(cmd)
}
