package Geek

import "sort"

func repeatedNTimes(A []int) int {
	sort.Ints(A)
	leftMid := len(A)/2 - 1
	if A[leftMid] == A[leftMid-1] {
		return A[leftMid]
	} else {
		return A[leftMid+1]
	}
}

func repeatedNTimes(A []int) int {
	for dis := 1; dis <= 3; dis++ {
		for i := dis; i < len(A); i++ {
			if A[i] == A[i-dis] {
				return A[i]
			}
		}
	}
	panic("题目出错")
}

/*
	题目链接: https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array/
	总结:
		1. 方法2利用了该重复元素的距离小于等于3
*/
