package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Fajar",
		MiddleName: "Fsec",
		LastName:   "Sylvana",
		Age:        43,
		Married:    true,
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
