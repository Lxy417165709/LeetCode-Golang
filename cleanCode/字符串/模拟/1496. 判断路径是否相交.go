package 模拟

const weightOfX = 1
const weightOfY = 1000000000

func isPathCrossing(path string) bool {
	hasReach := make(map[int]bool)
	positionWeight := 0
	for i := 0; i < len(path); i++ {
		hasReach[positionWeight] = true
		switch {
		case path[i] == 'N':
			positionWeight += weightOfY
		case path[i] == 'S':
			positionWeight -= weightOfY
		case path[i] == 'E':
			positionWeight += weightOfX
		case path[i] == 'W':
			positionWeight -= weightOfX
		}
		if hasReach[positionWeight] {
			return true
		}
	}
	return false
}

/*
	题目链接: https://leetcode-cn.com/problems/path-crossing/
*/
