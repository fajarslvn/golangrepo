package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Ali",
		MiddleName: "Ballan",
		LastName:   "Sylvana",
	}

	_ = encoder.Encode(customer)
	fmt.Println(customer)
}
