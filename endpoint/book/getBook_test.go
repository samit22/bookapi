package book
import(
	"testing"
	"context"
	"errors"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
	"github.com/graniticio/granitic/v2/test"
)
type MockFileReaderResp struct{
	books [][]string
	err error
}
func(m MockFileReaderResp)Read() ([][]string, error){
	return m.books, m.err
}
func TestGetBookLogic_Process(t *testing.T){
	t.Log("When getting books succeeds")
	{
		t.Run("returs availiable books", func(t *testing.T) {
			book := make([][]string, 1)
			book[0] = []string{"same", "last"}
			mockR := MockFileReaderResp{ books: book, err: nil}
			log := logging.CreateAnonymousLogger("TestLoggger", logging.Fatal)
			gl := GetBookLogic{ Log: log, FileManager: mockR }
			req := &ws.Request{}
			res := &ws.Response{}
			gl.Process(context.TODO(), req, res)
			test.ExpectString(t, res.Body.([]Book)[0].Name, "same")
		})
	}
	t.Log("When getting books fails")
	{
		t.Run("returs error", func(t *testing.T) {
			book := make([][]string, 1)
			book[0] = []string{"same", "last"}
			mockR := MockFileReaderResp{ books: book, err: errors.New("something went wrong")}
			log := logging.CreateAnonymousLogger("TestLoggger", logging.Fatal)
			gl := GetBookLogic{ Log: log, FileManager: mockR }
			req := &ws.Request{}
			res := &ws.Response{}
			gl.Process(context.TODO(), req, res)
			test.ExpectInt(t, res.HTTPStatus, 400)
			test.ExpectNil(t, res.Body)
		})
	}
}