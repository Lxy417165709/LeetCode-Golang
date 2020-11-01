package Geek

const INF = 1000000000000

func smallestRangeI(A []int, K int) int {
	return max(0, getMax(A)-getMin(A)-2*K)
}

func getMax(nums []int) int {
	result := -INF
	for i := 0; i < len(nums); i++ {
		result = max(nums[i], result)
	}
	return result
}

func getMin(nums []int) int {
	result := INF
	for i := 0; i < len(nums); i++ {
		result = min(nums[i], result)
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
