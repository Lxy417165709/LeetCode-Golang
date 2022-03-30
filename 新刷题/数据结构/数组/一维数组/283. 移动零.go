package 一维数组

// moveZeroes 将 0 移动到数组后面，不改变 非0数 的顺序。
func moveZeroes(nums []int) {
	// 1. 初始化。
	zero, notZero := 0, 0

	// 2. 处理。
	for zero <= len(nums)-1 && notZero <= len(nums)-1 {
		// 2.1 notZero 指向非0位。
		for notZero <= len(nums)-1 && nums[notZero] == 0 {
			notZero++
		}

		// 2.2 zero 指向0位。
		for zero < notZero && nums[zero] != 0 {
			zero++
		}

		// 2.3 交换。
		if zero <= len(nums)-1 && notZero <= len(nums)-1 {
			nums[zero], nums[notZero] = nums[notZero], nums[zero]
			notZero++
			zero++
		}
	}
}
