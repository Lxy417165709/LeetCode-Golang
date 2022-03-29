package 一维数组

// findDuplicates 给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整数，并以数组形式返回。
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findDuplicates(nums []int) []int {
	// 1. 让正数占领 nums[正数-1] 的位置。
	// (比如正数 1、2、3，分别占领 nums[0]、nums[1]、nums[2])
	for i := 0; i < len(nums); i++ {
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// 2. 寻找不在正确位置的正数，加入结果集。
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			result = append(result, nums[i])
		}
	}

	// 3. 返回。
	return result
}
