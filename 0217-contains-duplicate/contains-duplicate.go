// 1st try
func containsDuplicate(nums []int) bool {
	used := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := used[num]; exists {
			return true
		}
		used[num] = struct{}{}
	}
	return false
}

// better
func containsDuplicate(nums []int) bool {
	used := make(map[int]bool, len(nums))

	for _, num := range nums {
		if _, exists := used[num]; exists {
			return true
		}
		used[num] = true
	}
	return false
}