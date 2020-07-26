package Geek

const flagOfMissingNumber = -1

func missingNumber(nums []int) int {
	nums = append(nums, flagOfMissingNumber)
	for i := 0; i < len(nums); i++ {
		for !isRightPosition(nums, i) {
			makeNumToRightPosition(nums, i)
		}
	}
	return getMissNumber(nums)
}

func isRightPosition(nums []int, index int) bool {
	return nums[index] == flagOfMissingNumber || nums[index] == index
}

func makeNumToRightPosition(nums []int, indexOfNum int) {
	num := nums[indexOfNum]
	nums[num], nums[indexOfNum] = nums[indexOfNum], nums[num]
}

func getMissNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == flagOfMissingNumber {
			return i
		}
	}
	panic("题目出错")
}

/*
	题目链接: https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
	总结
		1. 这题和 _剑指 Offer 03. 数组中重复的数字_ 差不多，不过这题多了个前提 --- 数组是递增的。所以其实二分法也可以解这道题。

*/

