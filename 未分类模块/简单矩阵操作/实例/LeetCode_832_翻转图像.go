package main

func flipAndInvertImage(A [][]int) [][]int {
	horizontalFlip(A)
	invert(A)
	return A
}

func horizontalFlip(A [][]int) {
	if len(A) == 0 {
		return
	}
	m, n := len(A), len(A[0])
	for i := 0; i < m; i++ {
		for t := 0; t < n/2; t++ {
			A[i][t], A[i][n-1-t] = A[i][n-1-t], A[i][t]
		}
	}
}
func invert(A [][]int) {
	if len(A) == 0 {
		return
	}
	m, n := len(A), len(A[0])
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			A[i][t] ^= 1
		}
	}
}

/*
	总结
	1. 这道题invert和horizontalFlip其实可以合在二重循环实现，但是这样函数复用性就降低了，所以我还是
       两个函数分别执行实现题意所需要求。
*/
