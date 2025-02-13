func moveZeroes(nums []int) {
	lastNonZeroIdx := 0
	for _, num := range nums {
		if num != 0 {
			nums[lastNonZeroIdx] = num
			lastNonZeroIdx++
		}
	}
	for i := lastNonZeroIdx; i < len(nums); i++ {
		nums[i] = 0
	}
}