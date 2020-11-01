package main

func rotate(matrix [][]int) {
	n := len(matrix)
	leftBound, rightBound := 0, n-1
	upBound, downBound := 0, n-1
	for rightBound >= leftBound && downBound >= upBound {
		for i := 0; i <= rightBound-leftBound-1; i++ {
			tmp := matrix[upBound][leftBound+i]
			matrix[upBound][leftBound+i] = matrix[downBound-i][leftBound]
			matrix[downBound-i][leftBound] = matrix[downBound][rightBound-i]
			matrix[downBound][rightBound-i] = matrix[upBound+i][rightBound]
			matrix[upBound+i][rightBound] = tmp
		}
		leftBound++
		rightBound--
		upBound++
		downBound--
	}
}

/*
	题目链接: https://leetcode-cn.com/problems/rotate-image/
	总结:
		1. 做题技巧:
			1. 可以一个一个旋转，整一行旋转的有些复杂。
			2. 先在草稿纸画出矩阵，之后标明4个端点坐标，之后将4个端点的替换规则写到代码中。
			3. 判断端点的旋转方向。
			4. 在外层循环定义 i，i从[0,rightBound-leftBound-1],并将旋转方向用 +i,-i表示。
			5. 将+i,-i写入代码。
			6. 外边处理完毕后，调整4个边界。
*/
