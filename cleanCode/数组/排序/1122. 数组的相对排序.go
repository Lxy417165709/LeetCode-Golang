package 排序

import "sort"

// ------------------------ 自定义排序 ------------------------
// 执行用时：0 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.3 MB, 在所有 Go 提交中击败了 100.00% 的用户

const INF = 1000000000000

func relativeSortArray(arr1 []int, arr2 []int) []int {
	numWeight := getMapOfNumWeight(arr2)
	sort.Slice(arr1, func(i, j int) bool {
		if numWeight[arr1[i]] != numWeight[arr1[j]] {
			return numWeight[arr1[i]] > numWeight[arr1[j]]
		}
		return arr1[i] < arr1[j]
	})
	return arr1
}

func getMapOfNumWeight(arr []int) map[int]int {
	numWeight := make(map[int]int, 0)
	for i := 0; i < len(arr); i++ {
		numWeight[arr[i]] = INF - i // 保证索引越小，weight越大
	}
	return numWeight
}

/*
	题目链接: https://leetcode-cn.com/problems/relative-sort-array/
	总结:
		1. 提交的时候WA了两次，因为没有看清楚题目要求...
*/
