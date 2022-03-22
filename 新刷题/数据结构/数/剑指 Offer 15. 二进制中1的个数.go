package 数

func hammingWeight(num uint32) int {
	return HammingWeight(int(num)) // 二进制内部表示是相同的，因此 uint32 -> int 并不影响结果。
}

// HammingWeight 获取汉明重量。
func HammingWeight(num int) int {
	count := 0
	for num != 0 {
		num = removeLowestOne(num)
		count++
	}
	return count
}

// removeLowestOne 移除最低位1。
func removeLowestOne(num int) int {
	return num & (num - 1)
}
