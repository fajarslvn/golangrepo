package main

import (
	"embed"
	_ "embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}

	router.ServeFiles("/files/*filepath", http.FS(directory))

	req := httptest.NewRequest("GET", "http://localhost:3000/files/goodbye.txt", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	res := rec.Result()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "Goodbye Router", string(body))
}
