package utils

import (
	"os"
	"path/filepath"
)

func XDGCurrentDesktop() []string {
	env, ok := os.LookupEnv("XDG_CURRENT_DESKTOP")
	if !ok {
		return nil
	}
	return filepath.SplitList(env)
}

func XDGStateHome() (string, error) {
	dir := os.Getenv("XDG_STATE_HOME")

	if dir == "" || !filepath.IsAbs(dir) {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		dir = filepath.Join(home, ".share", "local")
	}

	return dir, nil
}
