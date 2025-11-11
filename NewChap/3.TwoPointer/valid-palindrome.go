package main

import (
	"fmt"
	"unicode"
)

func main() {

	a := "cat a tac"
	fmt.Println(isPalindrome(a))
}

func isPalindrome(s string) bool {
	runeString := []rune(s)
	l, r := 0, len(s)-1

	for l < r {
		for (l < r) && !isValid(runeString[l]) {
			l++
			fmt.Println(l)
		}
		for (r > l) && !isValid(runeString[r]) {
			fmt.Println(r)

			r--
		}

		if unicode.ToLower(runeString[l]) != unicode.ToLower(runeString[r]) {
			return false
		}
		l++
		r--
	}

	return true

}

func isValid(s rune) bool {
	if unicode.IsLetter(s) || unicode.IsDigit(s) {
		return true
	}
	return false
}
