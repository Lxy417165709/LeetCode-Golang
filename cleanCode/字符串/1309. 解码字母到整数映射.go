package 字符串

import "fmt"

func freqAlphabets(s string) string {
	charMap := getCharMap()
	result := make([]byte, 0)
	for i := len(s) - 1; i >= 0; {
		if s[i] == '#' {
			result = append(result, charMap[s[i-2:i+1]])
			i -= 3
		} else {
			result = append(result, charMap[s[i:i+1]])
			i -= 1
		}
	}
	return string(reverse(result))
}

func getCharMap() map[string]byte {
	charMap := make(map[string]byte)
	for i := 1; i <= 9; i++ {
		charMap[fmt.Sprintf("%d", i)] = byte(i - 1 + 'a')
	}
	for i := 10; i <= 26; i++ {
		charMap[fmt.Sprintf("%d#", i)] = byte(i - 1 + 'a')
	}
	return charMap
}

func reverse(array []byte) []byte {
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
	}
	return array
}

/*
	题目链接: https://leetcode-cn.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
	总结:
		1. 这题反向遍历更好理解。
*/
