package main

const (
	inf = 1000000000
)

func balancedString(s string) int {
	hash := make(map[uint8]int)
	hash['Q'], hash['W'], hash['E'], hash['R'] = 0, 1, 2, 3
	wholeCount := make([]int, 4) // 字符串的字符信息
	for i := 0; i < len(s); i++ {
		wholeCount[hash[s[i]]]++
	}
	windowNeedCount := make([]int, 4) // 窗口所需要的字符信息
	for i := 0; i < len(windowNeedCount); i++ {
		windowNeedCount[i] = wholeCount[i] - len(s)/4
	}
	windowOwnCount := make([]int, 4) // 窗口当前拥有的字符信息

	first, last := 0, 0
	ans := inf
	for last < len(s) {
		windowOwnCount[hash[s[last]]]++
		for first < len(s) && isFull(windowNeedCount, windowOwnCount) {
			ans = min(ans, last-first+1)
			windowOwnCount[hash[s[first]]]--
			first++
		}
		last++
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
// 判断窗口内的字符 满足 窗口所需要的字符
func isFull(windowNeedCount []int, windowOwnCount []int) bool {
	for i := 0; i < len(windowNeedCount); i++ {
		if windowNeedCount[i] > windowOwnCount[i] {
			return false
		}
	}
	return true
}
/*
	总结
	1. 这题是滑动窗口的变形版。
	2. 这题和之前的滑动串口不太一样的就是:
			(1) 需要先获得窗口所需要的信息，在这里所需要的信息就是: 子串至少包含哪些字符，才能使字符串转为平衡字符串。
*/
