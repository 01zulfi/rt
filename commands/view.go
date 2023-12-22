package commands

import (
	"fmt"

	"github.com/01zulfi/rt/files"
	"github.com/01zulfi/rt/utils"
)

func prettyPrintEntry(key string, value string) {
	fmt.Printf("* %s: %s\n", key, value)
}

func prettyPrint(data map[string]string) {
	for key, value := range data {
		prettyPrintEntry(key, value)
	}
}

var ViewAliases = [2]string{"v", "view"}

func View(key string) {
	path, err := files.Path()
	utils.Check(err)
	_, err = files.Exists(path + "/.rt.json")
	if err != nil {
		_, err := files.Create(path + "/.rt.json")
		utils.Check(err)
	}
	data, err := files.Read(path + "/.rt.json")
	utils.Check(err)
	if key == "" {
		prettyPrint(data)
	} else {
		value := data[key]
		if value == "" {
			fmt.Printf("No entry found for %s\n", key)
			fmt.Println("")
			prettyPrint(data)
		} else {
			fmt.Println(value)
		}
	}
}
