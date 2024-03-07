package main

import (
	"fmt"
	"os"
)

func main() {
	controller, err := InitializeEvent()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	controller.Start()
}
