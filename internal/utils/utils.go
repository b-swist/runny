package utils

import (
	"os"
	"path/filepath"
	"slices"
)

func XdgCurrentDesktop() []string {
	env, ok := os.LookupEnv("XDG_CURRENT_DIR")
	if !ok {
		return nil
	}
	return filepath.SplitList(env)
}

func Intersects[T comparable](a, b []T) bool {
	for _, v := range a {
		if slices.Contains(b, v) {
			return true
		}
	}
	return false
}

func Collect[T any](ch <-chan T) (result []T) {
	for v := range ch {
		result = append(result, v)
	}
	return
}
