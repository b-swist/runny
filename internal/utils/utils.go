package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
)

func ReadStdin() ([]byte, error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("error reading from stdin: %w", err)
	}

	return data, nil
}

func XDGCurrentDesktop() []string {
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
