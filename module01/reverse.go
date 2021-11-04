package module01

func Reverse(word string) string {
	/*
		Solution 1:
		var result string
		for i := len(word) - 1; i >= 0; i-- {
			result = result + string(word[i])
		}
			return result
	*/

	/* Solution 2:
	var result string
	for i := 0; i < len(word); i++ {
		result = string(word[i]) + result
	}
	return result
	*/

	// Solution 3 (passa the test on kanji character):
	var result string
	for _, r := range word {
		result = string(r) + result
	}
	return result
}
