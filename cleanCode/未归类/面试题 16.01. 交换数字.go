package 未归类

import "strings"

func swapNumbers(numbers []int) []int {
	numbers[0]=numbers[0]^numbers[1]
	numbers[1]=numbers[0]^numbers[1]
	numbers[0]=numbers[0]^numbers[1]

	return numbers
}

func swapNumbers(numbers []int) []int {
	numbers[0]=numbers[0]-numbers[1]
	numbers[1]=numbers[0]+numbers[1]
	numbers[0]=numbers[1]-numbers[0]
	return numbers
}

/*
	题目链接: https://leetcode-cn.com/problems/swap-numbers-lcci/
*/
