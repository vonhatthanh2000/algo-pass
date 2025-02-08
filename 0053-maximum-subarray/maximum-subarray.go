func maxSubArray(nums []int) int {
	maxSub := nums[0]
	subSum := nums[0]
	for _, num := range nums[1:] {
		if subSum < 0 {
			subSum = num
		} else {
			subSum += num
		}
		maxSub = max(maxSub, subSum)
	}
	return maxSub
}
