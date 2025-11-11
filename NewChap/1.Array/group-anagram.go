package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}

	indexScan := make(map[int]bool)
	for i := 0; i < len(strs); i++ {
		if indexScan[i] == true {
			continue
		}
		subRes := []string{strs[i]}
		indexScan[i] = true
		for j := i + 1; j < len(strs); j++ {
			if indexScan[j] == true {
				continue
			}

			if isAnagram(strs[i], strs[j]) == true {

				subRes = append(subRes, strs[j])
				indexScan[j] = true
			}
		}
		res = append(res, subRes)
	}
	return res
}

func groupAnagrams2(strs []string) [][]string {
	sortString := func(str string) string {
		strRunes := []rune(str)
		sort.Slice(strRunes, func(i, j int) bool {
			return strRunes[i] < strRunes[j]
		})
		return string(strRunes)
	}

	strMap := make(map[string][]string)

	for _, str := range strs {
		keyStr := sortString(str)
		strMap[keyStr] = append(strMap[keyStr], str)
	}

	var res [][]string

	for _, group := range strMap {
		res = append(res, group)
	}

	return res
}
func main() {
	input := []string{"act", "pots", "tops", "cat", "stop", "hat"}

	res := groupAnagrams(input)

	fmt.Println("Res", res)
}
