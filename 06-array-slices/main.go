package main

import (
	"fmt"
	"sort"
)

func main() {
	var sweets = []string {
		"jelly",
		"dark chocolate",
		"candy",
		"milk chocolate",
	}

	fmt.Println("Sweets at index 0 is", sweets[0])

	sweets[1] = "Vanilla"
	fmt.Println("Sweets at index 1 is", sweets[1])

	fmt.Println("Length of sweets is", len(sweets))

	/*
	var scores []int
	scores = append(scores, 10)
	scores = append(scores, 15)
	scores = append(scores, 20)
	scores = append(scores, 25)
	scores = append(scores, 30)
	*/

	scores := []int {11, 12, 13, 14, 15, 0}
	fmt.Println("Scores slices:", scores)

	// sorting
	sort.Ints(scores)
	fmt.Println("Sorting scores:", scores)

	for i := range sweets {
		fmt.Println("Index:", i, "Sweets of:", sweets[i])
	}

	for index, score := range scores {
		fmt.Println("index:", index, "Value of:", score)
	}
  
	fmt.Println("The average score is", avgScores(scores))
}

func avgScores(scores []int) float64 {
	total := 0
	for _, score := range scores {
		total += score
	}
	return float64(total / len(scores))
}