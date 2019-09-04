package file

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
)

type ReaderWriter struct {
	FileName string
}

// func (f ReaderWriter) Write(...string) error {
// }

func (f ReaderWriter) Read() ([][]string, error) {
	output := make([][]string, 0)
	file, err := os.OpenFile(f.FileName, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return output, errors.New("error opening file")
	}
	defer file.Close()
	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return output, errors.New("error reading file")
	}
	return data, nil
}
