package author

import (
	"context"

	"github.com/graniticio/granitic/v2/ws"
)

type GetAuthorLogic struct{}

type Author struct {
	Name string `json:"name"`
}

func (gl *GetAuthorLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	a := make([]Author, 2)
	a[0] = Author{"Author 2"}
	a[1] = Author{"Author 3"}
	res.Body = a
}
