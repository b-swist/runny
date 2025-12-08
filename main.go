package main

import (
	"fmt"
	"os"

	"github.com/b-swist/runny/cmd"
)

func main() {
	if err := cmd.Main(); err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	os.Exit(0)
}
