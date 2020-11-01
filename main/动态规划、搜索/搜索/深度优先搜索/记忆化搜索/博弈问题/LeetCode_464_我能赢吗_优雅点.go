package 博弈问题

var isVisit map[int]bool // 用于记忆化

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	isVisit = make(map[int]bool)
	return canIWinExec(maxChoosableInteger, desiredTotal, 0, true, 0)
}

// 这函数的意义是: 第一个玩家是否必赢。
// 前面两个参数是题目给定的
// total为累积和
// turn == true 表示轮到第一个玩家，false 表示轮到第二个
func canIWinExec(maxChoosableInteger int, desiredTotal int, total int, turn bool, isUse int) bool {
	if desiredTotal == 0 {
		return turn
	}
	// 达到条件时
	if desiredTotal <= total {
		return !turn
	}
	// 这里只是为了记忆化
	hashNumber := hash(isUse, total, turn)
	if x, ok := isVisit[hashNumber]; ok {
		return x
	}
	ans := false
	for i := maxChoosableInteger; i >= 1; i-- {
		if (isUse & toBit(i)) != 0 {
			continue
		}
		ans = canIWinExec(maxChoosableInteger, desiredTotal, total+i, !turn, isUse^toBit(i))
		if turn == ans {
			break
		}
	}
	isVisit[hashNumber] = ans
	return ans
}

// 哈希操作，将整个局势转换为一个数字 (需要保证不同的局势，哈希出来的数字是不同的)
func hash(isUse, total int, turn bool) int {
	turnNumber := 0
	if turn == true {
		turnNumber = 1
	}
	return (isUse << 30) | (total << 10) | turnNumber
}

// 位运算
func toBit(a int) int {
	return 1 << a
}
