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
		var key string
		var value string
		if len(args) > 1 {
			key = args[1]
		}
		if len(args) > 2 {
			value = args[2]
		}

		commands.Add(key, value)
	case "v", "view":
		if len(args) > 1 {
			commands.View(args[1])
		} else {
			commands.View("")
		}
	case "r", "remove":
		if len(args) > 1 {
			commands.Remove(args[1])
		} else {
			fmt.Println("remove: need argument")
		}
	case "x", "clear-all":
		commands.ClearAll()
	default:
		fmt.Println("Invalid command")
	}
}
