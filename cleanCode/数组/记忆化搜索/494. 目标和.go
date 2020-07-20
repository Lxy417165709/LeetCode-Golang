package 记忆化搜索

// ----------------------------- 记忆化搜索 -----------------------------
var ways map[int]int

func findTargetSumWays(nums []int, S int) int {
	ways = make(map[int]int)
	return getWays(nums, 0, len(nums)-1, S)
}

func getWays(nums []int, left, right int, curSum int) int {
	if left > right {
		if curSum == 0 {
			return 1
		}
		return 0
	}
	hashCode := (curSum << 20) + left
	if value, ok := ways[hashCode]; ok {
		return value
	}
	ways[hashCode] = getWays(nums, left+1, right, curSum-nums[left]) + getWays(nums, left+1, right, curSum+nums[left])
	return ways[hashCode]
}

// ----------------------------- 动态规划 -----------------------------

func findTargetSumWays(nums []int, S int) int {
	maxSum := getSum(nums)
	offset := maxSum * 2
	if abs(S) > maxSum {
		return 0
	}
	ways := get2DSlice(len(nums)+1, maxSum*4+1) // 将数组开得足够大，可以防止下面代码 t-nums[i-1]+offset 越界。
	ways[0][0+offset] = 1
	for i := 1; i <= len(nums); i++ {
		for t := -maxSum; t <= maxSum; t++ {
			ways[i][t+offset] = ways[i-1][t-nums[i-1]+offset] + ways[i-1][t+nums[i-1]+offset]
		}
	}
	return ways[len(nums)][S+offset]
}

func getSum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

/*
	题目链接: https://leetcode-cn.com/problems/target-sum/
*/
