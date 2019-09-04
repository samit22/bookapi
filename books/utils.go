package books

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func fetchBooks() ([]Book, error) {
	data, err := ioutil.ReadFile("assets/books.json")
	var books []Book
	if err != nil {
		fmt.Println("error:", err)
		return []Book{}, err
	}
	err = json.Unmarshal(data, &books)
	if err != nil {
		fmt.Println("error:", err)
		return []Book{}, err
	}
	return books, nil
}

func appendBooks(book Book) error {
	if book.Name == "" && book.Author == "" {
		return errors.New("Author and Name can't be empty")
	}
	books, err := fetchBooks()
	if err != nil {
		return err
	}
	books = append(books, book)
	file, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		return err
	}
	_ = ioutil.WriteFile("assets/books.json", file, 0644)
	return nil
}
