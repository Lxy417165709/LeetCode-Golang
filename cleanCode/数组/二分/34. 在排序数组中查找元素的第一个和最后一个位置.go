package 二分

func searchRange(nums []int, target int) []int {
	if isExist(nums, target) == false {
		return []int{-1, -1}
	}
	return []int{getIndexOfFirstGreaterOrEqual(nums, target), getIndexOfLastLessOrEqual(nums, target)}
}

func isExist(nums []int, ref int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == ref {
			return true
		}
		if nums[mid] > ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
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

func getIndexOfLastLessOrEqual(nums []int, ref int) int {
	return getIndexOfFirstGreater(nums, ref) - 1
}

func getIndexOfFirstGreater(nums []int, ref int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

/*
	题目链接: https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
*/
