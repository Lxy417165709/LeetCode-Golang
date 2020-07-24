package 位运算

func hammingWeight(num uint32) int {
	countOfOne := 0
	for num != 0 {
		num = removeLowestOne(num)
		countOfOne++
	}
	return countOfOne
}

func removeLowestOne(num uint32) uint32 {
	num ^= getValueOfLowestBit(num)
	return num
}

// 这样也可以
func removeLowestOne(num uint32) uint32 {
	num &= (num - 1)
	return num
}

func getValueOfLowestBit(num uint32) uint32 {
	return num & (-num)
}

/*
	题目链接: https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
*/
