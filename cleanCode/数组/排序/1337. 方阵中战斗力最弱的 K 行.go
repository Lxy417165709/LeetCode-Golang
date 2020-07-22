package 排序

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	units := make([]*Unit, 0)
	for i := 0; i < len(mat); i++ {
		units = append(units, &Unit{
			Row:            i,
			CountOfSoldier: getCountOfOne(mat[i]),
		})
	}
	sort.Slice(units, func(i, t int) bool {
		if units[i].CountOfSoldier == units[t].CountOfSoldier {
			return units[i].Row < units[t].Row
		}
		return units[i].CountOfSoldier < units[t].CountOfSoldier
	})
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, units[i].Row)
	}
	return result
}

func getCountOfOne(arr []int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}

type Unit struct {
	Row            int
	CountOfSoldier int
}

/*
	题目链接: https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/
*/
