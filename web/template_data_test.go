package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template data map",
		"Name":  "Fsec",
	})
}

func TestTemplateDataMap(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181", nil)
	rec := httptest.NewRecorder()

	TemplateDataMap(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template data struct",
		Name:  "Fsec",
		Address: Address{
			Street: "Wismas Pondok Cabe",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8181", nil)
	rec := httptest.NewRecorder()

	TemplateDataStruct(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
