package LAS

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sortEnvelopes(envelopes)
	return getLengthOfLAS(getHeightsOfAllEnvelopes(envelopes))
}

func sortEnvelopes(envelopes [][]int){
	sort.Slice(envelopes, func(i, t int) bool {
		if envelopes[i][0] != envelopes[t][0] {
			return envelopes[i][0] < envelopes[t][0]
		}
		return envelopes[i][1] > envelopes[t][1]
	})
}

func getHeightsOfAllEnvelopes(envelopes [][]int) []int{
	heights := make([]int,0)
	for i := 0; i < len(envelopes); i++ {
		heights = append(heights,envelopes[i][1])
	}
	return heights
}

func getLengthOfLAS(nums []int) int {
	LAS := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		index := getIndexOfFirstGreaterOrEqual(LAS, nums[i])
		if index==len(LAS){
			LAS = append(LAS,nums[i])
		}
		LAS[index]=nums[i]
	}
	return len(LAS)
}

func getIndexOfFirstGreaterOrEqual(nums []int, ref int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
