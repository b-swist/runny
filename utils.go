package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"slices"
)

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

func getFullExecPath(cmd string) (string, error) {
	if filepath.IsAbs(cmd) {
		return cmd, nil
	}

	path, err := exec.LookPath(cmd)
	if err != nil {
		return "", err
	}

	return path, nil
}
