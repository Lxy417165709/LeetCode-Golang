package Geek

// ----------------------- 方法1: 时空复杂度O(n),O(1) -----------------------
// 总体思路: 将数组元素存放到"正确"的位置上。
// 		所谓 "正确"位置 就是:
// 			举个例子，[1, 2, 4, 3], 只有元素 1, 2 位于"正确"的位置上, 而 4, 3 没有
//			而 [1, 2, 3, 4], 所有元素都在正确的位置上。
const disappearFlag = -1

func findDisappearedNumbers(nums []int) []int {
	makeNumsToRightPosition(nums)
	return getDisappearedNumbers(nums)
}

func makeNumsToRightPosition(nums []int) {
	for i := 0; i < len(nums); i++ {
		for !isRightPosition(nums, i) {
			if isNumInRightPosition(nums, i) {
				nums[i] = disappearFlag
			} else {
				makeNumToRightPosition(nums, i)
			}
		}
	}
}

func getDisappearedNumbers(nums []int) []int {
	disappearedNumbers := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == disappearFlag {
			disappearedNumbers = append(disappearedNumbers, i+1)
		}
	}
	return disappearedNumbers
}

func isRightPosition(nums []int, index int) bool {
	return nums[index] == disappearFlag || nums[index] == index+1
}

func isNumInRightPosition(nums []int, indexOfNum int) bool {
	rightPosition := nums[indexOfNum] - 1
	return nums[rightPosition] == nums[indexOfNum]
}

func makeNumToRightPosition(nums []int, indexOfNum int) {
	rightPosition := nums[indexOfNum] - 1
	nums[indexOfNum], nums[rightPosition] = nums[rightPosition], nums[indexOfNum]
}

/*
	题目链接: https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
	类似题型:
		1. [剑指 Offer 03. 数组中重复的数字](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/)
		2. [剑指 Offer 53 - II. 0～n-1中缺失的数字](https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/)
*/
