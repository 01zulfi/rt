package files

import (
	"encoding/json"
	"os"
)

func Write(data map[string]string, filepath string) error {
	contents, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, contents, 0644)
	if err != nil {
		return err
	}
	return nil
}
