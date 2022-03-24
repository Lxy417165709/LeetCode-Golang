package 一维数组

import "sort"

func twoSum(nums []int, target int) []int {
	// 1. 记录数在原数组的索引。
	numToIndexesMap := make(map[int][]int)
	for index, num := range nums {
		numToIndexesMap[num] = append(numToIndexesMap[num], index)
	}
	// 2. 排序。
	sort.Ints(nums)

	// 3. 获取结果集。
	left, right := 0, len(nums)-1
	firstResult := make([]int, 0)
	for left < right {
		if nums[left]+nums[right] == target {
			firstResult = []int{nums[left], nums[right]}
			break
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			left++
		}
	}

	// 4. 异常情况。
	if len(firstResult) == 0 {
		panic("题目出错")
	}

	// 5. 获取原数组索引。
	if firstResult[0] == firstResult[1] {
		return numToIndexesMap[firstResult[0]]
	} else {
		return append(numToIndexesMap[firstResult[0]], numToIndexesMap[firstResult[1]]...)
	}
}
