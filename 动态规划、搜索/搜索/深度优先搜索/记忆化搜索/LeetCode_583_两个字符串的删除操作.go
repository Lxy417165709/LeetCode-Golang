package main

// 函数入口
func minDistance(word1 string, word2 string) int {
	isVisit = make(map[string]int)
	return minDistanceExec(word1, word2)
}

var isVisit map[string]int	// 保存已经得到的结果

func minDistanceExec(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	hashVal := hash(word1, word2)
	if x, ok := isVisit[hashVal]; ok {
		return x
	}
	ans := 0
	// 如果相同，说明不用删除
	// 如果不同，那么就删除其中一个字符串的最后的字符，再进行向下求解
	if word1[len(word1)-1] == word2[len(word2)-1] {
		ans = minDistanceExec(word1[:len(word1)-1], word2[:len(word2)-1])
	} else {
		a := minDistanceExec(word1[:len(word1)-1], word2)	// 删除word1的最后一个字符
		b := minDistanceExec(word1, word2[:len(word2)-1])	// 删除word2的最后一个字符
		ans = min(a, b) + 1
	}
	isVisit[hashVal] = ans
	return ans
}
// 哈希，需要保证哈希值唯一
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

