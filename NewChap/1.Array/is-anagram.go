package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sCount, tCount := make(map[rune]int), make(map[rune]int)

	for _, char := range s {
		sCount[char]++
	}
	for _, char := range t {
		tCount[char]++
	}

	for i, sC := range sCount {
		if sC != tCount[i] {
			return false
		}
	}

	return true
}
