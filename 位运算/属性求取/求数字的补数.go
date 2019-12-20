package main

/*
	给定一个正整数，输出它的补数。补数是对该数的二进制表示取反。
	比如: 5 的二进制是: 101，则其补数为 2，对应于二进制就是010
*/

// 解法1 (模拟获取补数的过程)
func findComplement(num int) int {
	ans := 0
	count := uint8(0)
	for num != 0 {
		ans = ans | (((num & 1) ^ 1) << count)
		num >>= 1
		count++
	}
	return ans
}

// 解法2 (获取掩码，将该掩码与原num异或)
func findComplement(num int) int {
	mark := 0
	tmp := num
	for tmp != 0 {
		tmp >>= 1
		mark = (mark << 1) | 1
	}
	return mark ^ num
}

/*
	题目链接:
		https://leetcode-cn.com/problems/number-complement/submissions/		数字的补数
*/
