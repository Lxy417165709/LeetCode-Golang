package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		// 上面的条件判断可以写为 l < r，之后在函数体外返回 nums[l]就可以了，但是我习惯了l<=r，所以就在
		// 函数体里面判断了。
		if l == r {
			return nums[l]
		}
		mid := (l + r) / 2
		// 由于数组没有重复元素，num[r] == nums[mid] 时表示扫描区域只有一个元素了()
		if nums[r] == nums[mid] {
			// 这句不会被执行到，只是为了符合结构，我就这样写了。
			return nums[r]
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
		https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/		寻找旋转排序数组中的最小值
*/
/*
	总结
	1. 这题其它做法:
			(1) 排序后获取第一个元素: 时间复杂度O(nlogn)
			(2) 线性扫描一次数组，获取最小值: 时间复杂度O(n)
			(3) 线性扫描数组，遇到第一个变小的点就进行输出: 时间复杂度O(n)
	2. 这类题的共性就是需要找数组中的一些位置来确定数组局部的性质，再根据性质进行操作。

*/
