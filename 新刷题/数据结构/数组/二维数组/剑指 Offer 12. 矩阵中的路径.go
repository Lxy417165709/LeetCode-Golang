package 二维数组

import (
	"fmt"

	"github.com/Lxy417165709/LeetCode-Golang/新刷题/matrix_util"
	"github.com/Lxy417165709/LeetCode-Golang/新刷题/model_util"
)

func exist(matrix [][]byte, word string) bool {
	height, width := matrix_util.GetHeightAndWidthOfByteMatrix(matrix)
	for i := 0; i < height; i++ {
		for t := 0; t < width; t++ {
			if IsWordExist(matrix, make(map[string]bool), word, i, t) {
				return true
			}
		}
	}
	return false
}

// IsWordExist 判断单词是否存在。
func IsWordExist(matrix [][]byte, hadArrived map[string]bool, word string, row, column int) bool {
	// 1. 空判断。
	if len(word) == 0 {
		return true
	}

	// 2. 越界判断。
	height, width := matrix_util.GetHeightAndWidthOfByteMatrix(matrix)
	if row >= height || row <= -1 || column >= width || column <= -1 {
		return false
	}

	// 3. 到达判断。
	coordinateHash := fmt.Sprintf("%d_%d", row, column)
	if hadArrived[coordinateHash] {
		return false
	}

	// 4. 匹配判断。
	if word[0] != matrix[row][column] {
		return false
	}

	// 5. DFS。
	offsets := []*model_util.Offset{
		{
			X: 1,
			Y: 0,
		},
		{
			X: 0,
			Y: 1,
		},
		{
			X: -1,
			Y: 0,
		},
		{
			X: 0,
			Y: -1,
		},
	}
	for _, offset := range offsets {
		hadArrived[coordinateHash] = true
		if IsWordExist(matrix, hadArrived, word[1:], row+offset.Y, column+offset.X) {
			return true
		}
		hadArrived[coordinateHash] = false // 这里要回滚，否则会对下个方向的DFS造成影响。
	}

	// 6. 返回。
	return false
}
