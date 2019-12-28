package 脑筋急转弯

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接:
		https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/comments/		最长特殊序列 Ⅰ
*/
/*
	总结
	1. 思路:
		接下来我默认len(a)>=len(b)，且默认对a进行操作。
		(1) a,b 长度不等，显然整串a作为特殊子序列，b肯定不能组成该特殊子序列，
			于是返回max(len(a),len(b))
		(2) a,b 长度相等，假如a,b相等，那么从a中取出的任何一个子序列，b都有，于是不存在特殊序列。
			假如a,b不等，那么我继续把整串a作为特殊子序列，b肯定也不能组成该特殊子序列。
*/
