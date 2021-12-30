package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Ali",
		MiddleName: "Ballan",
		LastName:   "Sylvana",
		Hobbies:    []string{"coding", "reading", "basketball"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{
		"FirstName":"Ali",
		"MiddleName":"Ballan",
		"LastName":"Sylvana",
		"Hobbies":
			[
				"coding",
				"reading",
				"basketball"
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
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)

}
