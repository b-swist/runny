package utils

import (
	"os"
	"os/exec"
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

func LogPath() (string, error) {
	dir := os.Getenv("XDG_DATA_HOME")

	if dir == "" || !filepath.IsAbs(dir) {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		dir = filepath.Join(home, ".share", "local")
	}

	if err := os.MkdirAll(dir, 0o775); err != nil {
		return "", err
	}

	return filepath.Join(dir, "runny.log"), nil
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

func FullPath(cmd string) (string, error) {
	if filepath.IsAbs(cmd) {
		return cmd, nil
	}

	path, err := exec.LookPath(cmd)
	if err != nil {
		return "", err
	}

	return path, nil
}
