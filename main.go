package main

import (
	"fmt"
	"log"
)

func main() {
	entries, err := getEntries()
	if err != nil {
		log.Fatal(err)
	}

	for i, entry := range entries {
		fmt.Printf("%v: %v\n", i, (*entry).Name.Default)
	}
}
