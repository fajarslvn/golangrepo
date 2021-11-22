package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is empty")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func TestResponseCodeValid(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181/?name=fsec", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}
