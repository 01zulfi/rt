package commands

import (
	"github.com/01zulfi/rt/files"
	"github.com/01zulfi/rt/utils"
)

var AddAliases = [2]string{"a", "add"}

func Add(key string, value string) {
	path, err := files.Path()
	utils.Check(err)
	_, err = files.Exists(path + "/.rt.json")
	if err != nil {
		_, err := files.Create(path + "/.rt.json")
		utils.Check(err)
	}
	data, err := files.Read(path + "/.rt.json")
	utils.Check(err)
	if value == "" {
		data[key] = key
	} else {
		data[key] = value
	}
	err = files.Write(data, path+"/.rt.json")
	utils.Check(err)
}
