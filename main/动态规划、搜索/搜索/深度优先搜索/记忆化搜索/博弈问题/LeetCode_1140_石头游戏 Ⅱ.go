package 博弈问题

var sum []int

func stoneGameII(piles []int) int {
	sum = make([]int, len(piles)+1)
	isVisit = make(map[int]int)
	for i := 1; i <= len(piles); i++ {
		sum[i] = sum[i-1] + piles[i-1]
	}
	return stoneGameIIExec(piles, 0, 1)
}

// 获取 piles[l:r+1]区间的总和
func getSum(l, r int) int {
	return sum[r+1] - sum[l]
}

var isVisit map[int]int	// 备忘录

// 用数字表示局势，便于记忆化存储
func hash(index, M int) int {
	return (index << 10) | M
}

// 函数定义: 在piles[index:]中、M的限制下，先手能拿得的最大石头数
func stoneGameIIExec(piles []int, index int, M int) int {
	if index == len(piles) {
		return 0
	}
	hashNumber := hash(index, M)
	if x, ok := isVisit[hashNumber]; ok {
		return x
	}

	nowMaxSelect := 0
	nextMinSelect := 10000000000 // 无穷大
	// 先手拥有主动权，所以在所有选择中，先手执行最优选择就可以了 (即作出可以获得最多石头的选择)
	for i := 1; i <= 2*M && i+index <= len(piles); i++ {
		nextSelect := stoneGameIIExec(piles, index+i, max(M, i))
		nextMinSelect = min(nextMinSelect, nextSelect)
		nowMaxSelect = max(nowMaxSelect, getSum(index, len(piles)-1)-nextMinSelect) // 后手取最小时，先手就是最大了
	}
	isVisit[hashNumber] = nowMaxSelect
	return nowMaxSelect
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
