package main

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	indexMap[nums[0]] = 0

	for i := 1; i < len(nums); i++ {
		expected := target - nums[i]

		if expected == nums[0] {
			return []int{0, i}
		}

		if indexMap[expected] != 0 {
			return []int{indexMap[expected], i}
		} else {
			indexMap[nums[i]] = i
		}
	}
	return []int{}
}
