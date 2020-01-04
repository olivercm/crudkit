package utils

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}
	return string(fd)
}
