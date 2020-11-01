package main

func singleNumber(nums []int) int {
	count := make(map[int]int) // map就是一个哈希表

	// 统计数组中每个元素出现的次数
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	// 找出只出现一次的元素并返回
	for i := 0; i < len(nums); i++ {
		if count[nums[i]] == 1 {
			return nums[i]
		}
	}
	// 找不到时，
	return -1
}
