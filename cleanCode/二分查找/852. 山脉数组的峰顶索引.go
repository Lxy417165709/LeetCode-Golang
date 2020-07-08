package 二分查找

func peakIndexInMountainArray(A []int) int {
	left, right := 1, len(A)-2
	for left <= right {
		mid := left + (right-left)/2
		if A[mid] == A[mid+1] {
			panic("题目出错")
		}
		if A[mid] > A[mid+1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func peakIndexInMountainArray(A []int) int {
	left, right := 1, len(A)-2
	for left <= right {
		mid := left + (right-left)/2
		if A[mid] == A[mid-1] {
			panic("题目出错")
		}
		if A[mid] > A[mid-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

/*
	题目链接: https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/
	总结
		1. 刚刚开始已经想出了O(n)的做法，也意识到可能可以二分。
		2. 看了题解后确定了二分，之后就写出来AC了。
		3. 上面的两种二分都能AC。
*/
