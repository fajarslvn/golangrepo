package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	/* Solution 1
	result := 0
	// for _, num := range numbers {
	// 	result += num
	// }
	return result */

	// Solution 2
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0] + Sum(numbers[1:])
	return result
}
