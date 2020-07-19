package 数和问题

func twoSum(numbers []int, target int) []int {
	index1, index2 := getTwoIndexsWhichSumOfValueEqualRefInSortedArray(numbers, target)
	return []int{index1 + 1, index2 + 1}
}

func getTwoIndexsWhichSumOfValueEqualRefInSortedArray(array []int, ref int) (int, int) {
	left, right := 0, len(array)-1
	for left < right {
		curSum := array[left] + array[right]
		if curSum == ref {
			return left, right
		}
		if curSum > ref {
			right--
		} else {
			left++
		}
	}
	return -1, -1 // 不存在时才返回这个
}

/*
	题目链接: https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
	总结:
		1. 这题就是双指针~
*/
