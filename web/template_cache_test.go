package go_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// NOTES: cara ini yg di rekomendasikan dlm menggunakan template.

//go:embed templates/*.gohtml
var (
	temps       embed.FS
	myTemplates = template.Must(template.ParseFS(temps, "templates/*.gohtml"))
)

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
