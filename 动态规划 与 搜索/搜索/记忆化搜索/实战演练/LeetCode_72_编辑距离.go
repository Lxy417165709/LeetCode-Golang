package main

// 函数入口
func minDistance(word1 string, word2 string) int {
	isVisit = make(map[string]int)
	return minDistanceExec(word1, word2)
}

var isVisit map[string]int // 保存已经得到的结果

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
		// 与 "LeetCode_583_两个字符串的删除操作.go"相比，就这里修改了
		a := minDistanceExec(word1[:len(word1)-1], word2)                // 删除word1的最后一个字符
		b := minDistanceExec(word1[:len(word1)-1], word2[:len(word2)-1]) // 替换word1的最后一个字符为word2的最后一个字符
		c := minDistanceExec(word1, word2[:len(word2)-1])                // 在word1后添加word2的最后一个字符
		ans = min(a, b, c) + 1
	}
	isVisit[hashVal] = ans
	return ans
}

// 哈希，需要保证哈希值唯一
func hash(a, b string) string {
	return a + "|" + b
}

// 这里我重写了min函数，让它可以计算n个参数的最小值
func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a, b := arr[0], min(arr[1:]...)
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
	1. 这是采用记忆化搜索实现的编辑距离...但是时空效率让我想哭。
	2. 小伙伴可以比较一下和该目录下"LeetCode_583_两个字符串的删除操作.go"的差别。
*/
