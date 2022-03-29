package 一维数组

// firstMissingPositive 找到数组中第一个缺失的正数。
func firstMissingPositive(nums []int) int {
	// 1. 让正数占领 nums[正数-1] 的位置。 (比如正数 1、2、3，分别占领 nums[0]、nums[1]、nums[2])
	for i := 0; i < len(nums); i++ {
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// 2. 寻找第一个没被正数占领的位置。
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 3. 全被正数占领。
	return len(nums) + 1
}
