package 数组

// --------------------- 自己写的 ---------------------
func validMountainArray(A []int) bool {
	peak := 0
	for peak+1 < len(A) && A[peak] < A[peak+1] {
		peak++
	}
	if peak == 0 || peak == len(A)-1 {
		return false
	}
	for peak+1 < len(A) && A[peak] > A[peak+1] {
		peak++
	}
	return peak == len(A)-1
}

// --------------------- 参考评论写的 ---------------------
func validMountainArray(A []int) bool {
	left, right := 0, len(A)-1
	for left+1 < len(A) && A[left] < A[left+1] {
		left++
	}
	for right-1 >= 0 && A[right] < A[right-1] {
		right--
	}
	return left == right && left != len(A)-1 && right != 0
}

/*
	题目链接: https://leetcode-cn.com/problems/valid-mountain-array/submissions/
*/
