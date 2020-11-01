package Geek

// ----------- 迭代写法 -----------
// 执行用时：0 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 100.00% 的用户
func multiply(A int, B int) int {
	result := 0
	for weightBit := 0; B != 0; weightBit++ {
		if B&1 == 1 {
			result += A << uint(weightBit)
		}
		B >>= 1
	}
	return result
}

// ----------- 递归写法 -----------
// 执行用时：0 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.0 MB, 在所有 Go 提交中击败了 100.00% 的用户
func multiply(A int, B int) int {
	if B == 0 {
		return 0
	}
	if B&1 == 1 {
		return A + multiply(A, B>>1)<<1
	}
	return multiply(A, B>>1) << 1
}

/*
	总结:
		1. 这题就是使用乘法的位运算操作。
*/
