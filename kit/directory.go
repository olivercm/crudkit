package kit

import "os"

func MkdirAll(folderPath string) error {
	return os.MkdirAll(folderPath, os.ModePerm)
}
