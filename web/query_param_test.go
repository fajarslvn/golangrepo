package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	first_name := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")
	fmt.Fprintf(w, "Hello %s %s", first_name, last_name)
}

func TestQueryParam(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8181/hello?first_name=john&last_name=doe", nil)
	rec := httptest.NewRecorder()

	SayHello(rec, req)
	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func MultipleParamsVal(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprintf(w, "Hello %s", strings.Join(names, " "))
}

func TestMultipleParamsVal(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8181/hello?name=Jack&name=The&name=Ripper", nil)
	rec := httptest.NewRecorder()

	MultipleParamsVal(rec, req)
	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
