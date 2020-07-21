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
	for i := 0; i < len(bytes); i++ {
		bytes[0], bytes[i] = bytes[i], bytes[0]
		formPermutations(bytes[1:], append(nowPermutation, bytes[0]))
		bytes[0], bytes[i] = bytes[i], bytes[0]
	}
}

/*
	题目链接: https://leetcode-cn.com/problems/permutation-ii-lcci/
*/


