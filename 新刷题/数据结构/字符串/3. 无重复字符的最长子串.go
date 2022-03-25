package 字符串

import "github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"

func lengthOfLongestSubstring(s string) int {
	// 1. 空返回。
	if len(s) == 0 {
		return 0
	}

	// 2. 初始化。
	first, last := 0, 0 // 滑动窗口定义: [first, last) 窗口内的元素不重复。
	maxLength := 0
	isInWindow := make(map[byte]bool)

	// 3. 滑动。
	for last < len(s) {
		// 3.1 左指针向右滑动，确保窗口内无新元素。 (如果窗口中有新元素，那加入新元素后，窗口就有2个新元素，与定义不符)
		for first <= last && isInWindow[s[last]] {
			isInWindow[s[first]] = false
			first++
		}

		// 3.2 右指针向右滑动，加入新元素。
		isInWindow[s[last]] = true
		last++

		// 3.3 获取最长长度。
		maxLength = math_util.Max(maxLength, last-first)
	}

	// 4. 返回。
	return maxLength
}

