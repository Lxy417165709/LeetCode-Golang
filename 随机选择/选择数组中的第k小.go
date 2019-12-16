package main

import "math/rand"

// 随机划分 (l，r在合法范围内)
// 返回的是索引值
func randomPartition(nums []int, l int, r int) int {
	randomIndex := rand.Intn(r-l+1) + l                     // 选择随机索引
	nums[randomIndex], nums[l] = nums[l], nums[randomIndex] // 打乱数组
	guardIndex := l
	for l <= r {
		for l <= r && nums[l] <= nums[guardIndex] {
			l++
		}
		for l <= r && nums[r] >= nums[guardIndex] {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[guardIndex], nums[r] = nums[r], nums[guardIndex]
	return r
}

// 选择第k小的数 (重复的数次序不等)
func selectSmallKth(nums []int, k int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		index := randomPartition(nums, l, r)
		if index+1 == k {
			return nums[index]
		} else {
			if index+1 > k {
				r = index - 1
			} else {
				l = index + 1
			}
		}
	}
	return -100000000 // 表示没找到(k非法了)
}

// 选择第k大的数 (重复的数次序不等)
func selectBigKth(nums []int, k int) int {
	// 第k大 == 第 len(nums)-k+1 小
	return selectSmallKth(nums, len(nums)-k+1)
}

// LeetCode_414_第三大的数
func thirdMax(nums []int) int {

	// 这里是为了给数组去重
	A := make([]int, 0)
	isExist := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := isExist[nums[i]]; !ok {
			A = append(A, nums[i])
			isExist[nums[i]] = 1
		}
	}
	if len(A) < 3 {
		return selectBigKth(A, 1)
	} else {
		return selectBigKth(A, 3)
	}
}

/*
	题目链接:
		https://leetcode-cn.com/problems/third-maximum-number/		第三大的数
*/
/*
	总结
	1. 随机选择算法和快速选择算法都有用到随机划分 randomPartition。
	2. 这题还可以使用最小堆实现。
*/
