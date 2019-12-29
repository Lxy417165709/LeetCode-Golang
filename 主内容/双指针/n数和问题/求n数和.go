package 双指针

import (
	"fmt"
	"sort"
)

// nSum 函数调用者
// 返回: nums 中组成 sum 的 n 个数的组合 (不可重复选取、有重复、不一定大于0)
func nSum(nums []int, sum int, n int) [][]int {
	// 一定要先排序
	sort.Ints(nums)
	return nSumExec(nums, sum, n)
}

// nSum 函数执行者(递归执行)
func nSumExec(nums []int, sum int, n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	if n == 1 {
		for i := 0; i < len(nums); i++ {
			// 这里可以优化，可以采用二分查找
			if nums[i] == sum {
				return [][]int{{nums[i]}}
			}
		}
		return [][]int{}
	}
	if n == 2 {
		return twoSum(nums, sum)
	}
	combinations := make([][]int, 0)
	isVisit := make(map[int]bool)

	for i := len(nums) - 1; i >= 0; i-- {
		if isVisit[nums[i]] {
			continue
		}
		isVisit[nums[i]] = true

		// 解决 n-1 数和问题
		nextLayCombinations := nSumExec(nums[:i], sum-nums[i], n-1)

		// 合并
		for t := 0; t < len(nextLayCombinations); t++ {
			element := append(nextLayCombinations[t], nums[i])
			combinations = append(combinations, element)
		}
	}
	return combinations
}

// 这就是两数和
// 要求 nums 数组是升序排序的,这个函数的返回值是能组成 sum 的二元组
func twoSum(nums []int, sum int) [][]int {
	l, r := 0, len(nums)-1
	combinations := make([][]int, 0)
	for l < r {
		tmpSum := nums[l] + nums[r]
		leftAdjustFlag, rightAdjustFlag := false, false
		if sum == tmpSum {
			combinations = append(combinations, []int{nums[l], nums[r]})
			leftAdjustFlag, rightAdjustFlag = true, true
		} else {
			if tmpSum > sum {
				rightAdjustFlag = true
			} else {
				leftAdjustFlag = true
			}
		}
		// 下面的操作是为了保证不出现重复的二元组
		if leftAdjustFlag {
			reference := nums[l]
			for l < r && nums[l] == reference {
				l++
			}
		}
		if rightAdjustFlag {
			reference := nums[r]
			for l < r && nums[r] == reference {
				r--
			}
		}
	}
	return combinations
}


/*
	总结
	1. 这其实也是一种排列组合问题
*/