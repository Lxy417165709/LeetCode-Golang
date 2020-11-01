package 排序

import "sort"

func numSpecialEquivGroups(A []string) int {
	countOfEquaivalentString := make(map[string]int)
	for i := 0; i < len(A); i++ {
		countOfEquaivalentString[getOrderEquivalentString(A[i])]++
	}
	return len(countOfEquaivalentString)
}

func getOrderEquivalentString(s string) string {
	oddIndexChars, evenIndexChars := getOddIndexCharsAndEvenIndexChars(s)
	sort.Slice(oddIndexChars, func(i, t int) bool {
		return oddIndexChars[i] < oddIndexChars[t]
	})
	sort.Slice(evenIndexChars, func(i, t int) bool {
		return evenIndexChars [i] < evenIndexChars[t]
	})
	return string(oddIndexChars) + string(evenIndexChars)
}

func getOddIndexCharsAndEvenIndexChars(s string) ([]byte, []byte) {
	oddIndexChars, evenIndexChars := make([]byte, 0), make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if i%2 == 1 {
			oddIndexChars = append(oddIndexChars, s[i])
		} else {
			evenIndexChars = append(evenIndexChars, s[i])
		}
	}
	return oddIndexChars, evenIndexChars
}

/*
	题目链接: https://leetcode-cn.com/problems/groups-of-special-equivalent-strings/
	总结:
		1. 这题就是对字符串内部进行排序，从而一个标识字符串，通过这个标识字符串我们就能判断两个字符串是否等价了。
*/
