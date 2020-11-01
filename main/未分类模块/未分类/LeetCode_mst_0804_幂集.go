package main

var result [][]int

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	solve(nums, 0, []int{})
	return result
}

func solve(nums []int, idx int, tmp []int) {

	result = append(result, NewSlice(tmp))
	for i := idx; i < len(nums); i++ {
		solve(nums, i+1, append(tmp, nums[i]))
	}

}

func NewSlice(old []int) []int {
	newSlice := make([]int, len(old))
	copy(newSlice, old)
	return newSlice
}

/*
	总结
	1. 这题考查全排列 0.0..
*/
