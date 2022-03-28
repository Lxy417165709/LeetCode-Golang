package 字符串

// --------------------------------------------------- 1. 回溯1(开始) ---------------------------------------------------

// generateParenthesis 生成括号对。
// 递归思路: 括号之间分为2种关系:
// 1. 包含关系。
// 2. 平行关系。
//
// 贴士:
// 1. 函数还能进行优化，比如添加缓存，避免重复计算。
func generateParenthesis(n int) []string {
	// 1. 递归出口。
	if n == 1 {
		return []string{"()"}
	}

	// 2. 获取缓存字符串集。 (缓存字符串集里面有重复的字符串)
	bufferStrings := make([]string, 0)
	for _, innerString := range generateParenthesis(n - 1) {
		bufferStrings = append(bufferStrings, "("+innerString+")")
	}
	for i := 1; i <= n-1; i++ {
		leftStrings := generateParenthesis(i)
		rightStrings := generateParenthesis(n - i)
		for _, leftString := range leftStrings {
			for _, rightString := range rightStrings {
				bufferStrings = append(bufferStrings, leftString+rightString)
			}
		}
	}

	// 3. 去重
	result := deduplicateString(bufferStrings)

	// 4. 返回。
	return result
}

// deduplicateString 字符串去重。
func deduplicateString(strings []string) []string {
	result := make([]string, 0)
	isExist := make(map[string]bool)
	for _, s := range strings {
		if isExist[s] {
			continue
		}
		isExist[s] = true
		result = append(result, s)
	}
	return result
}

// --------------------------------------------------- 1. 回溯1(结束) ---------------------------------------------------

// --------------------------------------------------- 2. 回溯2(开始) ---------------------------------------------------

// generateParenthesis 生成括号对。
// 递归思路:
// 1. 每个位置只能由 ( 或 ) 占领。
// 2. 只能追加字符时，如果原字符串的右括号大于左括号数时，一定不可能形成合法的括号对。
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	genParenthesis(&result, "", n, 0, 0)
	return result
}

// genParenthesis 生成括号对。
func genParenthesis(result *[]string, curString string, n, countOfLeftParentheses, countOfRightParentheses int) {
	// 1. 左括号数或右括号数大于给定数值，返回。
	if countOfLeftParentheses > n || countOfRightParentheses > n {
		return
	}

	// 2. 右括号数大于左括号数，不可能形成合法括号字符串。 (剪枝操作)
	if countOfRightParentheses > countOfLeftParentheses {
		return
	}

	// 3. 左右括号等于给定目标值，此时 curString 为合法括号字符串。
	if countOfLeftParentheses == countOfRightParentheses && countOfLeftParentheses == n {
		*result = append(*result, curString)
		return
	}

	// 4. 递归。
	genParenthesis(result, curString+"(", n, countOfLeftParentheses+1, countOfRightParentheses)
	genParenthesis(result, curString+")", n, countOfLeftParentheses, countOfRightParentheses+1)
}

// --------------------------------------------------- 2. 回溯2(结束) ---------------------------------------------------
