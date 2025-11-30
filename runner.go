package main

import (
	"golang.org/x/sys/unix"
	"log"
	"os"
	"os/exec"

	"github.com/MatthiasKunnen/xdg/desktop"
)

func runApp(e *desktop.Entry) {
	if e.Terminal {
		runTerm(e)
	} else {
		runGUI(e)
	}
}

func runTerm(e *desktop.Entry) {
	tokens := getFinalExec(e)
	execPath, err := getFullExecPath(tokens[0])
	if err != nil {
		log.Fatal(err)
	}

	if err := unix.Exec(execPath, tokens, os.Environ()); err != nil {
		log.Fatal(err)
	}
}

func runGUI(e *desktop.Entry) {
	tokens := getFinalExec(e)
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
