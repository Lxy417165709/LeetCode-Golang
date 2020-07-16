package 数组

func findLucky(arr []int) int {
	const maxBound, minBound = 500, 1
	countOfNum := make([]int, maxBound+1)
	for i := 0; i < len(arr); i++ {
		countOfNum[arr[i]]++
	}
	for i := maxBound; i >= minBound; i-- {
		if countOfNum[i] == i {
			return i
		}
	}
	return -1
}

/*
	题目链接: https://leetcode-cn.com/problems/find-lucky-integer-in-an-array/submissions/
	总结:
		1. 水题
*/
