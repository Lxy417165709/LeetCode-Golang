package main

func stoneGame(piles []int) bool {
	isVisit = make(map[int]bool)
	return stoneGameExec(piles, 0, len(piles)-1, []int{0, 0}, true)
}

var isVisit map[int]bool // 记忆化

func stoneGameExec(piles []int, l, r int, people []int, turn bool) bool {
	if l > r {
		return people[1] > people[0]
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
			people[turnNumber] += piles[l]
			ans = stoneGameExec(piles, l+1, r, people, !turn)
			people[turnNumber] -= piles[l]
		} else {
			people[turnNumber] += piles[r]
			ans = stoneGameExec(piles, l, r-1, people, !turn)
			people[turnNumber] -= piles[r]
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
	return (people[0] << 45) | (people[1] << 25) | (l << 15) | (r << 5) | (boolToInt(turn))
}
func boolToInt(bo bool) int {
	if bo == true {
		return 1
	}
	return 0
}

/*
	总结
	1. 注意: 以上代码提交会超时！
	2. 实际上这道题要采用dp解，这样就不会超时了。
*/
