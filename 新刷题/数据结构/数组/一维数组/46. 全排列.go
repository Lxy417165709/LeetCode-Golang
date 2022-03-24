package 一维数组

func permute(nums []int) [][]int {
	return getPermute(nums)
}

// getPermute 获取全排列。 (nums中不存在重复数字)
func getPermute(nums []int) [][]int {
	// 1. 递归边界。
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	// 2. 递归。
	allPermute := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]
		subPermute := getPermute(nums[1:])
		for _, numsInSubPermute := range subPermute {
			allPermute = append(allPermute, append(numsInSubPermute, nums[0]))
		}
		nums[0], nums[i] = nums[i], nums[0] // 这里记得要回溯，因为递归会影响nums的顺序。
	}

	// 3. 返回。
	return allPermute
}
