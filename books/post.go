package books

import (
	"context"
	"fmt"

	"github.com/graniticio/granitic/v2/ws"
)

type PostLogic struct {
}

func (gl *PostLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, cb *Book) {

	message := "Book added successfully!"

	if err := appendBooks(*cb); err != nil {
		message = fmt.Sprint(err)
	}

	res.Body = Response{Message: message}

}

type Response struct {
	Message string `json:"message"`
}
