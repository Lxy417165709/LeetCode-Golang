package 一维数组

import (
	"fmt"
	"sort"
)

// merge 区间合并。
func merge(intervals [][]int) [][]int {
	// 1. 空返回。
	if len(intervals) == 0 {
		return nil
	}

	// 2. 排序。
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 3. 贪心。
	result := make([][]int, 0)
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		// 3.1 如果当前区间的右端点包括了下一区间的左端点，此时扩展当前区间。
		if right >= intervals[i][0] {
			right = max(right, intervals[i][1])
			continue
		}

		// 3.2 如果当前区间的右端点不包括下一区间的左端点，区间截断。
		result = append(result, []int{left, right})
		left, right = intervals[i][0], intervals[i][1]
	}
	result = append(result, []int{left, right})

	// 4. 返回。
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
