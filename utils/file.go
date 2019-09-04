package utils

import (
	"io/ioutil"
	"os"
)

type File struct{}

func (f *File) Read(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	} else {
		return string(dat), err
	}
}

func (f *File) Write(path string, val string) error {
	fp, err := os.Create(path)
	if err == nil {
		fp.WriteString(val)
	}
	return err
}
