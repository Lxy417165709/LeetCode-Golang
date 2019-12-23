package main

func lenLongestFibSubseq(A []int) int {
	ans := 0
	for i := 0; i < len(A); i++ {
		for t := 0; t < i; t++ {
			// 获取以 .......A[t],A[i] 为结尾的最长斐波那契子序列的长度
			ans = max(ans, getFibSubLength(A, 0, t-1, A[i], A[t]))
		}
	}
	if ans <= 2 {
		ans = 0
	}
	return ans
}

// 这函数的意思是: 在A[left: right+1] 中找最长斐波那契子序列，其中末项是lastItem,倒数第二项是lastItem2
func getFibSubLength(A []int, left, right, lastItem, lastItem2 int) int {
	length := 2
	curItem := lastItem - lastItem2
	for true {
		// 采用二分查找
		right = search(A, left, right, curItem) - 1
		// 因为search函数返回值-1表示没找到，所以这里-1-1 == -2时表示没找到
		if right == -2 {
			break
		}
		length++
		// 获得下一个查找的目标值
		tmp := curItem
		curItem = lastItem2 - curItem
		lastItem2 = tmp
	}
	return length
}

// 二分查找，找到返回索引，否则返回 -1
func search(A []int, l, r, target int) int {
	for l <= r {
		mid := (l + r) / 2
		if A[mid] == target {
			return mid
		} else {
			if A[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	总结
	1. 这题之前采用动态规划做的时候，方向做错了，于是WA..然后发现数组是递增的于是想到了用二分法做，
		以上是AC代码。
	2. 这题目因为没有重复元素，所以其实二分查找的可以不使用边界，即搜索整个数组就可以了。

*/
