package 滑动窗口

func minWindow(s string, t string) string {

	ans := s + "0" // 为了判断是否存在答案，我采用这种方式初始化ans

	/* 1. 初始化窗口数据结构，用于记录窗口内的信息 */
	windowMap := make(map[uint8]int)	// 窗口内的字符映射, key是字符, value是出现次数
	tMap := make(map[uint8]int)		// t的字符映射, key是字符, value是出现次数
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	first, last := 0, 0

	for last < len(s) {
		/* 2. 把 last 指向的元素加入窗口 */
		windowMap[s[last]]++

		/* 3. 判断当前窗口内的元素是否符合条件 */
		for first < len(s) && judge(windowMap, tMap) {
			// 如果题目求的是最小值，则符合题意条件时进入函数体

			/* 4.a 在这写更新窗口最小值的代码 */
			length := last - first + 1
			if length < 0 {
				return ""
			}
			if length < len(ans) {
				ans = s[first : last+1]
			}

			/* 5. 把 first 指向的元素移出窗口 */
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

// 判断窗口内字符串是否符合要求
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
