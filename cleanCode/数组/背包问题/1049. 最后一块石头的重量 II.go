package 背包问题

// ---------------------- 01背包问题 ----------------------
// 执行用时：0 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.3 MB, 在所有 Go 提交中击败了 100.00% 的用户

func lastStoneWeightII(stones []int) int {
	canForm := getRelationOfCanForm(stones)
	weightOfAllStone := getSum(stones)
	maxOffsetWeight := 2 * getCanFormNumNearestAndLessRef(canForm, weightOfAllStone/2) // Offset: 抵消
	return weightOfAllStone - maxOffsetWeight
}

func getRelationOfCanForm(stones []int) [][]bool {
	sum := getSum(stones)
	canForm := get2DBoolSlice(len(stones)+1, sum+1) // canForm[i][t]表示 在 stones[:i] 中取出一些数，能否组成 t 。
	canForm[0][0] = true
	for i := 1; i <= len(stones); i++ {
		for t := 0; t <= sum; t++ {
			if t-stones[i-1] >= 0 {
				canForm[i][t] = canForm[i-1][t] || canForm[i-1][t-stones[i-1]]
			} else {
				canForm[i][t] = canForm[i-1][t]
			}
		}
	}
	return canForm
}

func getCanFormNumNearestAndLessRef(canForm [][]bool, ref int) int {
	for i := ref; i >= 1; i-- {
		if canForm[len(canForm)-1][i] {
			return i
		}
	}
	panic("不应到达这里")
}

func getSum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func get2DBoolSlice(rows, column int) [][]bool {
	slice := make([][]bool, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]bool, column)
	}
	return slice
}

/*
	题目链接: https://leetcode-cn.com/problems/last-stone-weight-ii/
	总结:
		1. 这题就是要找出，能组成的最接近sum/2的数。
		2. 之前做过类似的题目，就是背包问题嘛。
*/
