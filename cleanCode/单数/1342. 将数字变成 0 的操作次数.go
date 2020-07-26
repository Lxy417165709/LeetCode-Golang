package 单数

func numberOfSteps(num int) int {
	times := 0
	for num != 0 {
		if num%2 == 1 {
			num--
		} else {
			num /= 2
		}
		times++
	}
	return times
}

/*
	题目链接: https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero/submissions/
	总结:
		1. 这题太水了。
*/
