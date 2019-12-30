package main

// 需要传入2个操作数
func hashOpt(num1, num2 int) int {
	// 操作
	// 等于  大于  小于
	// 0	 1	   2

	switch {
	case num1 == num2:
		return 0
	case num1 > num2:
		return 1
	case num1 < num2:
		return 2
	default:
		return -1
	}
}

// 这题不是判断真假了，而是返回摆动序列的长度
func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 状态机矩阵
	matrixOfDFA := [][]int{
		{0, 1, 2}, // 开始状态: 状态0
		{1, 1, 2}, // 大于状态: 状态1
		{2, 1, 2}, // 小于状态: 状态2
	}

	// 当前状态
	nowState := 0
	ans := 1

	// 执行状态
	for i := 1; i < len(nums); i++ {
		preState := nowState
		nowState = matrixOfDFA[nowState][hashOpt(nums[i-1], nums[i])]
		if preState != nowState {
			ans++
		}
	}
	return ans
}

/*
	总结
	1. 这个状态机和之前的状态机不太一样，之前的都是判断真假的，而这个是要求取摆动序列的长度。
	2. 官方有人有更简单的解法。
*/
