package 动态规划



// --------------------------------- 这种方法是超时的，可以参考最下面的总结部分进行算法改进 ------------------------
const INF = 100000000000000000

var minCoinsOfFormingSum int

func coinChange(coins []int, amount int) int {
	minCoinsOfFormingSum = INF
	formMinCoinsOfFormingSum(coins, 0, 0, amount)
	if minCoinsOfFormingSum==INF{
		return -1
	}
	return minCoinsOfFormingSum
}

func formMinCoinsOfFormingSum(array []int, nowCoins int, nowSum int, sumOfWantToSelect int) {
	if nowSum > sumOfWantToSelect {
		return
	}
	if nowSum == sumOfWantToSelect {
		minCoinsOfFormingSum = min(minCoinsOfFormingSum, nowCoins)
		return
	}
	for i := 0; i < len(array); i++ {
		array[0], array[i] = array[i], array[0]
		formMinCoinsOfFormingSum(array, nowCoins+1, nowSum+array[0], sumOfWantToSelect)
		array[0], array[i] = array[i], array[0]
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/coin-change/submissions/
	总结
		1. 这种解法超时了
		2. 这题的内部模型和 _377. 组合总和 Ⅳ_ 是一样。
		3. 可以利用 _377. 组合总和 Ⅳ_ 的思路，修改递归定义，使用记忆化搜索就能AC了。
*/
