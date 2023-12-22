package files

import "os"

func Path() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return dirname, nil
}
