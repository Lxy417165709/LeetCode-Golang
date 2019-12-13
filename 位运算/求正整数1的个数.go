package main

import "fmt"

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num = num & (num - 1)	// num = num & (num - 1) 会消除num最后一位的1
	}
	return count
}


func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}

/*
	题目链接: https://leetcode-cn.com/problems/number-of-1-bits/submissions/
*/
