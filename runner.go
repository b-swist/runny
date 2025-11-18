package main

import (
	"github.com/MatthiasKunnen/xdg/desktop"
)

func runApp(e *desktop.Entry) {
	if e.Terminal {
		runTerm(e)
	} else {
		runGUI(e)
	}
}

func runTerm(e *desktop.Entry) {}

func runGUI(e *desktop.Entry) {}
