package main

// 方法1: 2次循环
func isMonotonic(A []int) bool {
	flag1, flag2 := true, true
	// 判断数组是否单调递增
	for i := 1; i < len(A); i++ {
		if A[i-1] > A[i] {
			flag1 = false
		}
	}
	// 判断数组是否单调递减
	for i := 1; i < len(A); i++ {
		if A[i-1] < A[i] {
			flag2 = false
		}
	}
	return flag1 || flag2
}

// 方法2: 1次循环 (先判断单调性)
func isMonotonic(A []int) bool {
	if len(A) == 0 {
		return true
	}
	firstNum := A[0]
	lastNum := A[len(A)-1]
	// 相等单调时，表示数组所有元素全相等
	if firstNum == lastNum {
		for i := 0; i < len(A); i++ {
			if A[i-1] != A[i] {
				return false
			}
		}
	} else {
		// 表示非递减
		if firstNum < lastNum {
			for i := 1; i < len(A); i++ {
				if A[i-1] > A[i] {
					return false
				}
			}
		} else {
			// 表示非递增
			for i := 1; i < len(A); i++ {
				if A[i-1] < A[i] {
					return false
				}
			}
		}
	}

	return true
}

// 方法3: 1次循环 (先判断单调性) (方法2的优雅版本)
func isMonotonic(A []int) bool {
	if len(A) == 0 {
		return true
	}
	firstNum := A[0]
	lastNum := A[len(A)-1]
	for i := 1; i < len(A); i++ {
		if compare(A[i-1], A[i]) == 0 {
			continue
		}
		if compare(firstNum, lastNum) != compare(A[i-1], A[i]) {
			return false
		}
	}

	return true
}
// 比较大小
// a == b: return 0
// a >  b: return 1
// a <  b: return -1
func compare(a, b int) int {
	if a == b {
		return 0
	} else {
		if a > b {
			return 1
		} else {
			return -1
		}
	}
}

/*
	题目链接:
		https://leetcode-cn.com/problems/monotonic-array/submissions/		单调数列
*/
