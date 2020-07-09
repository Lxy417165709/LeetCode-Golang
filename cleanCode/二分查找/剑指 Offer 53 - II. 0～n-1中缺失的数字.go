package Geek

func missingNumber(nums []int) int {
	return getInsertIndexOfMissingNumber(nums)
}

func getInsertIndexOfMissingNumber(nums []int) int{
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			if nums[mid] > mid {
				right = mid - 1
			} else {
				panic("题目出错")
			}
		}
	}
	return right + 1
}


/*
	题目链接: https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
	总结
		1. 这题和 _剑指 Offer 03. 数组中重复的数字_ 差不多，不过这题多了个前提 --- 数组是递增的。所以其实二分法也可以解这道题。
		2. 这里采用二分解决
*/
