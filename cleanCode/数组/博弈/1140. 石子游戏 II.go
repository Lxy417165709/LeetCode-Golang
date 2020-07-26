package 博弈

const INF = 1000000000

var prefixSum []int
var hasVisit map[int]int

func stoneGameII(piles []int) int {
	hasVisit = make(map[int]int)
	prefixSum = getPrefixSumArray(piles)
	return getMaxGoalWhenFirstPick(0, len(piles)-1, 1)
}

func getPrefixSumArray(baseArray []int) []int {
	if len(baseArray) == 0 {
		return []int{}
	}
	sum := make([]int, len(baseArray))
	sum[0] = baseArray[0]
	for i := 1; i < len(baseArray); i++ {
		sum[i] = sum[i-1] + baseArray[i]
	}
	return sum
}

func getMaxGoalWhenFirstPick(left, right int, M int) int {
	if left > right {
		return 0
	}
	if value, ok := hasVisit[getHash(left, right, M)]; ok {
		return value
	}
	maxGoal := -INF
	for i := 1; i <= 2*M; i++ {
		maxGoal = max(maxGoal, getCountOfStore(left, right)-getMaxGoalWhenFirstPick(left+i, right, max(M, i)))
	}
	hasVisit[getHash(left, right, M)] = maxGoal
	return maxGoal
}

func getHash(left, right, M int) int {
	const overM = 201
	const overLeft = 100
	return M + left*overM + right*overLeft*overM
}

func getCountOfStore(left, right int) int {
	if right >= len(prefixSum) {
		right = len(prefixSum) - 1
	}
	if left == 0 {
		return prefixSum[right]
	}
	return prefixSum[right] - prefixSum[left-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
