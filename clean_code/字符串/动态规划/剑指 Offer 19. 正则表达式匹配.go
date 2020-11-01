package 动态规划

func isMatch(s string, p string) bool {
	match := get2DSlice(len(s)+1, len(p)+1)
	match[0][0] = true
	for i := 0; i <= len(s); i++ {
		for t := 1; t <= len(p); t++ {
			if p[t-1] == '*' {
				isPreMatch := (i >= 1) && (p[t-2] == s[i-1] || p[t-2] == '.')
				match[i][t] = match[i][t-2] || isPreMatch && match[i][t-1] || isPreMatch && match[i-1][t]
			} else {
				match[i][t] = (i >= 1) && (s[i-1] == p[t-1] || p[t-1] == '.') && match[i-1][t-1]
			}
		}
	}
	return match[len(s)][len(p)]
}

func get2DSlice(rows, cols int) [][]bool {
	slice := make([][]bool, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]bool, cols)
	}
	return slice
}

/*
	总结:
		1. 这题的麻烦之处在于，字符串 s 为空，p不为空时还要进行处理。 (所以 i 的下标是从 0 开始)
*/
