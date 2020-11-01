package 摩尔投票

import "sort"

// -------------------- 方法1: 排序 --------------------
// 执行用时：20 ms,  在所有 Go 提交中击败了 93.99%  的用户
// 内存消耗：5.9 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时空复杂度: O(n * log_n),O(1)
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// -------------------- 方法2: 哈希 --------------------
// 执行用时：24 ms,  在所有 Go 提交中击败了 65.33%  的用户
// 内存消耗：5.9 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时空复杂度: O(n),O(n)
func majorityElement(nums []int) int {
	appearTimes := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		appearTimes[nums[i]]++
	}
	for value, times := range appearTimes {
		if times > len(nums)/2 {
			return value
		}
	}
	panic("题目说了一定存在，所以这里一定不会到达")
}

// -------------------- 方法3: 摩尔投票法 --------------------
// 执行用时：16 ms,  在所有 Go 提交中击败了 99.20%  的用户
// 内存消耗：5.9 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时空复杂度: O(n),O(1)
func majorityElement(nums []int) int {
	countOfTicket, candidateNum := 0, 0
	for i := 0; i < len(nums); i++ {
		if countOfTicket == 0 {
			candidateNum = nums[i]
		}
		if nums[i] == candidateNum {
			countOfTicket++
		} else {
			countOfTicket--
		}
	}
	return candidateNum
}

/*
	题目链接: https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
*/
