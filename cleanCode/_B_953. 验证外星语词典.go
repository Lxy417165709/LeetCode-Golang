package main

func isAlienSorted(words []string, order string) bool {
	refWords := make([]string, 0)
	for i := 0; i < len(words); i++ {
		refWords = append(refWords, words[i])
	}
	letterToPriority := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		letterToPriority[order[i]] = 26 - i
	}
	sort.Slice(words, func(i, t int) bool {
		for j := 0; j < min(len(words[i]), len(words[t])); j++ {
			pi, pt := letterToPriority[words[i][j]], letterToPriority[words[t][j]]
			if pi != pt {
				return pi > pt
			}
		}
		return len(words[i]) < len(words[t])
	})
	for i := 0; i < len(refWords); i++ {
		if refWords[i] != words[i] {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}


/*
	总结
		1. 这题可以直接用O(n)时间复杂度的比较。
		2. 我采用的是先排序，再对比，空间复杂度和时间复杂度都有点大。
*/
