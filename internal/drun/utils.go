package drun

import (
	"slices"
	"strings"

	xdg "github.com/MatthiasKunnen/xdg/desktop"
	"github.com/b-swist/runny/internal/utils"
)

func isApplication(e *Entry) bool {
	return e.Type == "Application"
}

func isHidden(e *Entry) bool {
	return e.NoDisplay || e.Hidden
}

func isExcluded(e *Entry, desktop []string) bool {
	if len(desktop) == 0 {
		return len(e.OnlyShowIn) > 0
	}

	if len(e.OnlyShowIn) > 0 {
		return !utils.Intersects(e.OnlyShowIn, desktop)
	}

	return utils.Intersects(e.NotShowIn, desktop)
}

func loadEntry(path string) (*Entry, error) {
	entry, err := xdg.LoadFile(path)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func StripFieldCodes(e xdg.ExecValue) []string {
	return e.ToArguments(xdg.FieldCodeProvider{})
}

func sortEntries(entries []*Entry) {
	slices.SortFunc(entries, func(a, b *Entry) int {
		return strings.Compare(
			strings.ToLower(DefaultName(a)),
			strings.ToLower(DefaultName(b)),
		)
	})
}
