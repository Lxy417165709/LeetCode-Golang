package main

func minIncrementForUnique(A []int) int {
	count := make([]int, 80005) // 记录数组元素出现的次数
	for i := 0; i < len(A); i++ {
		count[A[i]]++
	}
	cost := 0
	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			continue
		}
		cost += count[i] - 1
		count[i+1] += count[i] - 1
	}
	return cost
}

/*
	总结
	1. 先统计元素出现次数，之后重复的元素进行+1搬移
*/
