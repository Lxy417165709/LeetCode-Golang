package sort_util

// QuickSort 快速排序。
func QuickSort(nums []int) {
	// 1. 元素小于2的数组，直接返回。
	if len(nums) < 2 {
		return
	}

	// 2. 保证: 左边的数 <= 基准值 <= 右边的数。
	reference := nums[0]
	left, right := 0, len(nums)-1
	for left <= right {
		// 2.1 自左向右，获取第一个大于基准值的索引。
		for left <= right && nums[left] <= reference {
			left++
		}

		// 2.2 自右向左，获取第一个小于基准值的索引。
		for left <= right && nums[right] >= reference {
			right--
		}

		// 2.3 保证两个索引存在。 (如果其中一个不存在，此时必然是 left = right+1)
		if left <= right {
			// 2.3.1 交换。
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	// 3. 摆正基准值位置。
	// 为什么这里一定要是 right 呢? 因为 right 最终指向的数一定小于等于基准值。场景: [1 2 9 6 8]
	nums[0], nums[right] = nums[right], nums[0]

	// 4. 递归。
	QuickSort(nums[:right])
	QuickSort(nums[right+1:])
}
