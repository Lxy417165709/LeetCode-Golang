package 搜索

// ------------------------------ 没有剪枝的数位DFS ------------------------------
// 输入: 1000000000
// 用时: 执行用时：1320 ms
var canRotateDigits = []int{0, 1, 2, 5, 6, 8, 9}
var goodDigits = []int{2, 5, 6, 9}
var countOfGoodNumber = 0

func rotatedDigits(N int) int {
	countOfGoodNumber = 0
	formCountOfGoodNumber(N, 0)
	return countOfGoodNumber
}

func formCountOfGoodNumber(maxValueBoundary int, curValue int) {
	if curValue > maxValueBoundary {
		return
	}
	if isGoodNum(curValue) {
		countOfGoodNumber++

	}
	for i := 0; i < len(canRotateDigits); i++ {
		nextValue := curValue*10 + canRotateDigits[i]
		if nextValue != 0 {
			formCountOfGoodNumber(maxValueBoundary, nextValue)
		}
	}
}

func isGoodNum(num int) bool {
	for num != 0 {
		curDigit := num % 10
		num /= 10
		if isGoodDigit(curDigit) {
			return true
		}
	}
	return false
}

func isGoodDigit(digit int) bool {
	for i := 0; i < len(goodDigits); i++ {
		if digit == goodDigits[i] {
			return true
		}
	}
	return false
}

// ------------------------------ 剪枝的数位DFS ------------------------------
// 输入: 1000000000
// 用时: 执行用时：848 ms
var canRotateDigits = []int{0, 1, 2, 5, 6, 8, 9}
var goodDigits = []int{2, 5, 6, 9}
var countOfGoodNumber = 0

func rotatedDigits(N int) int {
	countOfGoodNumber = 0
	formCountOfGoodNumber(N, 0, false)
	return countOfGoodNumber
}

// 有 isCurValueGoodNum 标志位，通过这个标志位能减少 isGoodNum 的调用。
func formCountOfGoodNumber(maxValueBoundary int, curValue int, isCurValueGoodNum bool) {
	if curValue > maxValueBoundary {
		return
	}
	if isCurValueGoodNum || isGoodNum(curValue) {
		countOfGoodNumber++
		isCurValueGoodNum = true
	}
	for i := 0; i < len(canRotateDigits); i++ {
		nextValue := curValue*10 + canRotateDigits[i]
		if nextValue != 0 {
			formCountOfGoodNumber(maxValueBoundary, nextValue, isCurValueGoodNum)
		}
	}
}

func isGoodNum(num int) bool {
	for num != 0 {
		curDigit := num % 10
		num /= 10
		if isGoodDigit(curDigit) {
			return true
		}
	}
	return false
}

func isGoodDigit(digit int) bool {
	for i := 0; i < len(goodDigits); i++ {
		if digit == goodDigits[i] {
			return true
		}
	}
	return false
}

/*
	题目链接: https://leetcode-cn.com/problems/rotated-digits/
	总结:
		1. 这题的范围很小，所以暴力可以AC。
		2. 上面的代码采用了dfs (也叫数位dp)，能使处理效率高出数个数量级。
*/
