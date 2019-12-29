package main

// 求一组数的最小公倍数
func LCM(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	x := LCM(nums[1:]...)
	gcd := GCD(x, nums[0])
	return x / gcd * nums[0]
}

// 求一组数的最大公约数
func GCD(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	x, y := nums[0], GCD(nums[1:]...)
	if y == 0 {
		return x
	}
	if x == 0 {
		return y
	}
	return GCD(y, x%y)
}
