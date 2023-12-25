package commands

import (
	"os"

	"github.com/01zulfi/rt/files"
	"github.com/01zulfi/rt/utils"
)

func ClearAll() {
	path, err := files.Path()
	utils.Check(err)
	err = os.WriteFile(path+"/.rt.json", []byte("{}"), 0644)
	utils.Check(err)
}
