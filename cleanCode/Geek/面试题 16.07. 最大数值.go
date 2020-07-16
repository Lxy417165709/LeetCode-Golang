package Geek

import "math"

func maximum(a int, b int) int {
	highestBit := int(uint(a-b) >> 63)
	return b*(highestBit) + a*(highestBit^1)
}

func maximum(a int, b int) int {
	mid := float64(a+b)/2.0
	distanceFromMidToAorB := math.Abs(float64(a)-mid)
	return int(mid+distanceFromMidToAorB)
}

/*
	题目链接:
		1. 这题真Geek。
		2. 我是看官方题解后才知道怎么写的，虽然知道要使用位运算..
*/
