package book

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/graniticio/granitic/v2/ws"
)

type GetBookLogic struct{}

type Book struct {
	Name   string
	Author string
}

func (gl *GetBookLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	books := make([]Book, 0)
	f, err := os.OpenFile("Book.csv", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf(err.Error())
		log.Fatal(err)
		return
	}
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Printf(err.Error())
		log.Fatal(err)
		return
	}
	books = make([]Book, len(data))
	for i, d := range data {
		books[i].Name = d[0]
		books[i].Author = d[1]
	}
	res.Body = books
	res.HTTPStatus = 200
}
