package 读写指针

/*
	给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。
	你可以返回满足此条件的任何数组作为答案。
*/

// 会改变数组元素的相对位置		(解法1)
func sortArrayByParity(A []int) []int {
	reader, writer := 0, 0
	for reader < len(A) {
		if A[reader]&1 == 0 {
			A[writer], A[reader] = A[reader], A[writer]
			writer++
		}
		reader++
	}
	return A
}

// 不会改变数组元素的相对位置的解法 (但是时间效率会低一些)		(解法2)
func sortArrayByParity(A []int) []int {
	reader, writer := 0, 0
	for reader < len(A) {
		if A[reader]&1 == 0 {
			// 这可以保证相对位置不变
			for i := reader; i >= writer+1; i-- {
				A[i], A[i-1] = A[i-1], A[i]
			}
			writer++
		}
		reader++
	}
	return A
}

// 另外一种指针定义: 左右指针法		(解法3)
func sortArrayByParity(A []int) []int {
	l, r := 0, len(A)-1
	for l <= r {
		for l <= r && A[l]&1 == 0 {
			l++
		}
		for l <= r && A[r]&1 == 1 {
			r--
		}
		if l <= r {
			A[l], A[r] = A[r], A[l]
		}
	}
	return A
}

/*
	题目链接:
		https://leetcode-cn.com/problems/sort-array-by-parity/submissions/			按奇偶排序数组
*/

/*
	问题:
		这题目的答案可以是满足此条件的任何数组，假如加大一点难度，要求数组元素的相对位置不能改变呢？
			所谓相对位置就是: 我在你前面，那么操作后我还是在你前面。
			比如 [2 1 4 3]，那么操作后只能是 [2 4 1 3]，而不能是[4 2 1 3]
			(因为4、2的相对位置发生改变，4原本在2后面，现在跑到前面去了)
	解法:
		解法很简单，看上面的代码
*/

/*
	总结
	1. 这题目的双指针定义还可以定义在左右两边，然后采用向中心推进的方式。左指针遇到奇数停下，右指针遇到偶数停下，
		此时执行交换，之后再重复以上的不走，知道l>=r
*/
