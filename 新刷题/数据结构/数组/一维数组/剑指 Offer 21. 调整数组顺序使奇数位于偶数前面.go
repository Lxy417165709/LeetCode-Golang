package 一维数组

func exchange(nums []int) []int {
	// 1. 调整位置。
	for left, right := 0, len(nums)-1; left < right; {
		// 1.1 尝试从左到右找第一个偶数。
		for left < right {
			if nums[left]%2 == 0 {
				break
			}
			left++
		}

		// 1.2 尝试从右到左找第一个奇数。
		for left < right {
			if nums[right]%2 == 1 {
				break
			}
			right--
		}

		// 1.3 交换。
		nums[left], nums[right] = nums[right], nums[left]
	}

	// 2. 返回。
	return nums
}
