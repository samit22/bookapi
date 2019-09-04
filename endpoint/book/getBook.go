package book

import (
	"context"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type GetBookLogic struct {
	Log         logging.Logger
	FileManager FileReader
}

type FileReader interface {
	Read() ([]Book, error)
}

type Book struct {
	Name   string
	Author string
}

func (gl *GetBookLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {

	books, err := gl.FileManager.Read()
	if err != nil {
		gl.Log.LogErrorf("Could not read data file: %v", err)
		res.HTTPStatus = 400
	}
	res.Body = books
}
