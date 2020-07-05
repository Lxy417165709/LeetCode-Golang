package main

func minNumberOfFrogs(croakOfFrogs string) int {
	const croak = "croak"
	var firstCharOfCroak = rune(croak[0])
	var lastCharOfCroak = rune(croak[len(croak)-1])
	var minNumberOfFrogs = 0
	var timesOfChar = make(map[rune]int)

	for _, char := range croakOfFrogs {
		timesOfChar[char]++
		for i := 0; i < len(croak)-1; i++ {
			if timesOfChar[rune(croak[i])] < timesOfChar[rune(croak[i+1])] {
				return -1
			}
		}
		croakingFrogs := timesOfChar[firstCharOfCroak] - timesOfChar[lastCharOfCroak]
		minNumberOfFrogs = max(minNumberOfFrogs, croakingFrogs)
	}
	for i := 0; i < len(croak)-1; i++ {
		if timesOfChar[rune(croak[i])] != timesOfChar[rune(croak[i+1])] {
			return -1
		}
	}
	return minNumberOfFrogs
}

func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := max(arr[1:]...)
	if a < b {
		return b
	}
	return a
}

/*

	题目链接：https://leetcode-cn.com/problems/minimum-number-of-frogs-croaking/
	总结
		1. 最开始的时候使用了模拟的方式，不出意料的超时了。
		2. 看了解答之后，才从模拟之中抽象出来。
		3. 题解有人说这个就类似"俄罗斯方块"
*/
