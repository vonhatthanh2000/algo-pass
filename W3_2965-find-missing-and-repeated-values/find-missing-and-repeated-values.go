func findMissingAndRepeatedValues(grid [][]int) []int {
	count := make(map[int]int)
	n := len(grid) * len(grid[0])

	for _, row := range grid {
		for _, num := range row {
			count[num]++
		}
	}

	var missing, repeated int

	for i := 1; i <= n; i++ {
		if count[i] == 2 {
			repeated = i
		} else if count[i] == 0 {
			missing = i
		}
	}

	return []int{repeated, missing}
}