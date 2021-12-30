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
	Hobbies    []string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Fajar",
		MiddleName: "Fsec",
		LastName:   "Sylvana",
		Age:        43,
		Married:    true,
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
