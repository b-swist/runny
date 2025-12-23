package apps

import (
	"slices"
	"strings"

	xdg "github.com/MatthiasKunnen/xdg/desktop"
	"github.com/b-swist/runny/internal/utils"
)

func isApplication(e *xdg.Entry) bool {
	return e.Type == "Application"
}

func (e *AppEntry) isHidden() bool {
	return e.NoDisplay || e.Hidden
}

func (e *AppEntry) isExcluded(desktop []string) bool {
	if len(desktop) == 0 {
		return len(e.OnlyShowIn) > 0
	}

	if len(e.OnlyShowIn) > 0 {
		return !utils.Intersects(e.OnlyShowIn, desktop)
	}

	return utils.Intersects(e.NotShowIn, desktop)
}

func loadEntry(path string) (*xdg.Entry, error) {
	entry, err := xdg.LoadFile(path)
	if err != nil {
		return nil, err
	}
	return (*xdg.Entry)(entry), nil
}

func stripFieldCodes(e xdg.ExecValue) []string {
	return e.ToArguments(xdg.FieldCodeProvider{})
}

func sortEntries(entries []*AppEntry) {
	slices.SortFunc(entries, func(a, b *AppEntry) int {
		return strings.Compare(
			strings.ToLower(a.DefaultName()),
			strings.ToLower(b.DefaultName()),
		)
	})
}
