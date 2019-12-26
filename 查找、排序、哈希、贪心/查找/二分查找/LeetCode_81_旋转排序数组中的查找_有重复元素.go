package main

// 旋转数组中的二分查找 (旋转数组: 按照升序排序的数组在预先未知的某个点上进行了旋转)
// 比如: 原数组 [1 2 3 4 5 6], 在5处发生了旋转，变为 旋转数组 [5 6 1 2 3 4]
// 注意: 此时旋转数组中可能有重复元素
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target == nums[mid] {
			return true
		}

		// 与没有重复元素的代码相比，就多了这一句
		if nums[mid] == nums[l] {
			l++
			continue
		}

		if nums[mid] > nums[l] {
			if target < nums[mid] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}

/*
	题目链接:
		https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/		搜索旋转排序数组2
*/

/*
	总结
	1. 这题和之前的差不多，只是此时有了重复元素，那么就会导致nums[l] == num[mid]使我们无法判断左右任意一边的情况，
       于是让左边界前移(因为此时我们可以确定左边界指向的值不是目标值)，缩小查找范围。
	2. 这题的最坏时间复杂度是O(n) (数组元素都一样且不存在目标值时)，平均是O(logn)
*/
