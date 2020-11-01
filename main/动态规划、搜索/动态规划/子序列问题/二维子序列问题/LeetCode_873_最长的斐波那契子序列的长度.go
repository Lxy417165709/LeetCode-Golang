package 二维子序列问题

// 这个哈希把2维信息转为1维信息
func hash(a int, b int) int {
	// a最大是1000，所以左移10位(因为2^10 > 1000)及以上就可以了，为了保险，我让b左移15位置。
	return (b << 15) + a
}


// dp[A[i]][diff] 表示以A[i]结尾，与前一项相差为diff的最长斐波那契子序列的长度
// 那么如果A[t]是其倒数第二项，通过数学推导可以推出，
// 此时以A[t]结尾的最长斐波那契子序列长度为: dp[A[t]][2*A[t]-A[i]]
// 注意: 下面的代码并不是这样的采用二维的dp，因为题目要求: 1 <= A[0] < A[1] < ... < A[A.length - 1] <= 10^9，
// 这使得，如果要使用二维dp,假设dp是map数组，那么我们需要开10^9个map，这显然会超时超内存。

// 于是，我采用了二维转一维的思想，把2个参数哈希为1个参数，即原来的 dp[A[i]][diff] 现在变为了 dp[hash(diff,A[i])]
// 此时就不用开很多空间了。

// 详情请看下面代码
// AC代码
func lenLongestFibSubseq(A []int) int {
	dp := make(map[int]int)
	ans := 0
	for i := 0; i < len(A); i++ {
		for t := 0; t < i; t++ {
			hash1, hash2 := hash(A[i], A[i]-A[t]), hash(A[t], 2*A[t]-A[i])
			dp[hash1] = max(dp[hash1], 2)
			dp[hash1] = max(dp[hash1], dp[hash2]+1)
			ans = max(ans, dp[hash1])
		}
	}
	if ans <= 2 {
		ans = 0
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// --------------------------------- 上面是AC的代码，下面是超时的代码 ----------------------------

// 超时代码
// 这个哈希和上面的哈希不一样!!!
// 这里我是想让dp只开1e7个map，然后通过哈希的方式把A[i]装入dp的第一维。 (我说的第一维的意思是: dp[第一维][第二维])
func hash(a int) int {
	return a % 999997
}
func lenLongestFibSubseq(A []int) int {
	dp := make([]map[int]int, 1000000) // 如果要装下1e9,那这里要开1e9个数组
	for i := 0; i < 1000000; i++ {
		dp[i] = make(map[int]int)
	}
	ans := 0
	for i := 0; i < len(A); i++ {
		for t := 0; t < i; t++ {
			dp[hash(A[i])][A[i]-A[t]] = max(dp[hash(A[i])][A[i]-A[t]], 2)
			dp[hash(A[i])][A[i]-A[t]] = max(dp[hash(A[i])][A[i]-A[t]], dp[hash(A[t])][2*A[t]-A[i]]+1)
			ans = max(ans, dp[hash(A[i])][A[i]-A[t]])
		}
	}
	if ans <= 2 {
		ans = 0
	}
	return ans
}

/*
	题目链接:
		https://leetcode-cn.com/problems/length-of-longest-fibonacci-subsequence/		最长的斐波那契子序列的长度
*/
/*
	总结
	1. 上面的这个方法使用了动态规划，期间遇到了诸多问题，如空间溢出问题。
*/
