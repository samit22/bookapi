package book

import (
	"context"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/types"
	"github.com/graniticio/granitic/v2/ws"
)

type AddBookLogic struct {
	Log         logging.Logger
	FileManager FileAccessor
}

type FileAccessor interface {
	Write(name string, auth string) error
}

type AddBookRequest struct {
	Name   *types.NilableString `json:"name"`
	Author *types.NilableString `json:"author"`
}

func (al *AddBookLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, cb *AddBookRequest) {

	err := al.FileManager.Write(cb.Name.String(), cb.Author.String())
	if err != nil {
		al.Log.LogErrorf("Could not open file: %v", err)
		res.HTTPStatus = 400
		res.Body = map[string]string{"message": "Something went wrong"}
	}
	res.HTTPStatus = 201
	res.Body = map[string]string{"message": "Successfully added"}
}
