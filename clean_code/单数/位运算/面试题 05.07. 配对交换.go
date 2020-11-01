package 位运算

// -------------------------- 位运算方法1 --------------------------
const countOfBitOfInt64 = 64

func exchangeBits(num int) int {
	for i := countOfBitOfInt64; i >= 2; i -= 2 {
		num = exchangingNthBitAndPreNthBit(num, i)
	}
	return num
}

func exchangingNthBitAndPreNthBit(num int, n int) int {
	xorValue := getNthBitValue(num, uint(n)) ^ (getNthBitValue(num, uint(n-1)) << 1)
	num = num ^ xorValue ^ (xorValue >> 1)
	return num
}

// 从右到左，分别是 第1位、第2位....第64位。
func getNthBitValue(num int, n uint) int {
	return num & (1 << (n - 1))
}

// -------------------------- 位运算方法2 --------------------------
func exchangeBits(num int) int {
	return (getAllOddPositionBitValue(num) << 1) | ((getAllEvenPositionBitValue(num)) >> 1)
}

func getAllOddPositionBitValue(num int) int {
	// 0x 555555555 == 0b 0101 0101 0101 0101 0101 0101 0101 0101
	return num & 0x55555555
}

func getAllEvenPositionBitValue(num int) int {
	// 0x 555555555 == 0b 1010 1010 1010 1010 1010 1010 1010 1010
	return num & 0xaaaaaaaa
}

/*
	题目链接: https://leetcode-cn.com/problems/exchange-lcci/submissions/
*/
