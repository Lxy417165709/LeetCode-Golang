package 数和问题

import "math"

// ----------------------- 方法1: 暴力 -----------------------
// 执行: 超时
//
// 时空复杂度：O(n), O(1)
func judgeSquareSum(c int) bool {
	for firstNum := 0; firstNum <= c/2; firstNum++ {
		secondNum := c - firstNum
		if isSquareNum(firstNum) && isSquareNum(secondNum) {
			return true
		}
	}
	return false
}

func isSquareNum(num int) bool {
	minBaseNum := getMinBaseNumWhichSquareGreaterOrEqualRef(num)
	return minBaseNum*minBaseNum == num
}

func getMinBaseNumWhichSquareGreaterOrEqualRef(ref int) int {
	return int(math.Floor(math.Sqrt(float64(ref))))
}

// ----------------------- 方法2: 双指针(存储所有平方数) -----------------------
// 执行用时：4 ms, 在所有 Go 提交中击败了 20.86%  的用户
// 内存消耗：7 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时空复杂度：O(sqrt(n)), O(sqrt(n))
func judgeSquareSum(c int) bool {
	squareNumSequence := getSquareNumSequence(getMinBaseNumWhichSquareGreaterOrEqualRef(c))
	return isHavingTwoIndexsWhichSumOfValueEqualRef(squareNumSequence, c)
}

func getSquareNumSequence(maxBaseNum int) []int {
	squareNumSequence := make([]int, 0)
	for i := 0; i <= maxBaseNum; i++ {
		squareNumSequence = append(squareNumSequence, i*i)
	}
	return squareNumSequence
}

func getMinBaseNumWhichSquareGreaterOrEqualRef(ref int) int {
	return int(math.Floor(math.Sqrt(float64(ref))))
}

func isHavingTwoIndexsWhichSumOfValueEqualRef(array []int, ref int) bool {
	// 允许索引相等
	index1, _ := getTwoIndexsWhichSumOfValueEqualRef(array, ref)
	return index1 != -1
}

func getTwoIndexsWhichSumOfValueEqualRef(array []int, ref int) (int, int) {
	left, right := 0, len(array)-1
	for left <= right {
		curSum := array[left] + array[right]
		if curSum == ref {
			return left, right
		}
		if curSum > ref {
			right--
		} else {
			left++
		}
	}
	return -1, -1 // 不存在时才返回这个
}

// ----------------------- 方法3: 双指针(不存储所有平方数) -----------------------
// 执行用时：0 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时空复杂度：O(sqrt(n)), O(1)
func judgeSquareSum(c int) bool {
	firstNum, secondNum := 0, getMinBaseNumWhichSquareGreaterOrEqualRef(c)
	for firstNum <= secondNum {
		curSquareSum := firstNum*firstNum + secondNum*secondNum
		if curSquareSum == c {
			return true
		}
		if curSquareSum > c {
			secondNum--
		} else {
			firstNum++
		}
	}
	return false
}

func getMinBaseNumWhichSquareGreaterOrEqualRef(ref int) int {
	return int(math.Floor(math.Sqrt(float64(ref))))
}

/*
	题目链接: https://leetcode-cn.com/problems/sum-of-square-numbers/submissions/
*/
