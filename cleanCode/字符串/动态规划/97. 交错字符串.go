package 动态规划

// ------------------------------ 暴力搜索 ------------------------------
// 执行用时：1072 ms, 在所有 Go 提交中击败了 8.09%   的用户
// 内存消耗：2.1 MB,  在所有 Go 提交中击败了 100.00% 的用户
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	if len(s3) == 0 {
		return true
	}
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}
	switch {
	case s1[0] == s3[0] && s2[0] == s3[0]:
		return isInterleave(s1[1:], s2, s3[1:]) || isInterleave(s1, s2[1:], s3[1:])
	case s1[0] == s3[0]:
		return isInterleave(s1[1:], s2, s3[1:])
	case s2[0] == s3[0]:
		return isInterleave(s1, s2[1:], s3[1:])
	}
	return false
}

// ------------------------------ 动态规划 ------------------------------
// 执行用时：0 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度: O(n^2)
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	canForm := get2DSlice(len(s1)+1, len(s2)+1) // 表示 s1[:i] 与 s2[:t] 能否交错组成 s3[:i+t]
	canForm[0][0] = true
	for i := 1; i <= len(s1); i++ {
		canForm[i][0] = s1[i-1] == s3[i-1] && canForm[i-1][0]
	}
	for t := 1; t <= len(s2); t++ {
		canForm[0][t] = s2[t-1] == s3[t-1] && canForm[0][t-1]
	}
	for i := 1; i <= len(s1); i++ {
		for t := 1; t <= len(s2); t++ {
			canForm[i][t] = s1[i-1] == s3[i+t-1] && canForm[i-1][t] || s2[t-1] == s3[i+t-1] && canForm[i][t-1]

			// 也可以像下面这样写:
			// switch {
			// case s1[i-1] == s3[i+t-1] && s2[t-1] == s3[i+t-1]:
			//	 canForm[i][t] = canForm[i-1][t] || canForm[i][t-1]
			// case s1[i-1] == s3[i+t-1]:
			//	 canForm[i][t] = canForm[i-1][t]
			// case s2[t-1] == s3[i+t-1]:
			//	 canForm[i][t] = canForm[i][t-1]
			// }
		}
	}
	return canForm[len(s1)][len(s2)]
}

func get2DSlice(rows, column int) [][]bool {
	slice := make([][]bool, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]bool, column)
	}
	return slice
}

/*
	题目链接: https://leetcode-cn.com/problems/interleaving-string/
*/
