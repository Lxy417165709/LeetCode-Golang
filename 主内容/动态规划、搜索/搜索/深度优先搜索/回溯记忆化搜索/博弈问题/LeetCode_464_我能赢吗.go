package 博弈问题

var isVisit map[int]bool	// 用于记忆化
var isUse int			// 用于判断该数字是否可取用

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal == 0 {
		return true
	}
	isUse = 0
	isVisit = make(map[int]bool)
	return canIWinExec(maxChoosableInteger, desiredTotal, 0, 0)
}

// 这函数的意义是: 第一个玩家是否必赢。
// 前面两个参数是题目给定的
// total为累积和
// turn == 0 表示轮到第一个玩家，1 表示轮到第二个
func canIWinExec(maxChoosableInteger int, desiredTotal int, total, turn int) bool {
	// 达到条件时
	if desiredTotal <= total {
		// 如果轮到第一个玩家，说明第二个玩家早已满足要求，所以返回false (第一个玩家必输)
		if turn == 0 {
			return false
		}
		// 如果轮到第二个玩家，说明第一个玩家早已满足要求，所以返回true (第一个玩家必赢)
		return true
	}

	// 这里只是为了记忆化
	hashNumber := hash(isUse, total, turn)
	if x, ok := isVisit[hashNumber]; ok {
		return x
	}

	// 这个for循环是做出选择，玩家可以取用[1, maxChoosableInteger]个
	ans := false
	for i := maxChoosableInteger; i >= 1; i-- {
		// 这里是判断 i 是否可取用
		if (isUse & toBit(i)) != 0 {
			continue
		}
		isUse ^= toBit(i)	// 标记已用
		ans = canIWinExec(maxChoosableInteger, desiredTotal, total+i, turn^1)	// 该玩家做出选择后，轮到下一个玩家, turn^1为轮转操作。
		isUse ^= toBit(i)	// 标记解除

		// 如果当前轮到第一个玩家，他的某个选择可以让第一个玩家(即自己)赢，那么他就做出这个选择。
		if turn == 0  && ans == true{
			break
		}
		// 如果当前轮到第二个玩家，他的某个选择可以让第一个玩家输，那么他就做出这个选择。
		if turn == 1 && ans == false{
			return false
		}
	}
	isVisit[hashNumber] = ans
	return ans
}

// 哈希操作，将整个局势转换为一个数字 (需要保证不同的局势，哈希出来的数字是不同的)
func hash(isUse, total, turn int) int {
	return (isUse << 30) | (total << 10) | (turn)
}

// 位运算
func toBit(a int) int {
	return 1 << a
}