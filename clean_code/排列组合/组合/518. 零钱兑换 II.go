package 组合

import "sort"

// ----------------- 求取所有组合，之后求其长度 (此法超时) -----------------
var combinations [][]int

func change(amount int, coins []int) int {
	combinations = make([][]int, 0)
	sort.Ints(coins)
	formCombinations(coins, []int{}, 0, amount)
	return len(combinations)
}

func formCombinations(sortedArray []int, nowCombinations []int, nowSum int, sumOfWantToSelect int) {
	if nowSum > sumOfWantToSelect {
		return
	}
	if nowSum == sumOfWantToSelect {
		combinations = append(combinations, NewSlice(nowCombinations))
		return
	}
	for i := 0; i < len(sortedArray); i++ {
		formCombinations(sortedArray[i:], append(nowCombinations, sortedArray[i]), nowSum+sortedArray[i], sumOfWantToSelect)
	}
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

// ----------------- 递归时不记录当前组合，最终只获得所有组合的长度 (此法也超时) -----------------
var countOfCombination int

func change(amount int, coins []int) int {
	countOfCombination = 0
	sort.Ints(coins)
	formCountOfCombination(coins, 0, amount)
	return countOfCombination
}

func formCountOfCombination(sortedArray []int, nowSum int, sumOfWantToSelect int) {
	if nowSum > sumOfWantToSelect {
		return
	}
	if nowSum == sumOfWantToSelect {
		countOfCombination++
		return
	}
	for i := 0; i < len(sortedArray); i++ {
		formCountOfCombination(sortedArray[i:], nowSum+sortedArray[i], sumOfWantToSelect)
	}
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

// ----------------- 修改递归函数定义, 不使用记忆化搜索 (超时) -----------------
func change(amount int, coins []int) int {
	return getCountOfCombination(coins, amount)
}

func getCountOfCombination(array []int, sumOfWantToSelect int) int {
	if sumOfWantToSelect == 0 {
		return 1
	}
	if len(array) == 0 {
		return 0
	}
	countOfCombination := 0
	for timesOfSelectFirstNum := 0; sumOfWantToSelect-timesOfSelectFirstNum*array[0] >= 0; timesOfSelectFirstNum++ {
		countOfCombination += getCountOfCombination(array[1:], sumOfWantToSelect-timesOfSelectFirstNum*array[0])
	}
	return countOfCombination
}

// ----------------- 修改递归函数定义, 使用记忆化搜索 (AC) -----------------
const overLengthOfArray = 10001

var countOfCombinationWithSpecificSumAndArray map[int]int

func change(amount int, coins []int) int {
	countOfCombinationWithSpecificSumAndArray = make(map[int]int)
	return getCountOfCombination(coins, amount)
}

func getCountOfCombination(array []int, sumOfWantToSelect int) int {
	hashCode := sumOfWantToSelect*overLengthOfArray + len(array)
	if value, ok := countOfCombinationWithSpecificSumAndArray[hashCode]; ok {
		return value
	}
	if sumOfWantToSelect == 0 {
		return 1
	}
	if len(array) == 0 {
		return 0
	}
	countOfCombination := 0
	for timesOfSelectFirstNum := 0; sumOfWantToSelect-timesOfSelectFirstNum*array[0] >= 0; timesOfSelectFirstNum++ {
		countOfCombination += getCountOfCombination(array[1:], sumOfWantToSelect-timesOfSelectFirstNum*array[0])
	}
	countOfCombinationWithSpecificSumAndArray[hashCode] = countOfCombination
	return countOfCombination
}

// ----------------- DP (AC) -----------------
func change(amount int, coins []int) int {
	return getCountOfCombination(coins, amount)
}

func getCountOfCombination(arrayOfElementGreaterThanZero []int, sumOfGreaterThanOrEqualZero int) int {
	// 这个dp其实经过了状态压缩
	// 原始二维dp是: dp[i][j]表示用数组的前i个元素可以凑成总和j的组合个数
	countOfCombinationWithSpecificSum := make([]int, sumOfGreaterThanOrEqualZero+1)
	countOfCombinationWithSpecificSum[0] = 1
	for i := 0; i < len(arrayOfElementGreaterThanZero); i++ {
		for t := 1; t <= sumOfGreaterThanOrEqualZero; t++ {
			if t-arrayOfElementGreaterThanZero[i] >= 0 {
				countOfCombinationWithSpecificSum[t] += countOfCombinationWithSpecificSum[t-arrayOfElementGreaterThanZero[i]]
			}
		}
	}
	return countOfCombinationWithSpecificSum[sumOfGreaterThanOrEqualZero]
}

/*
	题目链接: https://leetcode-cn.com/problems/coin-change-2/submissions/
	总结
		1. 这题和 _377. 组合总和 Ⅳ_ 可以形成对比，其是求取排列数，而本题是求取组合数。
		2. 我还没理清 2020年7月8日20:50:20
*/
