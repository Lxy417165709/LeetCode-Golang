package main

// ----------------- 超时的动态规划 -------------------
func nthUglyNumber(n int) int {
	countOfGottenUglyNumber := 1
	isUglyNumber := make(map[int]bool)
	isUglyNumber[1] = true
	readingNumber := 1
	for countOfGottenUglyNumber < n {
		readingNumber++
		// 这还能进行一些去冗余优化
		if readingNumber%2 == 0 && isUglyNumber[readingNumber/2] ||
			readingNumber%3 == 0 && isUglyNumber[readingNumber/3] ||
			readingNumber%5 == 0 && isUglyNumber[readingNumber/5] {
			countOfGottenUglyNumber++
			isUglyNumber[readingNumber] = true
		}

	}
	return readingNumber
}


// ----------------- 不超时的动态规划 -----------------
func nthUglyNumber(n int) int {
	return -1
}
func min(arr ...int) int{
	if len(arr)==1{
		return arr[0]
	}
	a := arr[0]
	b := min(arr[1:]...)
	if a>b{
		return b
	}
	return a
}
