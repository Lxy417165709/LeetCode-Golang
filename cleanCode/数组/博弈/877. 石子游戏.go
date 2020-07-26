package 博弈


const INF = 1000000000
var prefixSum []int
var hasVisit map[int]int

func stoneGame(stonePiles []int) bool {
	hasVisit = make(map[int]int)
	prefixSum = getPrefixSumArray(stonePiles)
	return getGameResult(stonePiles)
}


func getGameResult(stonePiles []int) bool{
	goalOfAlex := getMaxGoalWhenFirstPick(0, len(stonePiles)-1)
	goalOfLee:= getCountOfStore(0, len(prefixSum)-1) - goalOfAlex
	return goalOfAlex>goalOfLee
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

func getMaxGoalWhenFirstPick(left, right int) int {
	if left > right {
		return 0
	}
	if value, ok := hasVisit[getHash(left,right)]; ok {
		return value
	}

	maxGoal := max(
		getCountOfStore(left, right)-getMaxGoalWhenFirstPick(left+1, right),
		getCountOfStore(left, right)-getMaxGoalWhenFirstPick(left, right-1),
	)

	hasVisit[getHash(left,right)] = maxGoal
	return maxGoal
}

func getHash(left,right int) int{
	const overLeft = 500
	return left+right*overLeft
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


/*
	题目链接: https://leetcode-cn.com/problems/stone-game/
	总结
		1. 公式在手，天下我有！

*/
