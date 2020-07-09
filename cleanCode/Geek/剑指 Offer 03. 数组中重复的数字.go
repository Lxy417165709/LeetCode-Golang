package Geek

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for !isRightPosition(nums, i) {
			if isNumInRightPosition(nums, i) {
				return nums[i]
			}
			makeNumToRightPosition(nums, i)
		}
	}
	panic("题目出错")
}

func makeNumToRightPosition(nums []int, indexOfNum int) {
	num := nums[indexOfNum]
	nums[indexOfNum], nums[num] = nums[num], nums[indexOfNum]
}

func isNumInRightPosition(nums []int, indexOfNum int) bool {
	num := nums[indexOfNum]
	return nums[indexOfNum] == nums[num]
}

func isRightPosition(nums []int, index int) bool {
	return nums[index] == index
}

/*
	题目链接: https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
	总结
		1. 重构之后，代码更易懂，而且自己对这题的理解更深入了。
*/
