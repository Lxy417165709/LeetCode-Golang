package main

var isVisit map[int]bool // 保留已经得到的结果，该结构相当于一个备忘录

// 记忆化搜索函数调用者
func checkValidString(s string) bool {
	/* 1. 进行一些预处理 */
	isVisit = make(map[int]bool)

	/* 2. 开始调用记忆化搜索函数，返回记忆化搜索结果 */
	return checkValidStringExec(s, len(s)-1, 0, 0)
}

// 记忆化搜索函数调用者
// s[: nowIndex+1]为当前处理的字符串
// left, right 表示此时的左右括号数量
func checkValidStringExec(s string, nowIndex int, left, right int) bool {
	/* 3. 判断是否需要返回结果以及进行一些剪枝  (特殊情况处理) */
	if nowIndex == -1 {
		return left == right
	}
	if left > right {
		return false
	}

	// 如果该问题已经求解过了，那么直接返回结果
	hashNumber := hash(nowIndex, left, right)
	if x, ok := isVisit[hashNumber]; ok {
		return x
	}

	/* 4. 如果没求解，则继续调用记忆化搜索函数，得出结果  (一般情况处理) */
	ans := false
	lastChar := s[nowIndex]
	if lastChar == '(' || lastChar == '*'{
		ans = ans || checkValidStringExec(s, nowIndex-1, left+1, right)
	}
	if lastChar == ')' || lastChar == '*'{
		ans = ans || checkValidStringExec(s, nowIndex-1, left, right+1)
	}
	if lastChar == '*' {
		ans = ans || checkValidStringExec(s, nowIndex-1, left, right)
	}

	// 记录该问题的结果，加入备忘录
	isVisit[hashNumber] = ans
	return ans
}

// 这里主要是为了把3个变量哈希为1个变量，方便map存放
func hash(a, b, c int) int {
	return (a << 20) + (b << 10) + c
}

/*
	总结
	1. 刚刚开始我采用了DP做，但是没有做出来...因为思路已经错了。
	2. 这次采用记忆化搜索做，直接AC了，我以为会超时...
	3. 官方有人用栈做的，争取我也用栈做一下！
*/