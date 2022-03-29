package 一维数组

// findDisappearedNumbers 长度为 n 的数组nums，其中 nums[i] 在区间 [1, n] 内。找出所有在 [1, n] 范围内但没有出现在 nums 中的数字。
func findDisappearedNumbers(nums []int) []int {
	// 1. 让正数占领 nums[正数-1] 的位置。 (比如正数 1、2、3，分别占领 nums[0]、nums[1]、nums[2])
	for i := 0; i < len(nums); i++ {
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// 2. 寻找没被正数占领的位置，加入结果集。
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			result = append(result, i+1)
		}
	}

	// 3. 返回。
	return result
}
