package 二分

// 二分法判断是否为完全平方数 (num >= 0)
func isPerfectSquare(num int) bool {
	l, r := 0, 1000000000
	for l <= r {
		mid := (l + r) / 2
		if mid*mid == num {
			return true
		} else {
			if mid*mid > num {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}

/*
	题目链接:
		https://leetcode-cn.com/problems/valid-perfect-square/		有效的完全平方数
*/

/*
	总结
	1. 这道题认为0不是完全平方数，但是提交是可以AC的，因为它限定了只能输入正整数，
       如果要符合题目要求，那么把左边界l初始为1就可以了。
*/
