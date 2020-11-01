package main

var hasResult map[string]bool // 保留已经得到的结果，该结构相当于一个备忘录

// 记忆化搜索函数调用者
func isMatch(s string, p string) bool {
	/* 1. 进行一些预处理 */
	hasResult = make(map[string]bool)

	/* 2. 开始调用记忆化搜索函数，返回记忆化搜索结果 */
	return isMatchExec(s, p)
}

// 记忆化搜索函数
func isMatchExec(s string, p string) bool {
	/* 3. 判断是否需要返回结果以及进行一些剪枝  (特殊情况处理) */
	if s == p {
		return true
	}
	if p == "" {
		return s == ""
	}
	ends, endp := len(s)-1, len(p)-1
	if s == "" {
		if p[endp] == '*' {
			return isMatchExec(s, p[:endp-1])
		}
		return false
	}

	// 如果该问题已经求解过了，那么直接返回结果
	key := hash(s, p)
	if x, ok := hasResult[key]; ok {
		return x
	}

	/* 4. 如果没求解，则继续调用记忆化搜索函数，得出结果  (一般情况处理) */
	ans := false
	if s[ends] == p[endp] || p[endp] == '.' {
		ans = isMatchExec(s[:ends], p[:endp])
	} else {
		if p[endp] == '*' {
			if p[endp-1] == s[ends] || p[endp-1] == '.' {
				ans = isMatchExec(s, p[:endp]) || isMatchExec(s[:ends], p) || isMatchExec(s, p[:endp-1])
			} else {
				ans = isMatchExec(s, p[:endp-1])
			}
		}
	}

	// 记录该问题的结果，加入备忘录
	hasResult[key] = ans
	return ans
}

// 由于备忘录的键值是 1 个字符串，而记忆化搜索函数需要 2 个字符串参数才能唯一标识一个子问题，
// 所以，这里采用哈希的方式，把两个参数进行哈希，生成一个键值来唯一的标识这个参数组合，
// 即: 用「1个字符串」 唯一标识 「1个子问题」。
func hash(s, p string) string {
	return s + "|" + p
}

/*
	总结
	1. 这道题想要做出来，就需要理清先正则表达式的规则。
*/
