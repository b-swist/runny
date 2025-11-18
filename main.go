package main

import (
	"fmt"
	"log"
)

func main() {
	entries, err := getAppEntries()
	if err != nil {
		log.Fatal(err)
	}

	for i, entry := range entries {
		fmt.Printf("#%v %v: %v\n", i, getDefaultName(entry), getFinalExec(entry))
	}
	var index int
	fmt.Print("Index: ")
	fmt.Scan(&index)
	runApp(entries[index])
}
