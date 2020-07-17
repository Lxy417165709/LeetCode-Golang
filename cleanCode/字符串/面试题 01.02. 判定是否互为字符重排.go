package 字符串

func CheckPermutation(s1 string, s2 string) bool {
	countOfS1Char := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		countOfS1Char[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		countOfS1Char[s2[i]]--
		if countOfS1Char[s2[i]] == 0 {
			delete(countOfS1Char, s2[i])
		}
	}
	return len(countOfS1Char) == 0
}

/*
	题目链接: https://leetcode-cn.com/problems/check-permutation-lcci/submissions/
*/
