package main

func countPrimes(n int) int {
	return len(primeGenerator(n - 1)) // 因为题目要求的是 [1,n)的素数个数，所以这里要 -1
}

// 返回[1,n]的质数
// 最容易想到的素数筛 (没有优化)
func primeGenerator(n int) []int {
	isNotPrime := make([]bool, n+1)
	ans := []int{}

	for i := 2; i <= n; i++ {
		if isNotPrime[i] == true {
			continue
		}
		ans = append(ans, i)
		for t := i + i; t <= n; t += i {
			isNotPrime[t] = true
		}
	}
	return ans
}

// Sieve of Eratosthenes
// 返回[1,n]的质数
// 优化版素数筛，高效
func primeGenerator(n int) []int {
	isNotPrime := make([]bool, n+1)
	ans := []int{}

	for i := 2; i*i <= n; i++ {
		if isNotPrime[i] == true {
			continue
		}
		for t := i * i; t <= n; t += i {
			isNotPrime[t] = true
		}
	}
	for i := 2; i <= n; i++ {
		if isNotPrime[i] == false {
			ans = append(ans, i)
		}
	}
	return ans
}

func countPrimes(n int) int {
	return primeCount(n - 1) // 因为题目要求的是 [1,n)的素数个数，所以这里要 -1
}

// 为了提升提交时的时空效率，可以这样写
// 获取 [1,n]的质数个数
func primeCount(n int) int {
	isNotPrime := make([]bool, n+1)
	ans := 0

	for i := 2; i*i <= n; i++ {
		if isNotPrime[i] == true {
			continue
		}
		for t := i * i; t <= n; t += i {
			isNotPrime[t] = true
		}
	}
	for i := 2; i <= n; i++ {
		if isNotPrime[i] == false {
			ans++
		}
	}
	return ans
}


/*
	题目链接:
		https://leetcode-cn.com/problems/count-primes/		计数质数
*/
/*
	总结
	1. 第一版的素数筛时空效率有些差。
	2. 如果为了AC的时空效率更好，其实可以不用记录下所有的质数，只需要个数就可以了。
*/
