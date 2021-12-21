package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body":  "<p>Selamat belajar golang web<script>alert(9)</script></p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181", nil)
	rec := httptest.NewRecorder()

	TemplateAutoEscape(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8181",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body":  template.HTML("<p>Selamat belajar golang web</p><script>alert(9)</script>"),
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181", nil)
	rec := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8181",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181/?body=<p>Test</p>", nil)
	rec := httptest.NewRecorder()

	TemplateXSS(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func TestTemplateServerXSS(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8181",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
