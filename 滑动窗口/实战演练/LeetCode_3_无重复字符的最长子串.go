package main

// 计算无重复字符的最长子串长度
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	count := make(map[uint8]int) // 用来记录窗口中的每个字符出现了多少次
	first, last := 0, 0          // 窗口的左右边界
	for last < len(s) {
		count[s[last]]++
		for first <= last && count[s[last]] >= 2 {
			count[s[first]]--
			first++
		}
		maxLength = max(maxLength, last-first+1)
		last++
	}
	return maxLength
}

// 优化后的代码 (first不再一次跳一位，而是直接跳到对应的位置，使得窗口没有重复字符)
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	index := make(map[uint8]int) // 记录字符在字符串中的索引
	first, last := 0, 0			 // 窗口的左右边界
	for last < len(s) {
		// ok 表示字符出现过
		// idx >= first 表示该字符在窗口内
		if idx, ok := index[s[last]]; ok {
			if idx >= first {
				first = idx + 1
			}
		}
		index[s[last]] = last
		maxLength = max(maxLength, last-first+1)
		last++
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


/*
	题目链接:
		https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/submissions/		无重复字符的最长子串
*/

/*
	总结
	1. 这题还可以优化的，优化思路: 由于first每次都只是移动一位，但其实我们可以记录s[last]之前出现的位置，那么如果s[last]
       出现重复，first可以直接跳到s[last]的前一个位置上，这样可以优化一些时间效率。
	2. 滑动窗口更类似一个模板，而加上优化后，这个模板可能不太好识别..
*/
