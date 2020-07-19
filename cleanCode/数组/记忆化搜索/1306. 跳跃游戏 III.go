package 记忆化搜索

// ------------------------------ 记忆化搜索 ------------------------------
// 执行用时：40 ms, 在所有 Go 提交中击败了87.93% 的用户
// 内存消耗：6.6 MB, 在所有 Go 提交中击败了100.00% 的用户
var canReachSpecificPosition map[int]bool

func canReach(arr []int, curPosition int) bool {
	canReachSpecificPosition = make(map[int]bool)
	return isReachable(arr, curPosition)
}

func isReachable(arr []int, curPosition int) bool {
	if reachable, ok := canReachSpecificPosition[curPosition]; ok {
		return reachable
	}
	if len(arr) == 0 {
		return false
	}
	if curPosition < 0 || curPosition >= len(arr) {
		return false
	}
	if arr[curPosition] == 0 {
		return true
	}
	canReachSpecificPosition[curPosition] = false // 这句代码能防止回头
	canReachSpecificPosition[curPosition] = isReachable(arr, curPosition+arr[curPosition]) || isReachable(arr, curPosition-arr[curPosition])
	return canReachSpecificPosition[curPosition]
}

/*
	题目链接: https://leetcode-cn.com/problems/jump-game-iii/submissions/
	总结:
		1. 这题要注意回头路的问题。
		2. 官方有采用迭代AC这题的。
*/
