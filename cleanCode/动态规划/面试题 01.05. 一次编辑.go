package 动态规划

func oneEditAway(first string, second string) bool {
	return getEditDistance(first, second) <= 1
}

func getEditDistance(first string, second string) int {
	rows, cols := len(first), len(second)
	editDistance := get2DSlice(rows+1, cols+1)
	editDistance[0][0] = 0
	for col := 1; col <= cols; col++ {
		editDistance[0][col] = editDistance[0][col-1] + 1
	}
	for row := 1; row <= rows; row++ {
		editDistance[row][0] = editDistance[row-1][0] + 1
	}
	for row := 1; row <= rows; row++ {
		for col := 1; col <= cols; col++ {
			if first[row-1] == second[col-1] {
				editDistance[row][col] = editDistance[row-1][col-1]
			} else {
				editDistance[row][col] = min(editDistance[row-1][col-1], editDistance[row][col-1], editDistance[row-1][col]) + 1
			}
		}
	}
	return editDistance[rows][cols]
}


func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := min(arr[1:]...)
	if a > b {
		return b
	}
	return a
}




/*
	题目链接: https://leetcode-cn.com/problems/one-away-lcci/comments/
	总结
		1. 这题可以直接先求编辑距离，再通过判断编辑距离是否小于等于1，就能得到最终答案了。
		2. 我做这题的难处在于：如何表示空串。 ---> 解决方案就是 用[0]表示空串，用[1]表示字符串的 [0,0], 用[2]表示
			字符串[0,1]，以此类推。
*/
