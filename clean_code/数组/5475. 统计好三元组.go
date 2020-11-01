package 数组

// ------------------------ 纯暴力 ------------------------
func countGoodTriplets(arr []int, a int, b int, c int) int {
	countOfGoodTriplet := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if absDiff(arr[i], arr[j]) > a {
					continue
				}
				if absDiff(arr[j], arr[k]) > b {
					continue
				}
				if absDiff(arr[i], arr[k]) > c {
					continue
				}
				countOfGoodTriplet++
			}
		}
	}
	return countOfGoodTriplet
}
func absDiff(a, b int) int {
	return abs(a - b)
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// ------------------------ 暴力剪枝 ------------------------
func countGoodTriplets(arr []int, a int, b int, c int) int {
	countOfGoodTriplet := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if absDiff(arr[i], arr[j]) > a {
				continue
			}
			for k := j + 1; k < len(arr); k++ {
				if absDiff(arr[j], arr[k]) > b {
					continue
				}
				if absDiff(arr[i], arr[k]) > c {
					continue
				}
				countOfGoodTriplet++
			}
		}
	}
	return countOfGoodTriplet
}
func absDiff(a, b int) int {
	return abs(a - b)
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
