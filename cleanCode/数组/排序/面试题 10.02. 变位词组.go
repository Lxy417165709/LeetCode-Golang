package 排序

import "sort"

func groupAnagrams(strs []string) [][]string {
	stringsOfSameAnagram := make(map[string][]string)
	for _,str := range strs{
		specificOrderString := getStringOfObeyingLexicographicOrder(str)
		stringsOfSameAnagram[specificOrderString] = append(stringsOfSameAnagram[specificOrderString], str)
	}
	result := make([][]string, 0)
	for _, strings := range stringsOfSameAnagram {
		result = append(result, strings)
	}
	return result
}

func getStringOfObeyingLexicographicOrder(src string) string {
	bytes := []byte(src)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

/*
	题目链接: https://leetcode-cn.com/problems/group-anagrams-lcci/submissions/
	总结:
		1. 先排序，再哈希~
*/
