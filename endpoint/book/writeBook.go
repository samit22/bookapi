package book

import (
	"context"

	"github.com/graniticio/granitic/v2/ws"
)

type WriteBookLogic struct {
	File FileWriter
}
type FileWriter interface {
	Write(path string, val string) error
}

type WriteBookRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (gl *WriteBookLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, cb *WriteBookRequest) {
	s := cb.Name + "," + cb.Author
	err := gl.File.Write("tmp/file.txt", s)
	if err != nil {
		res.Body = "Error writing in file, ran out of ink"
		res.HTTPStatus = 404
	}
	res.Body = "Successfully written"
	res.HTTPStatus = 201
}
