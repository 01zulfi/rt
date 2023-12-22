package main

import (
	"fmt"
	"os"

	"github.com/01zulfi/rt/commands"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	cmd := args[0]

	switch cmd {
	case "a", "add":
		commands.Add(args[1], args[2])
	case "v", "view":
		if len(args) > 1 {
			commands.View(args[1])
		} else {
			commands.View("")
		}
	default:
		fmt.Println("Invalid command")
	}
}
