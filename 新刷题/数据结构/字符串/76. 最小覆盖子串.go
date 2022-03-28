package 字符串

import "fmt"

// ----------------------------------- 递归解法(超时，复杂度 O(n^2)) -----------------------------------

var cache map[string]interface{}

func minWindow(s string, t string) string {
	cache = make(map[string]interface{})
	return getMinString(s, t)
}

func getMinString(s string, t string) string {

	if s == "" {
		return ""
	}

	if minString, ok := getMinStringFromCache(s); ok {
		return minString
	}

	if !getSetIsContain(s, t) {
		return ""
	}

	head := s[1:]
	tail := s[:len(s)-1]

	minString1 := getMinString(head, t)
	minString2 := getMinString(tail, t)

	key := fmt.Sprintf("min_string:%s", s)
	cache[key] = compareOutMinString(s, minString1, minString2)
	return cache[key].(string)
}

func compareOutMinString(s, minString1, minString2 string) string {
	if len(minString1) == 0 && len(minString2) == 0 {
		return s
	} else if len(minString1) == 0 {
		return minString2
	} else if len(minString2) == 0 {
		return minString1
	} else if len(minString1) == len(minString2) {
		return minString1 // 任意返回。题目说只有一个满足条件。
	} else if len(minString1) > len(minString2) {
		return minString2
	} else {
		return minString1
	}
}

func getMinStringFromCache(s string) (string, bool) {
	key := fmt.Sprintf("min_string:%s", s)
	val, ok := cache[key]
	if !ok {
		return "", false
	}
	return val.(string), true
}

func getSetIsContain(s string, t string) bool {
	key := fmt.Sprintf("is_contain:%s_%s", s, t)
	val, ok := cache[key]
	if ok {
		return val.(bool)
	}
	m1, m2 := getSetCountMap(s), getSetCountMap(t)
	cache[key] = isContain(m1, m2)
	return cache[key].(bool)
}

func getSetCountMap(s string) map[byte]int {
	key := fmt.Sprintf("count_map:%s", s)
	val, ok := cache[key]
	if ok {
		return val.(map[byte]int)
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	cache[key] = m
	return m
}

// ----------------------------------- 二分解法(场景不适用) -----------------------------------

// 注意，这题不能用二分，因为:
// 1. 在左右两侧找到 minString 后，minString 依旧可能在字符串中间两侧。
// 2. 在左侧找到 minString 后，minString 依旧可能在字符串中间两侧。
// 3. ...
// 4. 无论如何，minString 依旧可能在字符串中间两侧。

//
//func minWindow(s string, t string) string {
//	cache = make(map[string]interface{})
//	return getMinString(s, t)
//}
//
//// 错误二分法
//func getMinString(s, t string) string {
//	fmt.Println(s, t)
//	if len(s) == 0 {
//		return ""
//	}
//	if len(s) == 1 {
//		if s == t {
//			return s
//		}
//		return ""
//	}
//
//	if minString, ok := getMinStringFromCache(s); ok {
//		return minString
//	}
//
//	if !getSetIsContain(s, t) {
//		return ""
//	}
//
//	left, right := 0, len(s)-1
//	mid := (left + right) / 2
//	part1, part2 := s[:mid+1], s[mid+1:]
//	contain1 := getSetIsContain(part1, t)
//	contain2 := getSetIsContain(part2, t)
//	key := fmt.Sprintf("min_string:%s", s)
//	if contain1 && contain2 {
//		a, b := getMinString(part1, t), getMinString(part2, t)
//		cache[key] = getShortestString(a, b)
//	} else if !contain1 && !contain2 {
//		a, b := getMinString(s[1:], t), getMinString(s[:len(s)-1], t)
//		if len(a) == 0 && len(b) == 0 {
//			cache[key] = s
//		} else if len(a) != 0 && len(b) != 0 {
//			cache[key] = getShortestString(a, b)
//		} else if len(a) != 0 {
//			cache[key] = a
//		} else {
//			cache[key] = b
//		}
//	} else if contain1 {
//		a, b := getMinString(s[1:], t), getMinString(s[:len(s)-1], t)
//		if len(a) == 0 && len(b) == 0 {
//			cache[key] = s
//		} else if len(a) != 0 && len(b) != 0 {
//			cache[key] = getShortestString(a, b)
//		} else if len(a) != 0 {
//			cache[key] = a
//		} else {
//			cache[key] = b
//		}
//
//		cache[key] = getMinString(part1, t)
//	} else {
//		cache[key] = getMinString(part2, t)
//	}
//	return cache[key].(string)
//}
//

// ----------------------------------- 滑动窗口(通过，复杂度 O(n)) -----------------------------------

// minWindow 最小覆盖子串。 (统一滑动窗口)
func minWindow(s string, t string) string {
	// 1. 空返回。
	if len(s) == 0 {
		return ""
	}

	// 2. 初始化窗口。
	windowMap := make(map[byte]int)
	left, right := 0, 0 // [left, right] 窗口内的元素满足要求。

	// 3. 构建子字符串映射。
	subStringMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		subStringMap[t[i]]++
	}

	// 4. 滑动
	minLengthString := s + s // 特殊赋值，无最小覆盖子串时，处理后 minLengthString 不变。
	for right < len(s) {
		// 4.1 加入尾元素。
		windowMap[s[right]]++

		// 4.2 左指针向右滑动，直到第一次不满足覆盖条件。
		for left <= right && isContain(windowMap, subStringMap) {
			// 4.2.1 获取最短长度。
			minLengthString = getShortestString(minLengthString, s[left:right+1])

			// 4.2.2 将最左的字符移出窗口。
			windowMap[s[left]]--

			// 4.2.3 处理窗口最左的下一字符。
			left++
		}

		// 4.3 右指针向右滑动，处理下一个字符。
		right++
	}

	// 5. 判断是否存在最小覆盖子串，不存在返回空。
	if minLengthString == s+s {
		return ""
	}

	// 6. 返回。
	return minLengthString
}

// isContain 判断 m1 是否包括 m2。
func isContain(m1, m2 map[byte]int) bool {
	for key, val := range m2 {
		if val > m1[key] {
			return false
		}
	}
	return true
}

// getShortestString 获取 a, b 中最短字符串。
func getShortestString(a, b string) string {
	if len(a) > len(b) {
		return b
	}
	return a
}
