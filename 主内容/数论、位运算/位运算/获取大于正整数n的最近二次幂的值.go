package main

import "fmt"

// 负数返回 0
// 获取最高位1的下一位1对应的数值 -> 负数返回0，正数在[0, 2^62-1]时返回比自己大的最小2次幂
func getHighestNextOne(a int) int {
	// 适用于64位及以下
	for i := 1; i <= 32; i *= 2 {
		a |= a >> uint8(i)
	}
	return a + 1
}

// 获取最高位1对应的数值 -> 负数返回0，正数在[0, 2^62-1]时返回不大于自己的最大2次幂
func getHighestOne(a int) int {
	for i := 1; i <= 32; i *= 2 {
		a |= a >> uint8(i)
	}
	// 适用于小于64位的正整数
	return (a + 1) >> 1
}
func main() {
	x := 0
	for {
		fmt.Scan(&x) // ctrl+v会使输入不正常...神马情况
		fmt.Println(x, getHighestNextOne(x), getHighestOne(x))
	}
}
