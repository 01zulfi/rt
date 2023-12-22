package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func prettyPrint(data map[string]string) {
	for key, value := range data {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func main() {
	var data map[string]string
	content, err := os.ReadFile("data.json")
	check(err)
	err = json.Unmarshal(content, &data)
	check(err)

	args := os.Args[1:]
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
		if value == "" {
			data[key] = key
			bytes, err := json.Marshal(data)
			check(err)
			os.WriteFile("data.json", bytes, 0644)
		} else {
			data[key] = value
			bytes, err := json.Marshal(data)
			check(err)
			os.WriteFile("data.json", bytes, 0644)
		}
	case "v", "view":
		var key string
		if len(args) > 1 {
			key = args[1]
		}
		if key == "" {
			prettyPrint(data)
		} else {
			value := data[key]
			if value != "" {
				fmt.Printf("%s: %s\n", key, value)
			} else {
				fmt.Println("Key not found")
				fmt.Println("")
				prettyPrint(data)
			}
		}
	default:
		prettyPrint(data)
	}
}
