package main

func prefixesDivBy5(A []int) []bool {
	ans := make([]bool, len(A))
	sum := 0
	for i := 0; i < len(A); i++ {
		sum = (sum << 1) | A[i]
		ans[i] = (sum%5 == 0)
		sum = sum % 5
	}
	return ans
}
/*
	总结
	1. 这里用到了取模运算的性质   (a * b + c) % m == ((a%m) * b + c) % m
*/
