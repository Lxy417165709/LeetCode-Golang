package 位运算

// ----------- 位运算解法 -----------
func singleNumber(nums []int) int {
	twos, ones := 0, 0
	for _, num := range nums {
		ones = (ones ^ num) & ^twos
		twos = (twos ^ num) & ^ones
	}
	return ones
}

// ----------- bitMap解法 -----------
func singleNumber(nums []int) int {
	countOfBit := make([]int, 64)
	for _, num := range nums {
		for i := 0; i < len(countOfBit); i++ {
			if getBit(num, i) == 1 {
				countOfBit[i]++
			}
		}
	}
	result := 0
	for i := 0; i < len(countOfBit); i++ {
		result |= (countOfBit[i] % 3) << i
	}
	return result
}

func getBit(num, k int) int {
	if num&(1<<k) == 0 {
		return 0
	}
	return 1
}

/*

	题目链接: https://leetcode-cn.com/problems/single-number-ii/
	总结:
		1. 这题就是用 2个bit，表示数组元素bit出现的次数。 (需要%3)
		2. 方法2就是统计所有bit出现的次数。 (由于go语言int是64位的，所以bitmap要64位)
*/
