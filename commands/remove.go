package commands

import (
	"fmt"

	"github.com/01zulfi/rt/files"
	"github.com/01zulfi/rt/utils"
)

func Remove(key string) {
	path, err := files.Path()
	utils.Check(err)
	_, err = files.Exists(path + "/.rt.json")
	if err != nil {
		_, err := files.Create(path + "/.rt.json")
		utils.Check(err)
	}
	data, err := files.Read(path + "/.rt.json")
	utils.Check(err)
	_, ok := data[key]
	if !ok {
		fmt.Printf("remove: can't find %s", key)
		return
	}
	delete(data, key)
	files.Write(data, path+"/.rt.json")
}
