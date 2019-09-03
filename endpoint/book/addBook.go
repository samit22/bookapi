package book

import (
	"context"

	"github.com/graniticio/granitic/v2/types"
	"github.com/graniticio/granitic/v2/ws"
)

type AddBookLogic struct{}

type AddBookRequest struct {
	Name   *types.NilableString `json:"name"`
	Author *types.NilableString `json:"author"`
}

func (gl *AddBookLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, cb *AddBookRequest) {

	a := new(Book)
	a.Name = cb.Name.String()
	a.Author = cb.Author.String()

	res.Body = a
	res.HTTPStatus = 201
}
