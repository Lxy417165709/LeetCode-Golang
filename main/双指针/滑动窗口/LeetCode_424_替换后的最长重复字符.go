package 滑动窗口

import "fmt"

func characterReplacement(s string, k int) int {
	count := make(map[uint8]int) // 用来记录窗口内字符出现次数
	maxCount := 0                // 表示到目前为止，窗口内的相同字符最多出现次数
	first, last := 0, 0
	maxLength := 0
	for last < len(s) {
		count[s[last]]++
		maxCount = max(maxCount, count[s[last]])
		for first < len(s) && last-first+1-maxCount > k {
			count[s[first]]--
			first++
		}
		maxLength = max(maxLength, last-first+1)
		last++
	}
	// 代码中的maxCount具有不减的属性， 只要没有出现更大的maxCount, 就不会有更长的重复子串; 因此left向后移动时不需要更新maxCount的值
	// 这是力扣有个人的解释。
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
		https://leetcode-cn.com/problems/longest-repeating-character-replacement/		替换后的最长重复字符
*/

/*
	总结
	1. 这题目使用了滑动窗口算法
	2. 在这题纠结了好久，就是first前移时，maxCount是否需要改变。上面是有个大佬的解释.
	3. 自己理解就是: 假如出现的maxCount == 6，那么之后如果要形成更长的字符串，maxCount必须＞6。
*/
