package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		text := "Product " + id
		fmt.Fprint(w, text)
	})

	req := httptest.NewRequest("GET", "http://localhost:3000/products/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	res := rec.Result()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "Product 1", string(body))
}
