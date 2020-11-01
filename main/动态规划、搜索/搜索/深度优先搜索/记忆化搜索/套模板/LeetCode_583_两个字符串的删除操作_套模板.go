package main

var isVisit map[string]int // 保留已经得到的结果，该结构相当于一个备忘录

// 记忆化搜索函数调用者
func minDistance(word1 string, word2 string) int {
	/* 1. 进行一些预处理 */
	isVisit = make(map[string]int)

	/* 2. 开始调用记忆化搜索函数，返回记忆化搜索结果 */
	return minDistanceExec(word1, word2)
}

// 记忆化搜索函数
func minDistanceExec(word1 string, word2 string) int {
	/* 3. 判断是否需要返回结果以及进行一些剪枝  (特殊情况处理) */
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}

	// 如果该问题已经求解过了，那么直接返回结果
	hashVal := hash(word1, word2)
	if x, ok := isVisit[hashVal]; ok {
		return x
	}

	/* 4. 如果没求解，则继续调用记忆化搜索函数，得出结果  (一般情况处理) */
	ans := 0
	if word1[len(word1)-1] == word2[len(word2)-1] {
		ans = minDistanceExec(word1[:len(word1)-1], word2[:len(word2)-1])
	} else {
		a := minDistanceExec(word1[:len(word1)-1], word2)
		b := minDistanceExec(word1, word2[:len(word2)-1])
		ans = min(a, b) + 1
	}

	// 记录该问题的结果，加入备忘录
	isVisit[hashVal] = ans
	return ans
}

// 由于备忘录的键值是 1 个字符串，而记忆化搜索函数需要 2 个字符串参数才能唯一标识一个子问题，
// 所以，这里采用哈希的方式，把两个参数进行哈希，生成一个键值来唯一的标识这个参数组合，
// 即: 用「1个字符串」 唯一标识 「1个子问题」。
func hash(a, b string) string {
	return a + "|" + b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接:
		https://leetcode-cn.com/problems/delete-operation-for-two-strings/		两个字符串的删除操作
*/

/*
	总结
	1. 在思路上，这题还可以利用最长公共子序列进行AC。
		过程就是:
		(1) 先获取最长公共子序列，
		(2)之后s1.length() + s2.length() - 2 * 最长公共子序列长度。
	2. 其实这题也可以使用迭代的dp进行AC，但是我更喜欢采用递归的记忆化搜索，因为不用考虑遍历方向而且终止条件也容易找。
	3. 这道题的升级版就是编辑距离了！
*/
