package main

// 记忆化搜索AC
func checkValidString(s string) bool {
	isVisit = make(map[int]bool)
	return checkValidStringExec(s, len(s)-1, 0, 0)
}

// 这里主要是为了把3个变量哈希为1个变量，方便map存放
func hash(a, b, c int) int {
	return (a << 20) + (b << 10) + c
}

var isVisit map[int]bool // 记忆化
// left,right 表示此时的左右括号数量
func checkValidStringExec(s string, nowIndex int, left, right int) bool {
	if nowIndex == -1 {
		return left == right
	}
	if left > right {
		return false
	}
	hashNumber := hash(nowIndex, left, right)
	if x, ok := isVisit[hashNumber]; ok {
		return x
	}
	lastChar := s[nowIndex]
	ans := false
	if lastChar == '(' || lastChar == '*'{
		ans = ans || checkValidStringExec(s, nowIndex-1, left+1, right)
	}
	if lastChar == ')' || lastChar == '*'{
		ans = ans || checkValidStringExec(s, nowIndex-1, left, right+1)
	}
	if lastChar == '*' {
		ans = ans || checkValidStringExec(s, nowIndex-1, left, right)
	}
	isVisit[hashNumber] = ans
	return ans
}
/*
	总结
	1. 刚刚开始我采用了DP做，但是没有做出来...因为思路已经错了。
	2. 这次采用记忆化搜索做，直接AC了，我以为会超时...
	3. 官方有人用栈做的，争取我也用栈做一下！
*/