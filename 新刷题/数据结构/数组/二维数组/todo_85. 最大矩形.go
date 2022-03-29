package 二维数组

import (
	"github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"
)

// 下面的动态规划是错误的，因为最大矩形不止以下两种情况。
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
