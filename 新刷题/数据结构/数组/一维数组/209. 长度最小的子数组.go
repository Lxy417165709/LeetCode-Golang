package 一维数组

import "fmt"

// ------------------------------------------ 1. 滑动窗口(开始) ------------------------------------------

// INF 无穷大。
const INF = 10000000000

// minSubArrayLen 获取最短的子数组，满足数和大于 target。
func minSubArrayLen(target int, nums []int) int {
	// 1. 异常返回。
	if len(nums) == 0 {
		panic("题目出错")
	}

	// 2. 初始化窗口。
	left, right := 0, 0 // 窗口区间为 [left, right]。
	windowSum := 0

	// 3. 滑动。
	minLength := INF
	for right < len(nums) {
		windowSum += nums[right]
		for left <= right && windowSum >= target {
			minLength = min(minLength, right-left+1)
			windowSum -= nums[left]
			left++
		}
		right++
	}

	// 4. 判断是否存在满足条件的窗口。
	if minLength == INF {
		return 0
	}

	// 5. 返回。
	return minLength
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// ------------------------------------------ 1. 滑动窗口(结束) ------------------------------------------

// ------------------------------------------ 2. 前缀和 + 二分(开始) ------------------------------------------

// minSubArrayLen 获取最短的子数组，满足数和大于 target。
func minSubArrayLen(target int, nums []int) int {
	// 1. 获取前缀和。
	prefixSumArray := getPrefixSumArray(nums)

	// 2. 二分。
	minLength := INF
	for i := 0; i < len(nums); i++ {
		// 2.1 查找满足条件的子数组边界。 (返回的边界相对于前缀和数组，前缀和数组索引 - 1 = 原数组索引)
		pos := i + 1
		left := pos
		right := getFirstGreaterOrEqual(
			prefixSumArray, target+prefixSumArray[left-1], left, len(prefixSumArray)-1)

		// 2.2 右边界越界说明不存在满足条件的子数组。
		if right <= len(prefixSumArray)-1 {
			minLength = min(minLength, right-left+1)
		}
	}

	// 3. 判断是否存在满足条件的长度。
	if minLength == INF {
		return 0
	}

	// 4. 返回。
	return minLength
}

// getPrefixSumArray 获取前缀和数组。
func getPrefixSumArray(nums []int) []int {
	prefixSum := make([]int, len(nums)+1)
	for index := range nums {
		pos := index + 1
		prefixSum[pos] = prefixSum[pos-1] + nums[index]
	}
	return prefixSum
}

// getFirstGreaterOrEqual 获取第一个大于或等于 target 的数的索引。
func getFirstGreaterOrEqual(nums []int, target int, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// ------------------------------------------ 2. 前缀和 + 二分(开始) ------------------------------------------
