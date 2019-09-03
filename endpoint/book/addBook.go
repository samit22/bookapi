package book

import (
	"context"

	"github.com/graniticio/granitic/v2/ws"
)

type AddBookLogic struct{}

type AddBookRequest struct {
	Name   string `json:"name"`
	Author string `json:author`
}

func (gl *AddBookLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, cb *AddBookRequest) {

	a := new(Book)
	a.Name = cb.Name
	a.Author = cb.Author

	res.Body = a
	res.HTTPStatus = 201
}
