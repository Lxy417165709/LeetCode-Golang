package 字符串

import "bytes"
// ----------------------- 哈希 -----------------------
func customSortString(S string, T string) string {
	countOfChar := getCountOfChar(T)
	result := bytes.Buffer{}
	for i := 0; i < len(S); i++ {
		for t := 0; t < countOfChar[S[i]]; t++ {
			result.WriteByte(S[i])
		}
		delete(countOfChar, S[i])
	}
	for char, count := range countOfChar {
		for t := 0; t < count; t++ {
			result.WriteByte(char)
		}
	}
	return result.String()
}

func getCountOfChar(str string) map[byte]int{
	countOfChar := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		countOfChar[str[i]]++
	}
	return countOfChar
}

/*
	题目链接: https://leetcode-cn.com/problems/custom-sort-string/comments/
	总结：
		1. 这题我采用哈希AC了，但其实还可以用排序AC。
*/
