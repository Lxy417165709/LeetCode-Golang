package main

var inf int				// 无穷大
var amount map[int]int 	// 保留已经得到的结果，该结构相当于一个备忘录

// 记忆化搜索函数调用者
func getMoneyAmount(n int) int {
	/* 1. 进行一些预处理 */
	amount = make(map[int]int)
	inf = 100000000000

	/* 2. 开始调用记忆化搜索函数，返回记忆化搜索结果 */
	return getMoneyAmountExec(1, n)
}

// 记忆化搜索函数
func getMoneyAmountExec(l, r int) int {
	/* 3. 判断是否需要返回结果以及进行一些剪枝  (特殊情况处理) */
	if l >= r {
		return 0
	}

	// 如果该问题已经求解过了，那么直接返回结果
	hashNumber := hash(l,r)
	if x, ok := amount[hashNumber]; ok {
		return x
	}

	/* 4. 如果没求解，则继续调用记忆化搜索函数，得出结果  (一般情况处理) */
	ans := inf
	for i := l; i <= r; i++ {
		left := getMoneyAmountExec(l, i-1)
		right := getMoneyAmountExec(i+1, r)
		ans = min(ans, max(left, right)+i)
	}

	// 记录该问题的结果，加入备忘录
	amount[hashNumber] = ans
	return ans
}

// 由于备忘录的键值是 1 个整数，而记忆化搜索函数需要 2 个整数参数才能唯一标识一个子问题，
// 所以，这里采用哈希的方式，把两个参数进行哈希，生成一个键值来唯一的标识这个参数组合，
// 即: 用「1个整数」 唯一标识 「1个子问题」。
func hash(l,r int) int{
	off := 10
	return (r << off) | l
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
