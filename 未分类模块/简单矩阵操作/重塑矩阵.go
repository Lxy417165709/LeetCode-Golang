package main

/*
	给出一个由二维数组表示的矩阵，以及两个正整数r和c，分别表示想要的重构的矩阵的行数和列数。
	重构后的矩阵需要将原始矩阵的所有元素以相同的行遍历顺序填充。
	如果具有给定参数的reshape操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。
*/

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nums
	}
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	ans := [][]int{}
	for i := 0; i < r; i++ {
		ans = append(ans, make([]int, c))
	}
	for i := 0; i < m*n; i++ {
		ans[i/c][i%c] = nums[i/n][i%n]
	}
	return ans
}
/*
	题目链接:
		https://leetcode-cn.com/problems/reshape-the-matrix/		重塑矩阵
*/
