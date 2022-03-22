package 一维数组

// findRepeatNumber 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
// 也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
// 链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
func findRepeatNumber(nums []int) int {
	// 1. for 循环让每个 index 对应的位置合适。
	for i := 0; i <= len(nums)-1; i++ {
		// 1.1 让 index 为 i 的位置合适。
		for !IsPositionFit(nums, i) {
			// 1.1.1 遇到重复数字。 (这里 i 一定不等于 nums[i]，因为位置不合适。)
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}

			// 1.1.2 没遇到重复数字，此时交换位置，让 index 为 nums[i] 的位置合适。
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}

	// 2. 异常。
	panic("不存在重复数字")
}

// IsPositionFit 位置是否合适。
// 当 nums[index] == index，此时 index 对应的位置合适。
func IsPositionFit(nums []int, index int) bool {
	return nums[index] == index
}
