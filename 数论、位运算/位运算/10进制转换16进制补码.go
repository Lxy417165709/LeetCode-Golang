package main

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	resolve := "0123456789abcdef" // 相当于一个映射
	ans := ""
	// 对于正数来说，位移运算可以保证num==0。
	// 但对负数来说，位移运算并不能保证num==0，而是需要使用32位int保证（对应16进制小于等于8位）。
	for num != 0 && len(ans) < 8 {
		ans = string(resolve[num&0xf]) + ans		// &0xf 相当于 %16
		num >>= 4 // ">>" 是算数位移，其中正数右移左边补0，负数右移左边补1。 >>4 相当于 /16
	}
	return ans
}

/*
	题目链接:
		https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/comments/		数字转换为十六进制数
*/

/*
	总结
	1. 这个模板改一下就能实现10进制到其他进制的转换了。
*/
