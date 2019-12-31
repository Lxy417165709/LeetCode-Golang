package main

func countSegments(s string) int {
	wordFlag := false	// 单词标志
	ans := 0
	for i := 0; i < len(s); i++ {
		// 空格则跳过
		if s[i] == ' ' {
			wordFlag = false
			continue
		}
		// 遇见单词首字母则进行计数
		if wordFlag == false {
			ans++
			wordFlag = true
		}
	}
	return ans
}
