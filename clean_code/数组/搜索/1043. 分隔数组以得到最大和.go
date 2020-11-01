package 搜索

// -------------------- 暴力搜索(超时) -----------------
const INF = 1000000000000

func maxSumAfterPartitioning(A []int, K int) int {
	return getMaxSumAfterPartitioning(A, K)
}

func getMaxSumAfterPartitioning(A []int, K int) int {
	if K == 0 {
		return 0
	}
	if len(A) == 0 {
		return 0
	}
	maxSum := 0
	for i := 0; i < K && i <= len(A)-1; i++ {
		maxSum = max(
			maxSum,
			getMax(A[:i+1])*(i+1)+getMaxSumAfterPartitioning(A[i+1:], K),
		)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMax(array []int) int {
	result := -INF
	for i := 0; i < len(array); i++ {
		result = max(result, array[i])
	}
	return result
}

// -------------------- 记忆化搜索 -----------------
// 执行用时：108 ms, 在所有 Go 提交中击败了 15.38% 的用户
// 内存消耗：3.9 MB, 在所有 Go 提交中击败了 100.00% 的用户

const INF = 1000000000000

var maxSumMap map[int]int

func maxSumAfterPartitioning(A []int, K int) int {
	maxSumMap = make(map[int]int)
	return getMaxSumAfterPartitioning(A, 0, len(A)-1, K)
}

func getMaxSumAfterPartitioning(A []int, left, right, K int) int {
	hashCode := (left << 10) | K
	if maxSum, ok := maxSumMap[hashCode]; ok {
		return maxSum
	}
	if K == 0 {
		return 0
	}
	if left > right {
		return -INF
	}
	maxSum := 0
	for i := left; i < left+K && i <= right; i++ {
		maxSum = max(
			maxSum,
			getMax(A[:i+1])*(i-left+1)+getMaxSumAfterPartitioning(A, i+1, right, K),
		)
	}
	maxSumMap[hashCode] = maxSum
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMax(array []int) int {
	result := -INF
	for i := 0; i < len(array); i++ {
		result = max(result, array[i])
	}
	return result
}

// -------------------- 记忆化搜索 (优化获取区间最大值的时间复杂度) -----------------
// 执行用时：44 ms,  在所有 Go 提交中击败了 23.08% 的用户
// 内存消耗：7.5 MB, 在所有 Go 提交中击败了 100.00% 的用户

const INF = 1000000000000

var maxSumMap map[int]int
var maxValueInInterval [][]int

func maxSumAfterPartitioning(A []int, K int) int {
	maxSumMap = make(map[int]int)
	maxValueInInterval = getMaxValueInInterval(A)
	return getMaxSumAfterPartitioning(A, 0, len(A)-1, K)
}

func getMaxValueInInterval(A []int) [][]int {
	maxValueInInterval := get2DSlice(len(A), len(A))
	maxValueInInterval[0][0] = A[0]

	for i := 0; i < len(A); i++ {
		maxValueInInterval[i][i] = A[i]
		for t := i + 1; t < len(A); t++ {
			maxValueInInterval[i][t] = max(maxValueInInterval[i][t-1], A[t])
		}
	}
	return maxValueInInterval
}

func getMaxSumAfterPartitioning(A []int, left, right, K int) int {
	hashCode := (left << 10) | K
	if maxSum, ok := maxSumMap[hashCode]; ok {
		return maxSum
	}
	if K == 0 {
		return 0
	}
	maxSum := 0
	for i := left; i < left+K && i <= right; i++ {
		maxSum = max(
			maxSum,
			getMax(left, i)*(i-left+1)+getMaxSumAfterPartitioning(A, i+1, right, K),
		)
	}
	maxSumMap[hashCode] = maxSum
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMax(left, right int) int {
	return maxValueInInterval[left][right]
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

// -------------------- 动态规划 -----------------
// 执行用时：16 ms,  在所有 Go 提交中击败了 34.62%  的用户
// 内存消耗：7.5 MB, 在所有 Go 提交中击败了 100.00% 的用户

const INF = 1000000000000

var maxValueInInterval [][]int

func maxSumAfterPartitioning(A []int, K int) int {
	maxValueInInterval = getMaxValueInInterval(A)
	maxSum := make([]int, len(A))
	for i := 0; i < K; i++ {
		maxSum[i] = getMax(0, i) * (i + 1)
	}
	for i := K; i < len(A); i++ {
		for t := 1; t <= K; t++ {
			maxSum[i] = max(maxSum[i], maxSum[i-t]+getMax(i-t+1, i)*t)
		}
	}
	return maxSum[len(A)-1]
}

func getMaxValueInInterval(A []int) [][]int {
	maxValueInInterval := get2DSlice(len(A), len(A))
	maxValueInInterval[0][0] = A[0]
	for i := 0; i < len(A); i++ {
		maxValueInInterval[i][i] = A[i]
		for t := i + 1; t < len(A); t++ {
			maxValueInInterval[i][t] = max(maxValueInInterval[i][t-1], A[t])
		}
	}
	return maxValueInInterval
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMax(left, right int) int {
	return maxValueInInterval[left][right]
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

/*
	题目链接: https://leetcode-cn.com/problems/partition-array-for-maximum-sum/submissions/
*/
