package main

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {

	n := len(status)

	// boxStatus[i]表示盒子i的状态
	// 00: 表示该没有该盒子，也没有钥匙
	// 01: 表示没有盒子，但是有钥匙
	// 10: 表示有盒子，但是没有钥匙
	// 11: 表示有盒子，有钥匙
	boxStatus := make([]int, n)

	// 获取盒子的钥匙状态
	for i := 0; i < len(status); i++ {
		if status[i] == 1 {
			boxStatus[i] |= 1
		}
	}
	stack := []int{} // 栈结构，能入栈的表示是可以打开的箱子
	// 获取盒子的拥有状态
	for i := 0; i < len(initialBoxes); i++ {
		targetBox := initialBoxes[i]
		boxStatus[targetBox] |= 2
		if boxStatus[targetBox] == 3 {
			stack = append(stack, targetBox)
		}
	}

	score := 0
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		score += candies[top]
		for i := 0; i < len(keys[top]); i++ {
			targetBox := keys[top][i]
			if boxStatus[targetBox] == 3 {
				continue
			}
			boxStatus[targetBox] |= 1
			// boxStatus[targetBox]==3时，该盒子已经入队了，所以不能再入队
			if boxStatus[targetBox] == 3 {
				stack = append(stack, targetBox)
			}
		}
		for i := 0; i < len(containedBoxes[top]); i++ {
			targetBox := containedBoxes[top][i]
			// boxStatus[targetBox]==3时，该盒子已经入队了，所以不能再入队
			if boxStatus[targetBox] == 3 {
				continue
			}
			boxStatus[targetBox] |= 2
			if boxStatus[targetBox] == 3 {
				stack = append(stack, targetBox)
			}
		}
	}
	return score
}

/*
	总结
	1. 这是借鉴了解答区大佬的想法后，模拟着写出来的。
	2. 这道题只要理清思路，AC是很简单的...AC不出来可能是由于参数太多，看着都晕了。
*/
