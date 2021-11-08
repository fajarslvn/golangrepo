package module01

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	result := 0
	multiplier := 1

	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for j, char := range charset {
			if char == rune(value[i]) {
				index = j
				break
			}
		}
		if index < 0 {
			panic("this shouldn't ever happen!")
		}
		result += index * multiplier
		multiplier *= base
	}
	return result
}

/*
Alternative solution:
{
	result := 0
	multiplier := 1
	for i := len(value) - 1; i >= 0; i-- {
		var val int
		fmt.Sscanf(string(value[i]), "%X", &val)
		result += multiplier * val
		multiplier *= base
	}
	return result
}
*/
