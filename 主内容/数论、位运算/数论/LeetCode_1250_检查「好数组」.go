package main

func isGoodArray(nums []int) bool {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = GCD(ans, nums[i])
	}
	return ans == 1
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

/*
	题目链接:
		https://leetcode-cn.com/problems/check-if-it-is-a-good-array/comments/			检查「好数组」
*/
/*
	总结
	1. 这题用到了裴蜀定理
		裴蜀定理:
			若a,b是整数,且gcd(a,b)=d，那么对于任意的整数x,y,ax+by都一定是d的倍数，
			特别地，一定存在整数x,y，使ax+by=d成立。
		    a,b互质的充要条件是存在整数x,y使ax+by=1。
			这个定理可以拓展为n个数

*/
