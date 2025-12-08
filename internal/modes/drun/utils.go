package drun

import (
	"slices"
	"strings"

	xdg "github.com/MatthiasKunnen/xdg/desktop"
	"github.com/b-swist/runny/internal/utils"
)

func (e *DrunEntry) isApplication() bool {
	return e.Type == "Application"
}

func (e *DrunEntry) isHidden() bool {
	return e.NoDisplay || e.Hidden
}

func (e *DrunEntry) isExcluded(desktop []string) bool {
	if len(desktop) == 0 {
		return len(e.OnlyShowIn) > 0
	}

	if len(e.OnlyShowIn) > 0 {
		return !utils.Intersects(e.OnlyShowIn, desktop)
	}

	return utils.Intersects(e.NotShowIn, desktop)
}

func loadEntry(path string) (*DrunEntry, error) {
	entry, err := xdg.LoadFile(path)
	if err != nil {
		return nil, err
	}
	return (*DrunEntry)(entry), nil
}

func stripFieldCodes(e xdg.ExecValue) []string {
	return e.ToArguments(xdg.FieldCodeProvider{})
}

func sortEntries(entries []*DrunEntry) {
	slices.SortFunc(entries, func(a, b *DrunEntry) int {
		return strings.Compare(
			strings.ToLower(a.DefaultName()),
			strings.ToLower(b.DefaultName()),
		)
	})
}
