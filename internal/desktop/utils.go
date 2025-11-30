package desktop

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func sortEntries(entries []*Entry) {
	slices.SortFunc(entries, func(a, b *Entry) int {
		return strings.Compare(
			strings.ToLower(GetDefaultName(a)),
			strings.ToLower(GetDefaultName(b)),
		)
	})
}

func getXdgCurrentDesktop() []string {
	env, ok := os.LookupEnv("XDG_CURRENT_DIR")
	if !ok {
		return nil
	}
	return filepath.SplitList(env)
}

func intersects(a, b []string) bool {
	for _, v := range a {
		if slices.Contains(b, v) {
			return true
		}
	}
	return false
}
