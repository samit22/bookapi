package book

import (
	"context"
	"strings"

	"github.com/graniticio/granitic/v2/ws"
)

type GetBookLogic struct {
	File FileReader
}

type FileReader interface {
	Read(path string) (string, error)
}

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (gl *GetBookLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	s, e := gl.File.Read("tmp/file.txt")
	if e != nil {
		res.Body = "Error reading file"
	} else {
		sp := strings.Split(s, ",")
		a := new(Book)
		a.Name = sp[0]
		a.Author = sp[1]

		res.Body = a
	}
}
