package 二分查找

import "sort"

// -------------- 二分查找法 -------------
func triangleNumber(nums []int) int {
	countOfTriangle := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for t := i + 1; t < len(nums); t++ {
			firstEdge, secondEdge := nums[i], nums[t]
			index := getIndexOfLastSmall(nums, firstEdge+secondEdge)
			countOfTriangle += max(index-t, 0) // 这所以要max(0)，是因为原数组存在0
		}
	}
	return countOfTriangle
}

func getIndexOfLastSmall(sortedArr []int, ref int) int {
	return getIndexOfFirstGreaterOrEqual(sortedArr, ref) - 1
}

func getIndexOfFirstGreaterOrEqual(sortedArr []int, ref int) int {
	left, right := 0, len(sortedArr)-1
	for left <= right {
		mid := left + (right-left)/2
		if sortedArr[mid] >= ref {
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

// -------------- 双指针法 -------------
func triangleNumber(nums []int) int {
	countOfTriangle := 0
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		longestEdge := nums[i]
		countOfTriangle += getCountOfPairWhoseSumGreaterThanRef(nums[:i], longestEdge)
	}
	return countOfTriangle
}

func getCountOfPairWhoseSumGreaterThanRef(sortedArr []int, ref int) int {
	left, right := 0, len(sortedArr)-1
	count := 0
	for left < right {
		sum := sortedArr[left] + sortedArr[right]
		if sum > ref {
			count += right - left
			right--
		} else {
			left++
		}
	}
	return count
}

/*
	题目链接： https://leetcode-cn.com/problems/valid-triangle-number/comments/
*/
