package main

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	return isOnlyHaving([]int{5, 2, 3}, num)
}

// 判断target中是否只有factorSet中的因子
// 需要保证: target > 0、 factorSet[i] > 0
func isOnlyHaving(factorSet []int, target int) bool {
	for i := 0; i < len(factorSet); i++ {
		for target%factorSet[i] == 0 {
			target /= factorSet[i]
		}
	}
	return target == 1
}

/*
	题目链接:
		https://leetcode-cn.com/problems/ugly-number/submissions/		丑数
*/
