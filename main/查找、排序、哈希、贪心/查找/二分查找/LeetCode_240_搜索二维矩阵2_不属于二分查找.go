package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0{
		return false
	}
	m,n := len(matrix),len(matrix[0])
	currentX := 0
	currentY := n-1
	for currentX<m && currentY>=0{
		if matrix[currentX][currentY]==target{
			return true
		}else{
			if matrix[currentX][currentY]>target{
				currentY--
			}else{
				currentX++
			}
		}
	}
	return false
}
/*
	题目链接:
		https://leetcode-cn.com/problems/search-a-2d-matrix-ii/submissions/		搜索二维矩阵 II
*/
/*
	总结
	1. 这题没有使用到二分查找，而是理由矩阵点的特殊性(比如右上角、左下角)。
	2. 小伙伴可以比对一下和搜索二维矩阵1的差异。
*/
