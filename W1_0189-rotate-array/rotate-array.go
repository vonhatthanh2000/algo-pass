func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n
	mod := (n - k) % n
	rotateArr := make([]int, n)

	for i := range nums {
		rotateArr[i] = nums[(i+mod)%n]
	}

	for i, num := range rotateArr {
		nums[i] = num
	}
}

func rotateArray(nums []int, k int) {
	n := len(nums)
	k = k % n // Handle cases where k is greater than n
	if k == 0 {
		return // No rotation needed
	}

	// Reverse the entire array
	reverse(nums, 0, n-1)
	// Reverse the first k elements
	reverse(nums, 0, k-1)
	// Reverse the remaining n - k elements
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start] // Swap elements
		start++
		end--
	}
}
