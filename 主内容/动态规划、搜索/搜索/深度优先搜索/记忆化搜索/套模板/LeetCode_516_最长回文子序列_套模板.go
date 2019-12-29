package main


var LPS map[string]int // 保留已经得到的结果，该结构相当于一个备忘录

// 采用 记忆化搜索 最长回文子序列
func longestPalindromeSubseq(s string) int {
	/* 1. 进行一些预处理 */
	LPS = make(map[string]int)

	/* 2. 开始调用记忆化搜索函数，返回记忆化搜索结果 */
	return longestPalindromeSubseqExec(s)
}

// 记忆化搜索函数
func longestPalindromeSubseqExec(s string) int {
	/* 3. 判断是否需要返回结果以及进行一些剪枝  (特殊情况处理) */
	if len(s) <= 1 {
		return len(s)
	}

	// 如果该问题已经求解过了，那么直接返回结果
	if x, ok := LPS[s]; ok {
		return x
	}

	/* 4. 如果没求解，则继续调用记忆化搜索函数，得出结果  (一般情况处理) */
	ans := 0
	if s[0] == s[len(s)-1] {
		ans = longestPalindromeSubseqExec(s[1:len(s)-1]) + 2
	} else {
		a := longestPalindromeSubseqExec(s[:len(s)-1])
		b := longestPalindromeSubseqExec(s[1:])
		ans = max(a, b)
	}

	// 记录该问题的结果，加入备忘录
	LPS[s] = ans
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接:
		https://leetcode-cn.com/problems/longest-palindromic-subsequence/submissions/		最长回文子序列
*/

/*
	总结
	1. 由官方的结果来看，使用迭代的dp时空效率都高于使用递归的记忆化搜索。
			(dp: 44ms 10.8MB) (记忆化搜索: 376ms 59.7MB)
	2. 从代码的简洁度来看，dp的比记忆化搜索的简洁。
	3. 从设计难度上看(个人感觉)，记忆化搜索比dp容易设计、编码。
*/