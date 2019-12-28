package main

var hasResult map[string]bool

func isMatch(s string, p string) bool {
	hasResult = make(map[string]bool)
	return solve(s, p)
}
func hash(s, p string) string {
	return s + "|" + p
}
func isMatchExec(s string, p string) bool {
	// 记忆化搜索
	key := hash(s, p)
	if x, ok := hasResult[key]; ok {
		return x
	}

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
	// 接下来是s,p都非空时
	ans := false

	// 这里表示s和p的末尾字符匹配
	if s[ends] == p[endp] || p[endp] == '.' {
		ans = isMatchExec(s[:ends], p[:endp])
	} else {
		// 这里表示p的末尾字符是'*'时
		if p[endp] == '*' {
			// 判断s的末尾字符和p的倒数第二个字符是否匹配
			if p[endp-1] == s[ends] || p[endp-1] == '.' {
				ans = isMatchExec(s, p[:endp]) || isMatchExec(s[:ends], p) || isMatchExec(s, p[:endp-1])
			} else {
				ans = isMatchExec(s, p[:endp-1])
			}
		}
	}
	hasResult[key] = ans
	return ans
}
/*
	总结
	1. 这道题想要做出来，就需要理清先正则表达式的规则。
*/

