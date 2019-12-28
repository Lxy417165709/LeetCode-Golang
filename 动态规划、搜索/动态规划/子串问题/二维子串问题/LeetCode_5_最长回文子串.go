package 二维子串问题


// 动态规划 解决最长回文子串
// dp[i][t] 表示: s[i][t]是否为回文串
func longestPalindrome(s string) string {
	dp := [1050][1050]bool{} // s的最大长度是1000,
	longestPalindromeString := ""
	for length := 1; length <= len(s); length++ {
		for begin := 0; begin+length-1 < len(s); begin++ {
			end := begin + length - 1
			if length <= 2 {
				dp[begin][end] = s[begin] == s[end]
			} else {
				dp[begin][end] = s[begin] == s[end] && dp[begin+1][end-1]
			}
			if dp[begin][end] == true && length > len(longestPalindromeString) {
				longestPalindromeString = s[begin : end+1]
			}
		}
	}
	return longestPalindromeString
}

/*
	题目链接: https://leetcode-cn.com/problems/longest-palindromic-substring/submissions/
*/
