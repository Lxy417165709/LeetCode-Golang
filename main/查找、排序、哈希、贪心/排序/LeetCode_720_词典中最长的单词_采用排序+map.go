package main

import (
	"sort"
	"strings"
)

func longestWord(words []string) string {
	isLiving := make(map[string]bool)
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return strings.Compare(words[i], words[j]) == -1
		}
		return len(words[i]) < len(words[j])
	})
	ans := ""
	for i := 0; i < len(words); i++ {
		if len(words[i]) == 1 || isLiving[words[i][:len(words[i])-1]] {
			isLiving[words[i]] = true
			if len(ans) < len(words[i]) {
				ans = words[i]
			}
		}
	}
	return ans
}

/*
	总结
	1. 这题我先对words进行按长度从小到大、字典序从小到大的排序。
	2. 扫描words，把所有可能的单词(或单词前驱)用哈希表记录，维护一个ans字符串，用于记录最长的可拼凑字符串。
*/
