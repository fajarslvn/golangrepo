package main

import "testing"

func TestAdd(t *testing.T) {
	tables := []struct {
		name string
		numberOne float64
		numberTwo float64
		expected float64
	}{
		 {"valid-addition-1", 4, 6, 10},
		 {"valid-addition-2", 5, 7, 12},
		 {"valid-addition-3", 10, 5, 15},
	}
	for _, tt := range tables {
		result, _ := Add(tt.numberOne, tt.numberTwo)

		if result != tt.expected {
			t.Error("Incorrect calculation, Expected:", tt.expected)
		}
	}
}

func TestDivision(t *testing.T){
	tables := []struct {
		name string
		dividen float64
		divisor float64
		expected float64
		isError bool 
	}{
		{"valid-division-1", 8, 4, 2, false},
		{"valid-division-2", 10, 5, 2, false},
		{"division-by-0", 334, 0, 0, true},
	}
	for _, tt := range tables {
		result, err  := Divide(tt.dividen, tt.divisor)

		if tt.isError {
			if err == nil {
				t.Error("expected an error, but didn't get ones")
			 }
			} else {
				if tt.expected != result {
					t.Error("expected", tt.expected, "but got one:", result)
				} 
		}
	}
}

func TestIncorrectDivision(t *testing.T) {
	_, err := Divide(10, 0)

	if err == nil {
		t.Error("expected an error, but didn't get one")
	} 
}

/*func TestAdd(t *testing.T) {
	 result, _ := Add(5, 6)

	 if result != 11 {
		 t.Error("Incorrect calculation, Expected:", 11)
	 }
}
func TestDivision(t *testing.T) {
	result, err := Divide(10, 5)

	if err != nil {
		t.Error("did'nt expect an error, but got one:", err.Error())
	} else {
		if result != 2.0 {
			t.Error("Incorect division calculation")
		}
	}
}*/ 