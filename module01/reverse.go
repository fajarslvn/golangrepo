package module01

func Reverse(word string) string {
	// Solution 3 (pass the kanji test with rune):
	var result string
	for _, i := range word {
		result = string(i) + result
		// result = iterate: word + d(3), r(2), o(1), w(0) ->
	}
	return result
}

/*
		Solution 1 (Not pass the kanji test):
		var result string
		for i := len(word) - 1; i >= 0; i-- {
			result = result + string(word[i])
		}
			return result

	Solution 2 (Not pass the kanji test):
	var result string
	for i := 0; i < len(word); i++ {
		result = string(word[i]) + result
	}
	return result
*/
