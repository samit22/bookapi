package file

import (
	"bookapi/endpoint/book"
	"encoding/csv"
	"errors"
	"os"

	"github.com/graniticio/granitic/v2/logging"
)

type Manager struct {
	FileName string
	Log      logging.Logger
}

func (f Manager) Write(name string, auth string) error {
	if !fileExists(f.FileName) {
		file, err := os.OpenFile(f.FileName, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			f.Log.LogErrorf("File doesn't exist, error creating one %v", err)
			return errors.New("couldn't create a new file")
		}
		err = writeToFile(file, []string{"Name", "Author"})
		if err != nil {
			f.Log.LogErrorf("Error writing header in new file %v", err)
			return errors.New("couldn't write into a new file")
		}
	}
	file, err := os.OpenFile(f.FileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		f.Log.LogErrorf("Could not open file: %v", err)
		return errors.New("couldn't open file")
	}
	err = writeToFile(file, []string{name, auth})
	if err != nil {
		f.Log.LogErrorf("Could not write into the file: %v", err)
		return errors.New("couldn't write data into file")
	}
	return nil
}

func (f Manager) Read() ([]book.Book, error) {
	books := make([]book.Book, 0)
	if !fileExists(f.FileName) {
		return books, errors.New("file doesn't exist")
	}
	file, err := os.OpenFile(f.FileName, os.O_RDONLY, 0644)
	if err != nil {
		return books, errors.New("error opening file")
	}
	defer file.Close()
	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return books, errors.New("error reading file")
	}
	books = make([]book.Book, len(data))
	for i, d := range data {
		books[i].Name = d[0]
		books[i].Author = d[1]
	}
	return books, nil

}

func writeToFile(file *os.File, data []string) error {
	w := csv.NewWriter(file)
	err := w.Write(data)
	if err != nil {
		return errors.New("error writing data to file")
	}
	w.Flush()
	return nil
}

func fileExists(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
