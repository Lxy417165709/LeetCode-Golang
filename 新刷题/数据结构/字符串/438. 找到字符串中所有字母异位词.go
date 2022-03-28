package 字符串

// findAnagrams 找到字符串中所有字母异位词的起始索引。
func findAnagrams(s string, p string) []int {
	// 1. 构建映射。
	subStringMap := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		subStringMap[p[i]]++
	}

	// 2. 初始化窗口。
	left, right := 0, 0
	windowMap := make(map[byte]int)

	// 3. 滑动。
	result := make([]int, 0)
	for right < len(s) {
		windowMap[s[right]]++
		for left <= right && isContain(windowMap, subStringMap) {
			if isMapEqual(windowMap, subStringMap) {
				result = append(result, left)
			}
			windowMap[s[left]]--
			left++
		}
		right++
	}

	// 4. 返回。
	return result
}

func isMapEqual(m1, m2 map[byte]int) bool {
	for key, value := range m1 {
		if m2[key] != value {
			return false
		}
	}
	for key, value := range m2 {
		if m1[key] != value {
			return false
		}
	}
	return true
}
