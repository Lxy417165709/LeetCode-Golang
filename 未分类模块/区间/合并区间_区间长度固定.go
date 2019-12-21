package main

import "sort"

func findPoisonedDuration(timeSeries []int, duration int) int {
	last := 0
	ans := 0
	sort.Ints(timeSeries)
	for i := 0; i < len(timeSeries); i++ {
		ans += duration
		if last >= timeSeries[i] {
			ans += timeSeries[i] - last
		}
		last = timeSeries[i] + duration
	}
	return ans
}
/*
	题目链接:
		https://leetcode-cn.com/problems/teemo-attacking/		提莫攻击
*/

/*
	总结
	1. 这就是一道给你一些区间，让你计算这些区间所占的长度的题。 (覆盖的区间不重复计算) (比如 [1,4] 和 [2,5], 算的是[1,5])
*/