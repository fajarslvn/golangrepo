package go_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// NOTES: cara ini yg di rekomendasikan dlm menggunakan template.

//go:embed templates/*.gohtml
var temps embed.FS

//go:embed templates/*.html
var tmps embed.FS

var myTemplates = template.Must(template.ParseFS(temps, "templates/*.gohtml"))
var myHtml = template.Must(template.ParseFS(tmps, "templates/*.html"))

func TemplateCache(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateCache(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181", nil)
	rec := httptest.NewRecorder()

	TemplateCache(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
