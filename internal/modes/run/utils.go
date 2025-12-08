package run

import (
	"os"
	"path/filepath"
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

func newEntry(name, path string) *RunEntry {
	return &RunEntry{
		name: name,
		path: path,
	}
}
