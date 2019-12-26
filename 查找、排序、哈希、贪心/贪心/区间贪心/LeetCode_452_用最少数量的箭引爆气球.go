package main

import "sort"

const inf = 100000000000

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	last := -inf
	count := 0
	for i := 0; i < len(points); i++ {
		// 无重叠区间是>=,这是>
		if points[i][0] > last {
			last = points[i][1]
			count++
		}
	}
	return count
}

/*
	题目链接:
		https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/		用最少数量的箭引爆气球
*/

/*
    总结
    1. 这题和LeetCode_453_无重叠区间是一样的，只是现在[1,2][2,3]认为是覆盖的，所以上面要修改一个地方(已经在上面指出)。
	2. 这题目是使用最少的弓箭射爆所有气球，所以要返回的是无覆盖区间的个数。
*/
