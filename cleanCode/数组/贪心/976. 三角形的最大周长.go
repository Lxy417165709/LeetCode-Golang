package 贪心

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1; i >= 2; i-- {
		lengthOfLargestLine := A[i]
		lengthOfLine1, lengthOfLine2 := A[i-1], A[i-2]
		if lengthOfLine1+lengthOfLine2 > lengthOfLargestLine {
			return lengthOfLine1 + lengthOfLine2 + lengthOfLargestLine
		}
	}
	return 0
}

/*
	题目链接: https://leetcode-cn.com/problems/largest-perimeter-triangle/
*/
