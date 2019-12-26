package 脑筋急转弯
// 计算字符串有多少种不同的字符
func calculateDifferentChar(s string) int {
	count := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	return len(count)
}

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	count := make(map[string]int) // 记录符合要求的子串的出现次数
	ans := 0
	for i := minSize - 1; i < len(s); i++ {
		if calculateDifferentChar(s[i+1-minSize:i+1]) <= maxLetters {
			count[s[i+1-minSize:i+1]]++
			ans = max(ans, count[s[i+1-minSize:i+1]])
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	总结
	1. 这道题的maxSize是没用的，因为如果maxSize的串满足要求，minSize的串也一定会满足要求，所以我们只需要求
	   minSize长度的子串是否符合要求就可以了。
	2. 这道题还可以用滚动哈希进行优化，这样就不用每次都把minSize长度的子串都扫一遍了。
	3. 比赛的时候感觉时间复杂度是O(n*minSize)会超时...
	   比赛后看了下限制条件，
			1 <= s.length <= 10^5
			1 <= maxLetters <= 26
			1 <= minSize <= maxSize <= min(26, s.length)
	   发现时间复杂度最大是O(26n)...怪不得提交的时候可以通过。
*/
