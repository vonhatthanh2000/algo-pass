func letterCombinations(digits string) []string {
	// Mapping of digits to letters
	obj := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	if len(digits) == 0 {
		return []string{}
	}

	// Initialize result array with the letters corresponding to the first digit
	resArr := obj[rune(digits[0])]

	// If only one digit, return the corresponding letters
	if len(digits) == 1 {
		return resArr
	}

	// Combine letters for the remaining digits
	for i := 1; i < len(digits); i++ {
		char := rune(digits[i])
		resArr = combine2arr(resArr, obj[char])
	}

	return resArr
}

// combine2arr combines two string slices into one.
func combine2arr(arr1 []string, arr2 []string) []string {
	resArr := []string{}
	for _, fstLetter := range arr1 {
		for _, sndLetter := range arr2 {
			resArr = append(resArr, fstLetter+sndLetter)
		}
	}
	return resArr
}