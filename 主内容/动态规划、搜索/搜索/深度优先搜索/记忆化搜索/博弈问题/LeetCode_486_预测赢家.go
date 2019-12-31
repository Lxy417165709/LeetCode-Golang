package main
func PredictTheWinner(nums []int) bool {
	isVisit = make(map[int]bool)
	return PredictTheWinnerExec(nums, 0, len(nums)-1, []int{0, 0}, true)
}

var isVisit map[int]bool // 记忆化

func PredictTheWinnerExec(nums []int, l, r int, people []int, turn bool) bool {
	if l > r {
		return people[1] >= people[0]
	}
	hashNumber := hash(l, r, turn, people)
	if x, ok := isVisit[hashNumber]; ok {
		return x
	}

	ans := false
	// 这里for循环的i表示选择，i==0表示取左边，i==1表示取右边，只是取左右两边
	for i := 0; i < 2; i++ {
		turnNumber := boolToInt(turn)
		if i == 0 {
			people[turnNumber] += nums[l]
			ans = PredictTheWinnerExec(nums, l+1, r, people, !turn)
			people[turnNumber] -= nums[l]
		} else {
			people[turnNumber] += nums[r]
			ans = PredictTheWinnerExec(nums, l, r-1, people, !turn)
			people[turnNumber] -= nums[r]
		}
		if turn == ans {
			break
		}
	}
	isVisit[hashNumber] = ans
	return ans
}
// 用于把局势哈希为唯一数字，便于记忆化存储
func hash(l, r int, turn bool, people []int) int {
	return (people[0] << 36) | (people[1] << 11) | (l << 6) | (r << 1) | (boolToInt(turn))
}
func boolToInt(bo bool) int {
	if bo == true {
		return 1
	}
	return 0
}
/*
	总结
	1. 这题和877.石子游戏是差不多的，我把石子游戏的代码copy过来，修改下hash函数就直接AC了。
*/
