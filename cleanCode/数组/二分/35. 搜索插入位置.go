package 二分

func searchInsert(nums []int, target int) int {
	return getIndexOfFirstGreaterOrEqual(nums, target)
}

func getIndexOfFirstGreaterOrEqual(nums []int, ref int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

/*
	题目链接: https://leetcode-cn.com/problems/search-insert-position/
*/
