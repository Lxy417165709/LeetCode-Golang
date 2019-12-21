package main

func convertToBase7(num int) string {
	if num==0{
		return "0"
	}
	isOpt := true
	if num < 0 {
		isOpt = false
		num = -num
	}
	resolve := "0123456"
	ans :=""
	for num!=0{
		ans = string(resolve[num%7])+ans
		num = num / 7
	}
	if isOpt==false{
		ans = "-" + ans
	}
	return ans
}

/*
	题目链接:
		https://leetcode-cn.com/problems/base-7/comments/		七进制数
*/

/*
	总结
	1. 这个模板改一下就能实现10进制到其他进制的转换了。
*/
