package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
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

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Fajar",
		Addresses: []Address{
			{
				Street:     "Wismamas",
				Country:    "Indonesia",
				PostalCode: "16516",
			},
			{
				Street:     "Elm Street",
				Country:    "Singapore",
				PostalCode: "17717",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{
		"FirstName":"Fajar", 
		"Addresses":[
				{"Street":"Wismamas","Country":"Indonesia","PostalCode":"16516"},
				{"Street":"Elm Street","Country":"Singapore","PostalCode":"17717"}
			]
		}`

	jsonBytes := []byte(jsonString)
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[
				{"Street":"Wismamas","Country":"Indonesia","PostalCode":"16516"},
				{"Street":"Elm Street","Country":"Singapore","PostalCode":"17717"}
			]`

	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Wismamas",
			Country:    "Indonesia",
			PostalCode: "16516",
		},
		{
			Street:     "Elm Street",
			Country:    "Singapore",
			PostalCode: "17717",
		},
	}

	bytes, err := json.Marshal(addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
