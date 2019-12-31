package 滑动窗口

func maxTurbulenceSize(A []int) int {

	first, last := 0, 0
	maxLength := 0
	for last < len(A) {
		// 与模板相比，主要修改了这里
		if first < len(A) {
			judgeNumber := judge(A, first, last)
			if judgeNumber == 0 {
				first = last
			}
			if judgeNumber == -1 {
				first = last - 1
			}
		}
		maxLength = max(maxLength, last-first+1)
		last++
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func compare(a, b int) int {
	if a == b {
		return 0
	}
	if a > b {
		return 1
	} else {
		return -1
	}

}

// -1: 不满足条件，因为变化情况不符合题意
// 0: 不满足条件，因为末尾两个数相等
// 1: 满足条件
func judge(A []int, first, last int) int {
	length := last - first + 1
	if length <= 1 {
		return 1
	}
	if compare(A[last], A[last-1]) == 0 {
		return 0
	}
	if length >= 3 && compare(A[last], A[last-1]) != compare(A[last-1], A[last-2]) {
		return 1
	}
	return -1
}

/*
	总结
	1. 这题目也是考查滑动窗口，但是由于规则的特殊性，滑动窗口模板需要进行一点修改。
*/