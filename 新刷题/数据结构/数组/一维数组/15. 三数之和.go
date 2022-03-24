package 一维数组

import (
	"fmt"
	"sort"
)

// threeSum 三数和。 (三指针朴素版)
func threeSum(nums []int) [][]int {
	// 1. 排序数组。
	sort.Ints(nums)

	// 2. 三指针获取和为0的子数组。
	result := make([][]int, 0)
	isExist := make(map[string]bool)
	for i := 0; i <= len(nums)-1; i++ {
		// 2.1 确定第一个指针。
		firstNumIndex := i

		// 2.2 利用两数和获取第二、第三个指针。
		secondNumIndex := i + 1
		thirdNumIndex := len(nums) - 1
		for secondNumIndex < thirdNumIndex {
			sum := nums[firstNumIndex] + nums[secondNumIndex] + nums[thirdNumIndex]
			key := fmt.Sprintf("%d_%d_%d", nums[firstNumIndex], nums[secondNumIndex], nums[thirdNumIndex])
			if sum == 0 {
				if !isExist[key] {
					result = append(result, []int{nums[firstNumIndex], nums[secondNumIndex], nums[thirdNumIndex]})
				}
				isExist[key] = true
				secondNumIndex++
			} else if sum > 0 {
				thirdNumIndex--
			} else {
				secondNumIndex++
			}
		}
	}

	// 3. 返回。
	return result
}

// threeSum2 三数和。 (三指针优化版)
func threeSum2(nums []int) [][]int {
	// 1. 排序数组。
	sort.Ints(nums)

	// 2. 获取 numCountLeftToRightMap。
	// numCountLeftToRightMap[x] 表示 nums[x:] 的元素个数映射。
	numCountLeftToRightMap := map[int]map[int]int{}
	for _, num := range nums {
		if numCountLeftToRightMap[0] == nil {
			numCountLeftToRightMap[0] = make(map[int]int)
		}
		numCountLeftToRightMap[0][num]++
	}
	for i := 1; i <= len(nums)-1; i++ {
		if numCountLeftToRightMap[i] == nil {
			numCountLeftToRightMap[i] = make(map[int]int)
		}
		for key, value := range numCountLeftToRightMap[i-1] {
			numCountLeftToRightMap[i][key] = value
		}
		numCountLeftToRightMap[i][nums[i-1]]--
	}

	// 2. 获取三指针和为0的数组。
	result := make([][]int, 0)
	for firstNumIndex := 0; firstNumIndex < len(nums); firstNumIndex = getFirstGreater(nums, nums[firstNumIndex]) {
		for secondNumIndex := firstNumIndex + 1; secondNumIndex < len(nums)-1; secondNumIndex = getFirstGreater(nums, nums[secondNumIndex]) {
			thirdNum := -nums[firstNumIndex] - nums[secondNumIndex]
			if numCountLeftToRightMap[secondNumIndex+1][thirdNum] >= 1 {
				result = append(result, []int{nums[firstNumIndex], nums[secondNumIndex], thirdNum})
			}
		}
	}

	// 3. 返回。
	return result
}

// getFirstGreater 获取第一个大于 target 的数的索引。
func getFirstGreater(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
