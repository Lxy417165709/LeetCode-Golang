package 动态规划

func numRollsToTarget(d int, f int, target int) int {
	arrangements := get2DSlice(d+1, target+1)
	arrangements[0][0] = 1
	for i := 1; i <= d; i++ {
		for s := 1; s <= target; s++ {
			for j := 1; j <= f && s-j >= 0; j++ {
				arrangements[i][s] += arrangements[i-1][s-j]
				arrangements[i][s] %= 1000000007
			}
		}
	}
	return arrangements[d][target]
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}


/*
	题目链接: https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/
	总结:
		1. 这题其实是考查求取组合数。
		2. 我采用dp AC了这题
*/
