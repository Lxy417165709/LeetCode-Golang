package main

/*
	给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
*/

// 判断 s 是否为 t 的子序列。	( 时间复杂度为O(n+m) )
func isSubsequence(s string, t string) bool {
	indexOfS, indexOfT := 0, 0
	for indexOfS < len(s) && indexOfT < len(t) {
		if s[indexOfS] == t[indexOfT] {
			indexOfS++
		}
		indexOfT++
	}
	return indexOfS == len(s)
}

/*
	题目链接:
		https://leetcode-cn.com/problems/is-subsequence/submissions/		判断子序列
*/

/*
	总结
	1. 这题目除了采用上面的方法做外，还可以使用哈希的方法喔。
*/
