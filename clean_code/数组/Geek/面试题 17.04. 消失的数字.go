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
	题目链接: https://leetcode-cn.com/problems/missing-number-lcci/
	总结
		1. 这题的解题思路和 _剑指 Offer 03. 数组中重复的数字_ 差不多，都是通过把数字存储到"正确"的位置上，将哈希信息存储到原数组中。
		2. 官方还有利用数学求和公式AC、以及用异或AC的...
		3. 这种方法有人叫做原址哈希。
*/
