package book

import (
	"context"

	"github.com/graniticio/granitic/v2/ws"
)

type GetBookLogic struct{}

type Book struct {
	Name   string
	Author string
}

func (gl *GetBookLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {

	a := new(Book)
	a.Name = "My Book"
	a.Author = "Ram"

	res.Body = a
}
