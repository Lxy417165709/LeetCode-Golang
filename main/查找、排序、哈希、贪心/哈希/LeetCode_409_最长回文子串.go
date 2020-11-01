package main

func longestPalindrome(s string) int {
	count := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	length := 0
	oddFlag := false
	for _, v := range count {
		if v%2 == 1 {
			length += v - 1
			oddFlag = true
		} else {
			length += v
		}
	}
	if oddFlag {
		length++
	}
	return length
}

// 更简洁的写法
func longestPalindrome(s string) int {
	count := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	length := 0
	for _, v := range count {
		length += v / 2 * 2
	}
	if length < len(s) {
		length++
	}
	return length
}

/*
	总结
	1. 这题我的思路是: 记录字符出现的个数，再遍历一次，如果偶数则全部添加入回文串，如果是
						奇数就把个数-1加入回文串，最后再判断该字符串中是否有个数为奇数的字符，
						如果有就+1(做回文字符串的中心)，否则不加(回文字符串长度为偶数)。
*/
