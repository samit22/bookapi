package book

import (
	"context"
	"errors"
	"testing"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/test"
	"github.com/graniticio/granitic/v2/types"
	"github.com/graniticio/granitic/v2/ws"
)

type MockFileWriter struct {
	err error
}

func (m MockFileWriter) Write(string, string) error {
	return m.err
}
func TestAddBookLogic_ProcessPayload(t *testing.T) {
	t.Log("When adding books succeeds")
	{
		t.Run("returs 201", func(t *testing.T) {
			mockW := MockFileWriter{}
			log := logging.CreateAnonymousLogger("TestLoggger", logging.Fatal)
			abl := AddBookLogic{Log: log, FileManager: mockW}
			ab := AddBookRequest{Name: types.NewNilableString("My Book"), Author: types.NewNilableString("Ram")}
			req := &ws.Request{}
			res := &ws.Response{}
			abl.ProcessPayload(context.TODO(), req, res, &ab)
			test.ExpectString(t, res.Body.(map[string]string)["message"], "Successfully added")
			test.ExpectInt(t, res.HTTPStatus, 201)
		})
	}
	t.Log("When adding books fails")
	{
		t.Run("returs error", func(t *testing.T) {
			mockW := MockFileWriter{err: errors.New("couldn't create a new file")}
			log := logging.CreateAnonymousLogger("TestLoggger", logging.Fatal)
			al := AddBookLogic{Log: log, FileManager: mockW}
			ar := AddBookRequest{Name: types.NewNilableString("My Book"), Author: types.NewNilableString("Ram")}
			req := &ws.Request{}
			res := &ws.Response{}
			al.ProcessPayload(context.TODO(), req, res, &ar)
			test.ExpectString(t, res.Body.(map[string]string)["message"], "Something went wrong")
			test.ExpectInt(t, res.HTTPStatus, 400)
		})
	}
}
