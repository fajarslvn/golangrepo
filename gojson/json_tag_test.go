package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Apple Mac Book Pro",
		ImageURL: "https://example.com/mbp.png",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","image_url":"https://example.com/mbp.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
