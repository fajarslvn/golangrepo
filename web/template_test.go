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

func SimpleHtml(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}
	// Tanpa harus check error
	// t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "Hello from template!")
}

func TestSimpleHtml(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181", nil)
	rec := httptest.NewRecorder()

	SimpleHtml(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func SimpleHtmlFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "simple.gohtml", "Hello from File!")
}

func TestSimpleHtmlFile(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181", nil)
	rec := httptest.NewRecorder()

	SimpleHtmlFile(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "simple.gohtml", "Hello from directory!")
}

func TestTemplateDirectory(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181", nil)
	rec := httptest.NewRecorder()

	TemplateDirectory(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "simple.gohtml", "Hello from embed!")
}

func TestTemplateEmbed(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181", nil)
	rec := httptest.NewRecorder()

	TemplateEmbed(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
