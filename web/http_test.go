package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8181/hello", nil)
	record := httptest.NewRecorder()

	HelloHandler(record, req)
	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(body)
	fmt.Println(bodyString)
}
