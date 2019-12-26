package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		// l == r 表示范围内只有一个元素了，那么这个元素就是最小值。
		if l == r {
			return nums[l]
		}
		mid := (l + r) / 2
		// nums[r] == nums[mid] 时，我们无法判断最小值位于左边还是右边，但能确定，最小值一定不在边界
		if nums[r] == nums[mid] {
			r--
		} else {
			if nums[r] > nums[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接:
		https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/		寻找旋转排序数组中的最小值 II
*/
/*
	总结
	1. 这题是寻找旋转排序数组中的最小值的升级版，此时数组中有了重复元素。
	2. 这题和原始的题解法差不多，只是修改了一个地方。

*/
