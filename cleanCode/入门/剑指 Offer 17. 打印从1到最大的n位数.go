package 入门

func printNumbers(n int) []int {
	countOfNum := fastPow(10, n) - 1
	sequence := make([]int, countOfNum)
	for i := 0; i < countOfNum; i++ {
		sequence[i] = i + 1
	}
	return sequence
}

func fastPow(x int, notNegativeExponent int) int {
	result, curWeight := 1, x
	for notNegativeExponent != 0 {
		if (notNegativeExponent & 1) == 1 {
			result *= curWeight
		}
		curWeight *= curWeight
		notNegativeExponent >>= 1
	}
	return result
}

/*
	总结:
		1. 水水水！
*/
