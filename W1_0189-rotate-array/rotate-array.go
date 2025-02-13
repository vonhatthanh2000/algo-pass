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