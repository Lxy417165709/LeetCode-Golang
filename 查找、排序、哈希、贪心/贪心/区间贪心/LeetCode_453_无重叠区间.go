package main

import "sort"

const inf = 100000000000
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	last := -inf
	count := 0
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] >= last {
			last = intervals[i][1]
			count++
		}
	}
	return len(intervals) - count
}
/*
	题目链接:
		https://leetcode-cn.com/problems/non-overlapping-intervals/submissions/		无重叠区间
*/

/*
    总结
    1. 这是一道区间贪心题，和活动安排问题同个性质，就是找区间中最多的不重叠区间，
        做法就是先按区间末端排序，然后遍历一次，每次遇到区间始点大于、或大于等于last时更新last,
        (大于还是大于等于取决于题意，如果[1,2][2,3]认为是覆盖的话，那么就大于，否则大于等于)
        并把不覆盖区间+1。
	2. 这道题在此基础上变形，问需要删除最少多少个区间，可以使区间可以不重叠，那么也很简单，先求出最多的
		不覆盖区间，再把 区间总数 - 最大不覆盖区间 就 求出了最少删除区间。
*/
