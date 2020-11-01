package main

var ans int		 // 记录有多少种符合条件的情况
var isUse []bool // 记录数字是否已被使用

func countArrangement(N int) int {
	ans = 0
	isUse = make([]bool, N+1)
	countArrangementExec(0, N)
	return ans
}

// index为摆放数字的位置
// [1, N] 为数字范围，也限定了摆放个数
func countArrangementExec(index int, N int) {
	// 递归终止条件
	if index == N {
		ans ++
		return
	}
	// 遍历每个数字
	for i := 1; i <= N; i++ {
		if isUse[i] {
			continue
		}
		// 满足条件就进入回溯
		if i%(index+1) == 0 || (index+1)%i == 0 {
			isUse[i] = true
			countArrangementExec(index+1, N)
			isUse[i] = false
		}
	}
}
/*
	总结
	1. 这题我采用回溯法。
	2. 官方还有人把答案写下来作为备忘录，实现O(1)时空复杂度
*/