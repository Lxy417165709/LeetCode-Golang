package 二维数组

import (
	"github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"
)

// ------------------------------------------ 错误解法 ------------------------------------------

// maximalRectangle 下面的动态规划是错误的，因为最大矩形不止以下两种情况。
// [["0","0","1","0"],
//  ["0","0","1","0"],
//  ["0","0","1","0"],
//  ["0","0","1","1"],
//  ["0","1","1","1"],
//  ["0","1","1","1"],
//  ["1","1","1","1"]]
// 此时最大面积在 matrix[4][1] -> matrix[6][3]，而不是 matrix[6][0] -> matrix[6][3] or matrix[3][3] -> matrix[6][3]
func maximalRectangle(matrix [][]byte) int {
	var width [201][201]int
	var height [201][201]int

	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[i]); t++ {
			x := i + 1
			y := t + 1
			if matrix[i][t] == '0' {
				width[x][y] = 0
				height[x][y] = 0
			} else {
				width[x][y] = width[x][y-1] + 1
				height[x][y] = height[x-1][y] + 1
			}

		}
	}

	var rightToLeftMinHeight [201][201]int
	var downToUpMinWidth [201][201]int
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[i]); t++ {
			x := i + 1
			y := t + 1
			if rightToLeftMinHeight[x][y-1] != 0 {
				rightToLeftMinHeight[x][y] = math_util.Min(rightToLeftMinHeight[x][y-1], height[x][y])
			} else {
				rightToLeftMinHeight[x][y] = height[x][y]
			}

			if downToUpMinWidth[x-1][y] != 0 {
				downToUpMinWidth[x][y] = math_util.Min(downToUpMinWidth[x-1][y], width[x][y])
			} else {
				downToUpMinWidth[x][y] = width[x][y]
			}
		}
	}

	//for i := 0; i < len(matrix); i++ {
	//	for t := 0; t < len(matrix[i]); t++ {
	//		x := i + 1
	//		y := t + 1
	//		fmt.Printf("%d ", rightToLeftMinHeight[x][y])
	//	}
	//	fmt.Println()
	//}
	//
	//fmt.Println("------------------")
	//for i := 0; i < len(matrix); i++ {
	//	for t := 0; t < len(matrix[i]); t++ {
	//		x := i + 1
	//		y := t + 1
	//		fmt.Printf("%d ", downToUpMinWidth[x][y])
	//	}
	//	fmt.Println()
	//}
	//fmt.Println("------------------")

	result := 0
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[0]); t++ {
			x := i + 1
			y := t + 1
			result = math_util.Max(result,
				downToUpMinWidth[x][y]*height[x][y],
				width[x][y]*rightToLeftMinHeight[x][y],
			)
		}
	}
	return result
}

// ------------------------------------------ 错误解法 ------------------------------------------

// ------------------------------------------ 单调栈解法 ------------------------------------------

// maximalRectangle 获取矩阵形成的最大面积。
func maximalRectangle(matrix [][]byte) int {
	// 1. 获取矩阵每行的最大柱子高度集。
	var height [201][201]int // height[x][y] 表示以 matrix[x-1][y-1] 为结尾的柱子的最大高度。
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[i]); t++ {
			x := i + 1
			y := t + 1
			if matrix[i][t] == '0' {
				height[x][y] = 0
			} else {
				height[x][y] = height[x-1][y] + 1
			}
		}
	}

	// 2. 获取矩阵形成的最大矩形面积。
	result := 0
	for i := 0; i < len(matrix); i++ {
		x := i + 1
		result = max(result, largestRectangleArea(height[x][1:len(matrix[i])+1]))
	}

	// 3. 返回。
	return result
}

// largestRectangleArea 获取柱子形成的矩形面积。
func largestRectangleArea(heights []int) int {
	indexStack := make([]int, 0)
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	result := 0
	for i := range heights {
		for len(indexStack) != 0 && heights[i] < heights[indexStack[len(indexStack)-1]] {
			height := heights[indexStack[len(indexStack)-1]]
			indexStack = indexStack[:len(indexStack)-1]
			result = max(result, height*(i-indexStack[len(indexStack)-1]-1))
		}
		indexStack = append(indexStack, i)
	}
	return result
}

// ------------------------------------------ 单调栈解法 ------------------------------------------
