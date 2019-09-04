package book

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/graniticio/granitic/v2/ws"
)

type AddBookLogic struct{}

type AddBookRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (gl *AddBookLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, ab *AddBookRequest) {
	a := new(Book)
	a.Name = ab.Name
	a.Author = ab.Author
	d := fmt.Sprintf("%s,%s\n", ab.Name, ab.Author)
	f, err := os.OpenFile("Book.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf(err.Error())
		log.Fatal(err)
		return
	}
	if _, err := f.Write([]byte(d)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
		return
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
		return
	}
	res.Body = a
	res.HTTPStatus = 201
}
