package book

import (
	"context"

	"github.com/graniticio/granitic/v2/ws"
)

type GetAuthorsLogic struct{}

type Author struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func (gl *GetAuthorsLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	s := make([]Author, 2)

	b := Author{Name: "blaaaaa", Country: "fasfsa"}

	s[0] = Author{Name: "bla", Country: "fasfsa"}
	s[1] = b

	res.Body = s
}
