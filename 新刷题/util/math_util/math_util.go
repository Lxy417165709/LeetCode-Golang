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
