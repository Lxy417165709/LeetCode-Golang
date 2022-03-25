package math_util

// GetDigitSum 获取数位和。
func GetDigitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

// FastPow 快速幂。
func FastPow(x float64, n int) float64 {
	// 1. 指数小于0时，底数取倒，指数取反。
	if n < 0 {
		return FastPow(1/x, -n)
	}

	// 2. 指数大于0时，使用快速幂计算结果。
	result := 1.0
	for curWeight := x; n != 0; {
		if n&1 == 1 {
			result *= curWeight
		}
		curWeight *= curWeight
		n >>= 1
	}

	// 3. 返回。
	return result
}

// Max 获取数组最大值。
func Max(nums ...int) int {
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	return Max(nums[0], Max(nums[1:]...))
}
