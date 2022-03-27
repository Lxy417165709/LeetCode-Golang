package 字符串

import (
	"fmt"

	"github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"
)

// cache 缓存。
var cache = make(map[string]int)

// minDistance 两个字符串间的最小编辑距离。
func minDistance(word1 string, word2 string) int {
	// 1. 递归出口。
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}

	// 2. 读缓存。
	key := fmt.Sprintf("%s_%s", word1, word2)
	if val, exist := cache[key]; exist {
		return val
	}

	// 3. 写缓存。
	if word1[0] == word2[0] {
		cache[key] = minDistance(word1[1:], word2[1:]) // 抵消。
	} else {
		cache[key] = math_util.Min(
			minDistance(word1[1:], word2[1:]), // 替换。
			minDistance(word1[1:], word2),     // 删除 word1 的第一个字符。
			minDistance(word1, word2[1:]),     // 删除 word2 的第一个字符。
		) + 1
	}

	// 4. 返回。
	return cache[key]
}