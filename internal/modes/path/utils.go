package path

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func isExecutable(f os.FileInfo) bool {
	return f.Mode()&0111 != 0
}

func Path() []string {
	path, ok := os.LookupEnv("PATH")
	if !ok {
		return nil
	}

	result := filepath.SplitList(path)
	for i, d := range result {
		result[i] = filepath.Clean(os.ExpandEnv(d))
	}

	return result
}

func newEntry(name string, path []string) *pathEntry {
	return &pathEntry{
		name: name,
		path: path,
	}
}

func sortEntries(entries []*pathEntry) {
	slices.SortFunc(entries, func(a, b *pathEntry) int {
		return strings.Compare(
			strings.ToLower(a.Name()),
			strings.ToLower(b.Name()),
		)
	})
}
