package main

// 旋转数组中的二分查找 (旋转数组: 按照升序排序的数组在预先未知的某个点上进行了旋转)
// 比如: 原数组 [1 2 3 4 5 6], 在5处发生了旋转，变为 旋转数组 [5 6 1 2 3 4]
// 注意: 以下代码要求旋转数组中没有重复元素
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] >= nums[l] {
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
	return -1
}

/*
	题目链接:
		https://leetcode-cn.com/problems/search-in-rotated-sorted-array/	搜索旋转排序数组
*/

/*
	总结
	1. 这题是二分查找的变形。
	2. 当我们得到中点mid时，由于我们此时不知道mid左右两边的递增情况，所以我们需要另外一个位置的数
       来确定左右两边的情况，上面的代码我采用的是和num[l]做比较。
       例如:  [5 6 1 2 3 4] l, r == 0, 5。
			那么此时mid == (l + r) / 2 == 2。由于nums[mid] < nums[l], 这说明nums[mid: r+1]部分是递增的，
			而nums[l: mid]这部分的情况我们还是不知道。	(由该旋转数组的性质可得出以上结论)
	3. 要注意，这个题目说了没有重复元素，如果有重复元素，那么就会出现 nums[mid] == nums[l] 的情况，此时我们并不能
	   得出左右两边的递增情况。
			比如 [3 -2 -1 1 2 3 3] [3 4 5 1 2 3 3]  对于mid == 2来说，它的左边可能存在2种情况[3 -2 -1]和[3 4 5]，可见递增情况是不确定的。
				 [3 3 3 -2 -1]	   [3 3 3 4 5]		对于mid == 2来说，它的右边也可能存在2种情况[3 -2 -1]和[3 4 5]，可见递增情况我们也不能确定。
	4. 上面写了很多的if判断可以使用位运算来简化，只是简化之后代码可读性就变差了。
*/
