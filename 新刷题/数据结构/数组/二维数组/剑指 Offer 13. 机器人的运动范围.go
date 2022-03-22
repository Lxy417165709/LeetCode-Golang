package 二维数组

import (
	"fmt"

	"github.com/Lxy417165709/LeetCode-Golang/新刷题/util/model_util"

	"github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"
)

func movingCount(m int, n int, k int) int {
	return getMovingCount(m, n, k, make(map[string]bool), 0, 0)
}

// getMovingCount 从 (curX, curY) 出发，机器人能走的步数。
func getMovingCount(m int, n int, k int, hadArrived map[string]bool, curX, curY int) int {
	// 1. 越界判断。
	if curX <= -1 || curX >= n || curY <= -1 || curY >= m {
		return 0
	}

	// 2. 到达判断。
	coordinateHash := fmt.Sprintf("%d_%d", curX, curY)
	if hadArrived[coordinateHash] {
		return 0
	}

	// 3. 数位判断。
	if math_util.GetDigitSum(curX)+math_util.GetDigitSum(curY) > k {
		return 0
	}

	// 4. DFS。
	count := 1
	hadArrived[coordinateHash] = true
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
		count += getMovingCount(m, n, k, hadArrived, curX+offset.X, curY+offset.Y)
	}

	// 5. 返回。
	return count
}
