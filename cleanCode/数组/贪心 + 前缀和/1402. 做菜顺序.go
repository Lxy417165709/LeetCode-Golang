package 贪心___前缀和

import "sort"

var INF = 1000000000

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	prefixSum := getPrefixSum(reverse(satisfaction))
	preprefixSum := getPrefixSum(prefixSum)

	// prefixSum[X] =    satisfaction[N] + satisfaction[N-1] + ... + satisfaction[X]
	// preprefixSum[X] = (N-X+1) * satisfaction[N] + (N-X) * satisfaction[N-1] + ... + satisfaction[X]
	// 其中 X 属于 [0, N], N == len(satisfaction) - 1

	return max(0, getMax(preprefixSum))
}

func reverse(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

func getPrefixSum(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	prefixSum := make([]int, len(arr))
	prefixSum[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}
	return prefixSum
}

func getMax(arr []int) int {
	result := -INF
	for i := 0; i < len(arr); i++ {
		result = max(result, arr[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	总结:
		1. 这题的贪心策略是: 满意度最高的放在数组最后。
*/
