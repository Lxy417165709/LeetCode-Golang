package 博弈问题

func canWinNim(n int) bool {
	isVisit = make(map[int]bool)
	return canWinNimExec(n, 3, true)
}

var isVisit map[int]bool	// 记忆化搜索

func hash(n int, turn bool) int {
	turnNumber := 0
	if turn == true {
		turnNumber = 1
	}
	return (n << 2) | turnNumber
}
func canWinNimExec(n int, maxPick int, turn bool) bool {
	if n == 0 {
		return !turn
	}
	if maxPick >= n {
		return turn
	}
	hashNumber := hash(n, turn)
	if x, ok := isVisit[hashNumber]; ok {
		return x
	}
	ans := false
	for i := 1; i <= maxPick; i++ {
		ans = canWinNimExec(n-i, maxPick, !turn)
		// 保证双方都是聪明的
		if ans == turn {
			break
		}
	}
	isVisit[hashNumber] = ans
	return ans
}

/*
	总结
	1. 对于这个题目，采用记忆化搜索会超时！！！这里只是采用记忆化搜索来探究博弈问题。
	2. nim游戏有很简单的解法，就是 n%(maxPick+1)==0时，先手必输，否则必赢。
*/
