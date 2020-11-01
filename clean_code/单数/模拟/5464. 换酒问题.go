package 模拟

func numWaterBottles(numBottles int, numExchange int) int {
	maxCountOfDrinking := 0
	countOfBlankBottle := 0
	for numBottles != 0 {
		maxCountOfDrinking += numBottles
		countOfBlankBottle = numBottles + countOfBlankBottle
		numBottles = countOfBlankBottle / numExchange
		countOfBlankBottle = countOfBlankBottle % numExchange
	}
	return maxCountOfDrinking
}

/*
	题目链接: https://leetcode-cn.com/problems/water-bottles/submissions/
*/
