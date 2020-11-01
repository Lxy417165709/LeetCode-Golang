package main

var permutations []string

func permutation(S string) []string {
	permutations = make([]string, 0)
	formPermutations([]byte(S), []byte{})
	return permutations
}

func formPermutations(bytes []byte, nowPermutation []byte) {
	if len(bytes) == 0 {
		permutations = append(permutations, string(nowPermutation))
		return
	}
	hasCharBeenInHead := make(map[byte]bool)
	for i := 0; i < len(bytes); i++ {
		if hasCharBeenInHead[bytes[i]] {
			continue
		}
		hasCharBeenInHead[bytes[i]] = true
		bytes[0], bytes[i] = bytes[i], bytes[0]
		formPermutations(bytes[1:], append(nowPermutation, bytes[0]))
		bytes[0], bytes[i] = bytes[i], bytes[0]
	}
}

/*
	总结:
		1. 这题和 _面试题 08.08. 有重复字符串的排列组合_ 一样。
*/
