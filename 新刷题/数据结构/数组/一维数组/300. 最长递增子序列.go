package 一维数组

// lengthOfLIS 获取最长递增子序列的长度。
func lengthOfLIS(nums []int) int {
	// 1. 贪心遍历。遍历后，内部的最长递增子序列长度可能是错误的，但长度是正确的。
	fakeLisSeq := make([]int, 0)
	for _, num := range nums {
		index := getIndexOfFirstGreaterOrEqual(fakeLisSeq, num)
		if index == len(fakeLisSeq) {
			fakeLisSeq = append(fakeLisSeq, num)
		} else {
			fakeLisSeq[index] = num
		}
	}

	// 2. 返回。
	return len(fakeLisSeq)
}

// getIndexOfFirstGreaterOrEqual 获取第一个大于或等于目标值的索引。
func getIndexOfFirstGreaterOrEqual(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
