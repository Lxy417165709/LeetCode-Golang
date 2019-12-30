package main

func findAndReplacePattern(words []string, pattern string) []string {
	ans := []string{}
	for i := 0; i < len(words); i++ {
		if judge(words[i], pattern) {
			ans = append(ans, words[i])
		}
	}
	return ans
}

func judge(s, p string) bool {
	if len(s) != len(p) {
		return false
	}
	sHashTo := make(map[uint8]uint8)
	pHashTo := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		sHashTo[s[i]] = p[i]
		pHashTo[p[i]] = s[i]
	}
	for i := 0; i < len(s); i++ {
		if sHashTo[s[i]] != p[i] || pHashTo[p[i]] != s[i] {
			return false
		}
	}
	return true
}

/*
	总结
	1. 这题建立映射关系就可以了。
	2. 这题和205同构字符串做法是一样的。
*/