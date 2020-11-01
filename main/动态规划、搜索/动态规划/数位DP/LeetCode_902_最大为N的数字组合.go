package 数位DP

import "strconv"

func atMostNGivenDigitSet(D []string, N int) int {
	nStr := strconv.Itoa(N)
	// dp[i]表示: 小于等于 N[i:] 且 长度为 |N| - i 的数的个数      (|N|表示 N 的长度)
	// 比如 N == 234 时:
	//		dp[2]表示: 小于等于 N[2:] 即 4，且长度为 3 - 2 == 1 的数的个数
	//		dp[1]表示: 小于等于 N[1:] 即 34，且长度为 3 - 1 == 2 的数的个数
	// 		dp[0]表示: 小于等于 N[0:] 即 234，且长度为 3 - 0 == 3 的数的个数

	// 当 N[i] > D[t]时，那么说明 N[i+1:] 的数字可以任选 D 中的数字，
	// 而 N[i+1:] 一共有 len(N)-i-1 位，所以dp[i] += pow( len(D), len(N)-i-1 )
	//
	// 当 N[i] == D[t]时，这是 N[i+1:] 的数字不能任意选了，而是等于: 小于等于 N[i+1:] 且 长度为 |N| - i - 1 的数的个数，即 dp[i+1]。
	// 于是 dp[i] += dp[i+1]
	//
	// 当 N[i] < D[t]时，显然此时无论怎么选择，N[i:] 一定是不满足条件的，所以 dp[i] += 0

	// 注意，当我们求出 dp[0]后，此时dp[0]的含义是: 小于等于 234，且长度为 3 的数的个数
	// 这里少了 小于等于 234 且 长度为1、 2 的数的个数，于是想得到最终答案，还需要一些操作。
	// 而这个很简单，如下:
	// 		小于等于 234 且 长度为 1 的数的个数为: pow(len(D), 1) 	(1位长度任选，因为一定小于N)
	// 		小于等于 234 且 长度为 2 的数的个数为: pow(len(D), 2)	(2位长度任选，因为一定小于N)
	// 上面很抽象，下面就是代码了。

	dp := make([]int, len(nStr)+1)
	dp[len(nStr)] = 1 // 记得初始化这个，因为 N 的个位和 D[t] 相等时，答案是有效的。
	// 从上面的分析看出，相等时 dp[i] += dp[i+1], 而此时 i 位于个位，即 i == len(nStr)-1，所以这里要初始化dp[len(nStr)] == 1。
	for i := len(nStr) - 1; i >= 0; i-- {
		nDigit := int(nStr[i] - '0')
		for t := 0; t < len(D); t++ {
			dDigit, _ := strconv.Atoi(D[t])

			// N[i] == D[t] 时
			if nDigit == dDigit {
				dp[i] += dp[i+1]
			} else {
				// N[i] > D[t] 时
				if nDigit > dDigit {
					dp[i] += quickPow(len(D), len(nStr)-i-1)
				}
			}
		}
	}
	ans := dp[0]
	// 添加 1、2....n-1 位的数字
	for i := 1; i < len(nStr); i++ {
		ans += quickPow(len(D), i)
	}
	return ans
}

// 快速幂
func quickPow(a, b int) int {
	sum := 1
	for b != 0 {
		if b&1 == 1 {
			sum *= a
		}
		a = a * a
		b >>= 1
	}
	return sum
}

/*
	总结
	1. 第一次做的时候采用了暴力搜索的方法，然后显然超时了，于是想在暴力搜索的基础上进行改进(因为发现了一些可以简化的)，
		但是实在不知道怎么简化...
	2. 看了官方题解后，才发现这道题应该采用数位DP，思路在上面了。
*/
