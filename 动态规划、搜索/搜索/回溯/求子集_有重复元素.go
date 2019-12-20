package main

// 回溯 + 外部变量 解决求数组子集问题 (数组中有重复元素，不可重复选取)
var subsetSequence [][]int
func subsetsWithDup(nums []int) [][]int {
	subsetSequence = make([][]int, 0)
	sort.Ints(nums) // 与没有重复元素的相比，多了这一步。  		(增加点①)
	subsetsExec(nums, []int{})
	return subsetSequence
}

func subsetsExec(nums []int, sequence []int) {
	subsetSequence = append(subsetSequence, newSlice(sequence))
	isVisit := make(map[int]bool) // 与没有重复元素的相比，多了这一步。  	(增加点②)
	for i := 0; i < len(nums); i++ {
		if isVisit[nums[i]] {
			continue
		}
		isVisit[nums[i]] = true
		subsetsExec(nums[i+1:], append(sequence, nums[i]))
	}
}

// 深拷贝
func newSlice(slice []int) []int {
	s := make([]int, len(slice))
	copy(s, slice)
	return s
}

/*
	题目链接:
		https://leetcode-cn.com/problems/subsets-ii/submissions/	子集2
*/

/*
	总结
	1. 这题求的是数组中的子集，相当于求数组中的所有组合。
	2. 这题要求的数组是有重复元素的，之前求的是没有重复元素的，上面已经标识出了增加了代码的部分。
*/
