package main

import "fmt"

func calculateNumber(numbers ...float64) float64 {
	fmt.Println(numbers)
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
	fmt.Println("The Average:", calculateNumber(2.3, 3.4, 4.5, 5.6))
}