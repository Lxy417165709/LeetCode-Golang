package 贪心

import "strconv"

func maximum69Number(num int) int {
	// 获得最高位的6，将其改为9，如果没有，就不用更改
	numString := strconv.Itoa(num)
	numBytes := []byte(numString)
	for i := 0; i <= len(numBytes)-1; i++ {
		if numBytes[i] == '6' {
			numBytes[i] = '9'
			break
		}
	}
	max69Number, _ := strconv.Atoi(string(numBytes))
	return max69Number
}
/*
	题目链接: https://leetcode-cn.com/problems/maximum-69-number/
*/
