package 数组

func minCount(coinPiles []int) int {
	minTimesOfTakingAwayAllCoins := 0
	for _, coins := range coinPiles {
		minTimesOfTakingAwayAllCoins += getMinTimesOfTakingAway(coins)
	}
	return minTimesOfTakingAwayAllCoins
}

func getMinTimesOfTakingAway(coins int) int {
	return (coins + 1) / 2
}


/*
	题目链接：https://leetcode-cn.com/problems/na-ying-bi/
*/
