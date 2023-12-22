package files

import (
	"encoding/json"
	"os"
)

func Read(filepath string) (map[string]string, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var data map[string]string
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
