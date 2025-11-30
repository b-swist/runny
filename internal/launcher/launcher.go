package launcher

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/b-swist/runny/internal/desktop"
	"golang.org/x/sys/unix"
)

func Launch(e *desktop.Entry) {
	if e.Terminal {
		runTerm(e)
	} else {
		runGUI(e)
	}
}

func runTerm(e *desktop.Entry) {
	tokens := desktop.GetFinalExec(e)
	execPath, err := getFullExecPath(tokens[0])
	if err != nil {
		log.Fatal(err)
	}

	if err := unix.Exec(execPath, tokens, os.Environ()); err != nil {
		log.Fatal(err)
	}
}

func runGUI(e *desktop.Entry) {
	tokens := desktop.GetFinalExec(e)
	execPath, err := getFullExecPath(tokens[0])
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(execPath, tokens[1:]...)
	cmd.Env = os.Environ()
	cmd.SysProcAttr = &unix.SysProcAttr{Setsid: true}

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err = cmd.Process.Release(); err != nil {
		log.Fatal(err)
	}
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
