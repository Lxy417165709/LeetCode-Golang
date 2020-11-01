package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	result := make([]int, 0)
	leftBound, rightBound := 0, len(matrix[0])-1
	upBound, downBound := 0, len(matrix)-1
	for leftBound <= rightBound && upBound <= downBound {
		for i := leftBound; i <= rightBound ; i++ {
			result = append(result, matrix[upBound][i])
		}
		upBound++
		for i := upBound; i <= downBound; i++ {
			result = append(result, matrix[i][rightBound])
		}
		rightBound--
		for i := rightBound; i >= leftBound && upBound <= downBound; i-- {
			result = append(result, matrix[downBound][i])
		}
		downBound--
		for i := downBound; i >= upBound && leftBound <= rightBound; i-- {
			result = append(result, matrix[i][leftBound])
		}
		leftBound++
	}
	return result
}
/*
	题目链接: https://leetcode-cn.com/problems/spiral-matrix/
	总结:
		1. 这题要注意只有一行的情况
*/
