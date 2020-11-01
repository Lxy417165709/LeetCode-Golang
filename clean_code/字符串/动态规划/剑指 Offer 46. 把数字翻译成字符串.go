package 动态规划

import "strconv"

func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	ways := make([]int, len(numStr)+1)
	ways[0], ways[1] = 1, 1
	for i := 2; i <= len(numStr); i++ {
		ways[i] = ways[i-1]
		if numStr[i-2] > '0' && numStr[i-2] < '2' || numStr[i-2] == '2' && numStr[i-1] <= '5' {
			ways[i] += ways[i-2]
		}
	}
	return ways[len(numStr)]
}

/*
	题目链接: https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
	总结:
		1. 这题数据量很小。 (注意测试数据的时候，数据不要超过Int32)
*/
