package main

// 滑动窗口
func equalSubstring(s string, t string, maxCost int) int {
	maxLength := 0
	cost := 0
	first, last := 0, 0
	for last < len(s) {
		cost += calculateCost(s[last], t[last])
		for first <= last && cost > maxCost {
			cost -= calculateCost(s[first], t[first])
			first++
		}
		maxLength = max(maxLength, last-first+1)
		last++
	}
	return maxLength
}

func calculateCost(a, b uint8) int {
	if a > b {
		return int(a - b)
	}
	return int(b - a)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 二分优化
func equalSubstring(s string, t string, maxCost int) int {
	// costSum[0] == 0 表示没进行转换时的总花费
	// costSum[i] 表示 将s[:i]转为t[:i]的总花费
	costSum := make([]int, 1+len(s))
	for i := 1; i <= len(s); i++ {
		costSum[i] = costSum[i-1] + calculateCost(s[i-1], t[i-1])
	}

	maxLength := 0
	first, last := 0, 0 // costSum[last]-costSum[first-1] 为s[first:last+1]转为t[first:last+1]的花费
	for last < len(costSum) {
		first = lowwerBound(costSum, costSum[last]-maxCost) + 1
		maxLength = max(maxLength, last-first+1)
		last++
	}
	return maxLength
}

// 在arr中找到第一个大于等于target的数
func lowwerBound(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			r = mid - 1
		} else {
			if arr[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return l
}

/*
	题目链接:
		https://leetcode-cn.com/problems/get-equal-substrings-within-budget/submissions/		尽可能使字符串相等
*/

/*
	总结
	1. 滑动窗口的题目都类似，一般就是先求窗口内的代价，之后判断代价是否满足要求，不满足则缩小，满足则继续前进。
	2. 这题还可以使用二分优化。
	3. 为什么使用 lowwerBound(costSum, costSum[last]-maxCost) + 1 来确定first呢？
			解释如下:
				我们定义了 costSum[last]-costSum[first-1] 为窗口内的字符串转为t[first:last+1]的花费
				则为了满足题目要求: 需要满足   	costSum[last] - costSum[first-1] <= maxCost
				把cost[first-1]化到左边: 则有	costSum[first-1] >= costSum[last] - maxCost
				于是我们使用二分查找，从costSum中 查找第一个大于等于costSum[last] - maxCost的数的索引，记为x
				则可以得到  first = x + 1。
				完毕。
*/
