package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Header", "Header by Hax0r")
	fmt.Fprint(w, "\nOK")
}

func TestHeader(t *testing.T) {
	req := httptest.NewRequest("POST", "localhost:8181", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	RequestHeader(rec, req)
	ResponseHeader(rec, req)

	custHeader := rec.Header().Get("X-Header")

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("X-Header:", custHeader)
	fmt.Println("Content-Type:", string(body))
}
