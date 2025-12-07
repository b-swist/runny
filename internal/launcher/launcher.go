package launcher

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/b-swist/runny/internal/drun"
	"golang.org/x/sys/unix"
)

func Launch(e *drun.Entry) error {
	if e.Terminal {
		if err := runTerm(e); err != nil {
			return err
		}
	} else {
		if err := runGUI(e); err != nil {
			return err
		}
	}

	return nil
}

func runTerm(e *drun.Entry) error {
	tokens := drun.StripFieldCodes(e.Exec)
	execPath, err := getFullExecPath(tokens[0])
	if err != nil {
		return err
	}

	if err := unix.Exec(execPath, tokens, os.Environ()); err != nil {
		return err
	}

	return nil
}

func runGUI(e *drun.Entry) error {
	tokens := drun.StripFieldCodes(e.Exec)
	execPath, err := getFullExecPath(tokens[0])
	if err != nil {
		return err
	}

	cmd := exec.Command(execPath, tokens[1:]...)
	cmd.Env = os.Environ()
	cmd.SysProcAttr = &unix.SysProcAttr{Setsid: true}

	if err = cmd.Start(); err != nil {
		return err
	}

	if err = cmd.Process.Release(); err != nil {
		return err
	}

	return nil
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
