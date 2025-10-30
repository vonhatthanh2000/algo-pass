package main

func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int)
	arrCount := make([][]int, len(nums)+1)

	for _, num := range nums {
		frequencyMap[num]++
	}

	for num, count := range frequencyMap {
		arrCount[count] = append(arrCount[count], num)
	}

	res := []int{}

	for i := len(arrCount) - 1; i > 0; i-- {
		for _, num := range arrCount[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}

	}

	return res
}
