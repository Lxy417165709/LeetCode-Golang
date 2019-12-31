package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	sumCount := make(map[int]int)
	for i := 0; i < len(C); i++ {
		for t := 0; t < len(D); t++ {
			sumCount[C[i]+D[t]]++
		}
	}
	ans := 0
	for i := 0; i < len(A); i++ {
		for t := 0; t < len(B); t++ {
			ans = ans + sumCount[-(A[i] + B[t])]
		}
	}
	return ans

}
/*
	总结
	1. 这道题咋一看没思路，之后想到了许多种不符合时间要求的解法。
			(1) 4个for循环嵌套暴力: O(n^4)
			(2) 采用hash + 3个for循环嵌套暴力: O(n^3)
			(3) 采用二分 + 3个for循环嵌套暴力: O(n^3logn)
	2. 这道题思路是:
			(1) 先把C、D能组合的总和采用哈希表记录。
			(2) 2个for循环嵌套获取A、B能组合的总和，那么如果要A+B+C+D要组成0，
				sumCount[-(A[i]+B[t])]就是当前包括A[i]和B[t] 且 A+B+C+D为0的个数，采用ans累加该次数。
			(3) ans就是答案了。
	3. 解法时间复杂度为O(n^2)

*/
