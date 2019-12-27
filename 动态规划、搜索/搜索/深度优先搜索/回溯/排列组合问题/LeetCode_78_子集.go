package 排列组合问题

// 回溯 + 外部变量 解决求数组子集问题 (数组中没有重复元素，不可重复选取)
var subsetSequence [][]int
func subsets(nums []int) [][]int {
	subsetSequence = make([][]int,0)
	subsetsExec(nums,[]int{})
	return subsetSequence
}

func subsetsExec(nums []int,sequence []int){
	subsetSequence = append(subsetSequence,newSlice(sequence))
	for i:=0;i<len(nums);i++{
		subsetsExec(nums[i+1:],append(sequence,nums[i]))
	}
}

func newSlice(slice []int) []int{
	s := make([]int,len(slice))
	copy(s,slice)
	return s
}

/*
	题目链接:
		https://leetcode-cn.com/problems/subsets/	子集
*/

/*
	总结
	1. 这题求的是数组中的子集，相当于求数组中的所有组合。
*/