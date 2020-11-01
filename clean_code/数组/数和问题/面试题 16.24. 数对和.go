package 数和问题

import "sort"

// ---------------- 方法1: 利用排序+两数和思想 ----------------
func pairSums(nums []int, target int) [][]int {
	pairs := make([][]int, 0)
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			pairs = append(pairs, []int{nums[left], nums[right]})
			left++
			right--
		} else {
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return pairs
}

// ---------------- 方法2: 哈希 ----------------
func pairSums(nums []int, target int) [][]int {
	countOfNum := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		countOfNum[nums[i]]++
	}
	pairs := make([][]int, 0)
	for num := range countOfNum {
		for  {
			countOfNum[target-num]--
			countOfNum[num]--
			if countOfNum[num] < 0 || countOfNum[target-num] < 0 {
				break
			}
			pairs = append(pairs, []int{num, target - num})
		}
	}
	return pairs
}
