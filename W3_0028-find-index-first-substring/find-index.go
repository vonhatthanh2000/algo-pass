func strStr(haystack string, needle string) int {
	subLength := len(needle)

	if subLength == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-subLength; i++ {
		if haystack[i] != needle[0] {
			continue
		}
		subStr := strings.TrimSpace(haystack[i : i+subLength])

		if subStr == needle {
			return i
		}
	}
	return -1
}