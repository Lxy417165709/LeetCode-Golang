package main

func isMatch(s string, p string) bool {
	isVisit = make(map[string]bool)
	return isMatchExec(s, p)
}

var isVisit map[string]bool

func hash(s, p string) string {
	return s + "|" + p
}

func isMatchExec(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(p) == 0 {
		return false
	}
	if len(s) == 0 {
		for i := 0; i < len(p); i++ {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}
	hashStr := hash(s, p)
	if x, ok := isVisit[hashStr]; ok {
		return x
	}

	lastOfS, lastOfP := len(s)-1, len(p)-1
	ans := false
	if s[lastOfS] == p[lastOfP] || p[lastOfP] == '?' {
		ans = isMatchExec(s[:lastOfS], p[:lastOfP])
	} else {
		if p[lastOfP] == '*' {

			for i := len(s); i >= 0 && ans == false; i-- {
				ans = isMatchExec(s[:i], p[:lastOfP])
			}
		}
	}

	isVisit[hashStr] = ans
	return ans
}

/*
	总结
	1. 这个版本超时了。
*/
