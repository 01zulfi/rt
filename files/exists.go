package files

import "os"

func Exists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil {
		return true, nil
	}
	return false, err
}
