package main

import "errors"


func Divide(dividen, divisor float64) (float64, error)  {
	// can't  divide by 0
	var result float64

	if divisor == 0 {
		return result, errors.New("can't divide by zero")
	}
	result = dividen / divisor
	return result, nil
}

func Add(num1, num2 float64) (float64, error) {
	result := num1 + num2
	return result, nil
}

func main() {}