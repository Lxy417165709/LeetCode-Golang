package 字符串

import "fmt"

// todo: 这个滑动窗口的定义要确定好！
func lengthOfLongestSubstring(s string) int {
	// 1. 空返回。
	if len(s) == 0 {
		return 0
	}

	// 2. 初始化窗口。
	first, last := 0, 0
	isInWindow := make(map[byte]bool)

	// 3. 滑动。
	maxLength := 0
	for last < len(s) {
		// 3.1 左指针向右滑动，直到窗口内无重复元素。
		for first <= last && isInWindow[s[last]] {	// 为什么是 s[last]，而不是s[first]呢？
			isInWindow[s[first]] = false
			first++
		}

		// 3.2 右指针向右滑动，加入不重复的元素
		isInWindow[s[last]] = true

		// 3.3 获取最长长度。
		maxLength = max(maxLength, last-first+1)
		fmt.Println(s[first:last])

		last++
	}

	// 4. 返回。
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
