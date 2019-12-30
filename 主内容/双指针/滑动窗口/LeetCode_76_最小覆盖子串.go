package 滑动窗口

func minWindow(s string, t string) string {
	first, last := 0, 0
	windowMap := make(map[uint8]int)	// 窗口内的字符映射, key是字符, value是出现次数
	tMap := make(map[uint8]int)			// t的字符映射, key是字符, value是出现次数

	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	ans := s + "0" // 为了判断是否存在答案，我采用这样的初始化ans
	for last < len(s) {
		windowMap[s[last]]++
		for first < len(s) && judge(windowMap, tMap) {
			length := last - first + 1
			if length < 0 {
				return ""
			}
			if length < len(ans) {
				ans = s[first : last+1]
			}
			windowMap[s[first]]--
			first++
		}
		last++
	}
	// 表示没有答案
	if len(ans) == len(s)+1 {
		return ""
	}
	return ans
}

// 判断窗口内是否符合要求
func judge(windowMap, tMap map[uint8]int) bool {
	for k, v := range tMap {
		if windowMap[k] < v {
			return false
		}
	}
	return true
}

/*
	总结
	1. 这道题直接套框架~
	2. 刚开始还想利用位运算做一个小容量的map,但是做到一半才发现了问题，就是位运算的map只能记录是否出现，
		而不能记录出现多少次。
*/
