package main

func canPermutePalindrome(s string) bool {
	timesOfChar := getTimesOfChar(s)
	return isTimesOfCharLessThanOneOddTimes(timesOfChar)
}

func getTimesOfChar(s string) map[byte]int {
	timesOfChar := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		timesOfChar[s[i]]++
	}
	return timesOfChar
}

func isTimesOfCharLessThanOneOddTimes(timesOfChar map[byte]int) bool {
	oddTimes := 0
	for _, times := range timesOfChar {
		if times%2 == 1 {
			oddTimes++
		}
	}
	return oddTimes <= 1
}

/*
	题目链接: https://leetcode-cn.com/problems/palindrome-permutation-lcci/
*/
