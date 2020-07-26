package 未归类

import "sort"

const INF = 1000000000

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minDifference := getMinDifferenceFromSortedArray(arr)
	return getPairsWhoseDifferenceIsSpecificFromSortedArray(arr, minDifference)
}

func getMinDifferenceFromSortedArray(arr []int) int {
	minDifference := INF
	for i := 1; i < len(arr); i++ {
		minDifference = min(minDifference, arr[i]-arr[i-1])
	}
	return minDifference
}

func getPairsWhoseDifferenceIsSpecificFromSortedArray(arr []int, specificDifference int) [][]int {
	pairs := make([][]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == specificDifference {
			pairs = append(pairs, []int{arr[i-1], arr[i]})
		}
	}
	return pairs
}

func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := min(arr[1:]...)
	if a > b {
		return b
	}
	return a
}
