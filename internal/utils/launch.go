package utils

import (
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

func LaunchGui(argv []string) error {
	args := argv[1:]
	path, err := FullPath(argv[0])
	if err != nil {
		return err
	}

	cmd := exec.Command(path, args...)
	cmd.Env = os.Environ()
	cmd.SysProcAttr = &unix.SysProcAttr{Setsid: true}

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Process.Release(); err != nil {
		return err
	}

	return nil
}

func LaunchTerm(argv []string) error {
	path, err := FullPath(argv[0])
	if err != nil {
		return err
	}

	if err := unix.Exec(path, argv, os.Environ()); err != nil {
		return err
	}

	return nil
}
