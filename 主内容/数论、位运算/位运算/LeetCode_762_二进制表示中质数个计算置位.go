package main

func countPrimeSetBits(L int, R int) int {
	isPrime := [40]bool{}
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 31}
	for i := 0; i < len(prime); i++ {
		isPrime[prime[i]] = true
	}

	ans := 0
	for i := L; i <= R; i++ {
		if isPrime[countOne(i)] {
			ans++
		}
	}
	return ans
}

func countOne(a int) int {
	count := 0
	for a != 0 {
		a = a & (a - 1)
		count++
	}
	return count
}
/*
	总结
	1. 这题关键就是求数x有多少个1。
*/
