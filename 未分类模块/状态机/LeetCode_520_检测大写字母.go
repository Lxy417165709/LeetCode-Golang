package main

// 今天迷上自动机了，这题采用自动机进行解决
func hashOpt(opt uint8) int {
	// 操作
	// 小写字母: 0
	// 大写字母: 1
	switch {
	case opt >= 'a' && opt <= 'z':
		return 0
	case opt >= 'A' && opt <= 'Z':
		return 1
	default:
		return 0
	}
}

func detectCapitalUse(s string) bool {
	// 状态机矩阵
	matrixOfDFA := [][]int{
		{1, 2}, // 空白态     状态0
		{1, 4}, // 小写态     状态1
		{1, 3}, // 单个大写字母态      状态2
		{4, 3}, // 多个大写字母态      状态3
		{4, 4}, // 非法态     状态4
	}

	// 上面有多少个状态，下面就要有多少个表项，是一一对应的关系
	// 比如此时状态0,对应的是true,状态4对应的是false
	// true表示是符合题目要求的状态
	finishStates := []bool{true, true, true, true, false}

	// 当前状态
	nowState := 0

	// 开始通过操作，实现状态转移
	for i := 0; i < len(s); i++ {
		nowState = matrixOfDFA[nowState][hashOpt(s[i])]
	}
	return finishStates[nowState]
}

/*
	总结
	1. 状态机实现步骤:  (以下基于个人总结)
			(1) 确认操作种类
			(2) 确认状态种类
			(3) 找出操作后的状态转移矩阵
			(4) 确定符合要求的状态
			(5) 输入"操作"，开始状态转移	(在上面，"操作"就是字符串中的字符)
			(6) 通过最终的状态，判断是否符合题意要求
*/
