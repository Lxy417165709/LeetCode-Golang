package 一维数组

import "fmt"

func removeDuplicates(nums []int) int {
	return removeDuplicateElements(nums, 2)
}

// removeDuplicateElements 删除重复元素。
func removeDuplicateElements(nums []int, maxCountOfDuplicateNum int) int {
	// 1. 初始化。
	writeIndex, readIndex := 0, 0

	// 2. 删除重复项。
	for readIndex < len(nums) {
		// 2.1 获取 nums[readIndex] 的索引区间。
		startIndex, endIndex := getIndexInterval(nums, readIndex)

		// 2.2 写入原数组。
		count := min(endIndex-startIndex+1, maxCountOfDuplicateNum)
		for i := 1; i <= count; i++ {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}

		// 2.3 读取下一部分。
		readIndex = endIndex + 1
	}

	// 3. 返回。
	return writeIndex
}

// getIndexInterval 获取 nums[index] 的索引区间。
func getIndexInterval(nums []int, index int) (int, int) {
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
