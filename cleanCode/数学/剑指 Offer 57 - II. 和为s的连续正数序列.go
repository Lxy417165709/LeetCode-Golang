package 数学

import "math"


// ------------------------------- 数学法 -------------------------------
// 概述:
// 		假如用区间 [firstValue, lastValue] 表示数列。 ( [2,5]表示数列{2,3,4,5} )
// 		则以下算法的思想就是:
// 			以公差为 1，和为 target 的条件为基础，暴力枚举 firstValue，
// 			通过数学运算，我们就能获得 lastValue。 如果 lastValue 是整数，
// 			则说明该 [firstValue, lastValue] 合法，否则非法。
const eps = 10e-8
const flagOfInvalidLastValue = -1

func findContinuousSequence(target int) [][]int {
	result := make([][]int, 0)
	for firstValue := 1; firstValue <= target/2; firstValue++ {
		lastValue := getLastValue(target, firstValue)
		if lastValue != flagOfInvalidLastValue {
			result = append(result, getSequence(firstValue, lastValue))
		}
	}
	return result
}

func getLastValue(target int, firstValue int) int {
	possibleLastValue := getPossibleLastValue(target, firstValue)
	if !isInt(possibleLastValue) {
		return flagOfInvalidLastValue
	}
	return int(possibleLastValue)
}

func getSequence(firstValue int, lastValue int) []int {
	sequence := make([]int, 0)
	for i := firstValue; i <= lastValue; i++ {
		sequence = append(sequence, i)
	}
	return sequence
}

func getPossibleLastValue(target int, firstValue int) float64 {
	// 这里涉及到等差数列求和公式, 对于本题来说，公差是1。
	// 则有 target == (firstValue+lastValue)*(lastValue-firstVal+1)/2
	// 移项后，可得到 lastValue = sqrt(float64(target*2-firstValue+firstValue*firstValue)+0.25) - 0.5
	return math.Sqrt(float64(target*2-firstValue+firstValue*firstValue)+0.25) - 0.5
}

func isInt(num float64) bool {
	return num-float64(int(num)) <= eps
}



/*
	题目链接: https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
	总结:
		1. 这题官方有人用滑动窗口AC。
		2. 我采用数学法AC了这题。
*/
