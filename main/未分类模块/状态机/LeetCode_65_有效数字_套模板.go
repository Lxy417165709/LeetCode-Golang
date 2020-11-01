package main

func hashOpt(opt uint8) int {

	/* 1. 将操作映射为数字，使用数字表示操作。 */
	switch {
	case opt >= '0' && opt <= '9':
		return 0
	case opt == 'e':
		return 1
	case opt == '+' || opt == '-':
		return 2
	case opt == '.':
		return 3
	case opt == ' ':
		return 4
	default:
		return 5
	}
}

func isNumber(s string) bool {
	/* 2. 通过分析，构建状态机矩阵 */
	matrixOfDFA := [][]int{
		{1, 10, 8, 11, 0, 10},
		{1, 3, 10, 2, 6, 10},
		{7, 3, 10, 10, 6, 10},
		{5, 10, 4, 10, 10, 10},
		{5, 10, 10, 10, 10, 10},
		{5, 10, 10, 10, 6, 10},
		{10, 10, 10, 10, 6, 10},
		{7, 3, 10, 10, 6, 10},
		{9, 10, 10, 11, 10, 10},
		{9, 3, 10, 2, 6, 10},
		{10, 10, 10, 10, 10, 10},
		{7, 10, 10, 10, 10, 10},
	}
	/* 3. 通过分析，构建终止状态向量 */
	finishStates := []bool{false, true, true, false, false, true, true, true, false, true, false, false}

	/* 4. 确定当前状态 */
	nowState := 0

	/* 5. 执行状态转换 */
	for i := 0; i < len(s); i++ {
		nowState = matrixOfDFA[nowState][hashOpt(s[i])]
	}

	/* 6. 返回最终结果 */
	return finishStates[nowState]
}

/*
	总结
	1. 这题主要就是找出状态，以及获得状态转换的规则。
	2. AC的代码可能还有一些地方考虑不周...这题是真的难搞。
	3. 状态机很好用！！！
	4. 上面的状态可能还可以进行压缩。
*/
