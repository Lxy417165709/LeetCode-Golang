package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 进行排序，让区间开端小的排在前面，如果相等，则让尾端小的排在前面
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	// 执行区间合并
	intervalSequence := make([][]int, 0)
	first, last := intervals[0][0], intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		// 这里控制 [1, 3] [3, 5] 是否合并为[1, 5]
		// >  : 不合并
		// >= : 合并
		if last >= intervals[i][0] {
			last = max(last, intervals[i][1])
		} else {
			intervalSequence = append(intervalSequence, []int{first, last})
			first, last = intervals[i][0], intervals[i][1]
		}
	}
	intervalSequence = append(intervalSequence, []int{first, last})
	return intervalSequence
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接:
		https://leetcode-cn.com/problems/merge-intervals/submissions/		合并区间
*/

/*
	总结
	1. 这题就是给一堆区间给你，让你合并为一些没有重叠部分的区间。
	2. 注意: 这题目对于区间 [1, 3] [3, 5] 是需要合并的，合并为[1, 5]
*/
