package 一维数组

import "fmt"

// todo: 这个方法虽然直观，但是不够精妙。

func removeDuplicates(nums []int) int {
	writeIndex, readIndex := 0, 0
	maxCountOfDuplicateNum := 2
	for readIndex < len(nums) {
		startIndex, endIndex := expanse(nums, readIndex)
		count := min(endIndex-startIndex+1, maxCountOfDuplicateNum)
		for i := 1; i <= count; i++ {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
		readIndex = endIndex + 1
	}
	return writeIndex
}

func expanse(nums []int, index int) (int, int) {
	target := nums[index]
	for cur := index; cur < len(nums); cur++ {
		if nums[cur] != target {
			return index, cur - 1
		}
	}
	return index, len(nums) - 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
