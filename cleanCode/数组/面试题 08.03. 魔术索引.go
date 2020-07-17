package 数组

// ----------------------- 方法1: 暴力 -----------------------
func findMagicIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i {
			return i
		}
	}
	return -1
}

// ----------------------- 方法2: 分治 -----------------------
const INF = 10000000000

func findMagicIndex(nums []int) int {
	magicIndex := getMagicIndex(nums, 0, len(nums)-1)
	if magicIndex == INF {
		return -1
	}
	return magicIndex
}

func getMagicIndex(nums []int, left, right int) int {
	if left > right {
		return INF
	}
	mid := left + (right-left)/2
	if nums[mid] == mid {
		return min(mid, getMagicIndex(nums, left, mid-1))
	}
	magicIndex := getMagicIndex(nums, left, mid-1)
	if magicIndex!=INF{
		return magicIndex
	}
	return getMagicIndex(nums, mid+1, right)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// ----------------------- 方法3: 二分 -----------------------
// 能二分吗？


/*
	题目链接: https://leetcode-cn.com/problems/magic-index-lcci/comments/
	总结:
		1. 这题好像有问题，数组是有序的，但是好像不能二分。
*/
