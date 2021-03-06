package main

// 注意: 此题其实是在求排列，而不是求组合。

// ----------------- 求取所有排列，之后求其长度 (此法超时) -----------------
var permutations [][]int

func combinationSum4(nums []int, target int) int {
	permutations = make([][]int, 0)
	formPermutations(nums, []int{}, 0, target)
	return len(permutations)
}

func formPermutations(array []int, nowPermutation []int, nowSum int, sumOfWantToSelect int) {
	if nowSum > sumOfWantToSelect {
		return
	}
	if nowSum == sumOfWantToSelect {
		permutations = append(permutations, NewSlice(nowPermutation))
		return
	}
	for i := 0; i < len(array); i++ {
		array[0], array[i] = array[i], array[0]
		formPermutations(array, append(nowPermutation, array[0]), nowSum+array[0], sumOfWantToSelect)
		array[0], array[i] = array[i], array[0]
	}
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

// ----------------- 递归时不记录当前排列，最终只获得所有排列的长度 (此法也超时) -----------------
var countOfPermutation int

func combinationSum4(nums []int, target int) int {
	countOfPermutation = 0
	formCountOfPermutation(nums, 0, target)
	return countOfPermutation
}

func formCountOfPermutation(array []int, nowSum int, sumOfWantToSelect int) {
	if nowSum > sumOfWantToSelect {
		return
	}
	if nowSum == sumOfWantToSelect {
		countOfPermutation++
		return
	}
	for i := 0; i < len(array); i++ {
		array[0], array[i] = array[i], array[0]
		formCountOfPermutation(array, nowSum+array[0], sumOfWantToSelect)
		array[0], array[i] = array[i], array[0]
	}
}

// ----------------- 修改递归函数定义, 不使用记忆化搜索 (超时) -----------------
func combinationSum4(nums []int, target int) int {
	return getCountOfPermutation(nums, target)
}

func getCountOfPermutation(array []int, sumOfWantToSelect int) int {
	if sumOfWantToSelect < 0 {
		return 0
	}
	if sumOfWantToSelect == 0 {
		return 1
	}
	countOfPermutation := 0
	for i := 0; i < len(array); i++ {
		countOfPermutation += getCountOfPermutation(array, sumOfWantToSelect-array[i])
	}
	return countOfPermutation
}

// ----------------- 修改递归函数定义, 使用记忆化搜索 (AC) -----------------
var countOfPermutationWithSpecificSum map[int]int

func combinationSum4(nums []int, target int) int {
	countOfPermutationWithSpecificSum = make(map[int]int)
	return getCountOfPermutation(nums, target)
}

func getCountOfPermutation(array []int, sumOfWantToSelect int) int {
	if value, ok := countOfPermutationWithSpecificSum[sumOfWantToSelect]; ok {
		return value
	}
	if sumOfWantToSelect < 0 {
		return 0
	}
	if sumOfWantToSelect == 0 {
		return 1
	}
	countOfPermutation := 0
	for i := 0; i < len(array); i++ {
		countOfPermutation += getCountOfPermutation(array, sumOfWantToSelect-array[i])
	}
	countOfPermutationWithSpecificSum[sumOfWantToSelect] = countOfPermutation
	return countOfPermutation
}

// ----------------- DP (AC) -----------------
func combinationSum4(nums []int, target int) int {
	return getCountOfPermutation(nums,target)
}


func getCountOfPermutation(arrayOfElementGreaterThanZero []int, sumOfGreaterThanOrEqualZero int) int {
	// 这个dp其实经过了状态压缩
	// 原始二维dp是: dp[i][j]表示用数组的前i个元素可以凑成总和j的排列个数
	countOfPermutationWithSpecificSum := make([]int, sumOfGreaterThanOrEqualZero+1)
	countOfPermutationWithSpecificSum[0] = 1
	for i := 1; i <= sumOfGreaterThanOrEqualZero; i++ {
		for t := 0; t < len(arrayOfElementGreaterThanZero); t++ {
			if i-arrayOfElementGreaterThanZero[t]>=0{
				countOfPermutationWithSpecificSum[i] += countOfPermutationWithSpecificSum[i-arrayOfElementGreaterThanZero[t]]
			}
		}
	}
	return countOfPermutationWithSpecificSum[sumOfGreaterThanOrEqualZero]
}


/*
	题目链接: https://leetcode-cn.com/problems/combination-sum-iv/submissions/
	总结
		1. 此题其实是在求排列，而不是求组合。
*/
