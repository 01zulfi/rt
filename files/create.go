package files

import "os"

func Create(filepath string) (*os.File, error) {
	f, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}
	_, err = f.Write([]byte("{}"))
	return f, err
}
