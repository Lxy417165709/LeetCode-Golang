package main

var off uint8			// 哈希偏移量
var inf int				// 无穷大
var amount map[int]int 	// 区间 [l, r] 所需的最小花费

// 记忆化搜索 解决 猜数字大小 II
func getMoneyAmount(n int) int {
	amount = make(map[int]int)
	off = 10
	inf = 100000000000
	return getMoneyAmountExec(1, n)
}

func getMoneyAmountExec(l, r int) int {
	// l > r 时，表示区间非法，于是返回0
	// l == r 时，说明这个数组就是我们要选的数字，也是返回0
	if l >= r {
		return 0
	}
	hashNumber := (r << off) | l
	if x, ok := amount[hashNumber]; ok {
		return x
	}
	ans := inf

	// 这个循环表示选择哪个数
	// i可以从(l+r)/2开始的，这应该需要用到贪心，但是我不知道怎么证明，所以这里还是从l开始。
	for i := l; i <= r; i++ {
		left := getMoneyAmountExec(l, i-1)
		right := getMoneyAmountExec(i+1, r)
		// 是否为目标数只有2个状态:
		// 		(1) 选择的这个数是目标数: 	是目标数的话，选取代价为0。
		//		(2) 选择的这个数不是目标数:	那么就根据这个数，向其左右两边找，返回花费最大的，再加上选取代价 i。
		// 所以，下面也可以写为: ans = min(ans, max(0, max(left, right)+i))
		//
		// 为什么left和right时返回最大的？因为我们要保证最坏情况时，我们的钱足够。 (目标数是哪个数我们是不确定的，主动权不在我们自己)
		// 为什么是ans选取的是最小的呢？  因为选择哪个数的主动权在我们自己。
		ans = min(ans, max(left, right)+i)
	}
	amount[hashNumber] = ans
	return amount[hashNumber]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接:
		https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/submissions/		猜数字大小 II
*/

/*
	总结
	1. 这题使用了记忆化搜索，个人认为 记忆化搜索 = 递归 + 备忘录。
	2. 其实也可以使用DP，不过就要考虑遍历的顺序。
*/
