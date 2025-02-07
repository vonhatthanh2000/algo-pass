func twoSum(nums []int, target int) []int {
	arrMap := make(map[int]int)

	for i, num := range nums {
		expected := target - num

		if idx, exist := arrMap[expected]; exist {
			return []int{idx, i}
		}

		arrMap[num] = i
	}
	return []int{}
}
