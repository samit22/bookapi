package books

import (
	"context"

	"github.com/graniticio/granitic/v2/ws"
)

type GetLogic struct{}

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (gl *GetLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	book, _ := fetchBooks()
	res.Body = book
}
