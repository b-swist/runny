package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
)

func LogFile() (string, error) {
	dir, err := XDGDataHome()
	if err != nil {
		return "", err
	}

	f := filepath.Join(dir, "runny", "runny.log")

	return f, nil
}

func ReadStdin() ([]byte, error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("error reading from stdin: %w", err)
	}

	return data, nil
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
