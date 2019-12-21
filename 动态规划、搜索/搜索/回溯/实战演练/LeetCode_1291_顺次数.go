package main

import "sort"

/*
	我们定义「顺次数」为：每一位上的数字都比前一位上的数字大 1 的整数。
	请你返回由 [low, high] 范围内所有顺次数组成的 有序 列表（从小到大排序）。
*/
// 外部变量 + 回溯解法
var sequence []int
func sequentialDigits(low int, high int) []int {
	sequence = make([]int, 0)
	for i := 1; i <= 9; i++ {
		sequentialDigitsExec(0, i, low, high)
	}
	sort.Ints(sequence)
	return sequence
}

// num: 表示先前的顺次数
// digit: 表示接下来要添加到 先前的顺次数 末尾的数字
// low,high: 结果的上下界[low, high]
func sequentialDigitsExec(num int, digit int, low int, high int) {
	num = num*10 + digit
	if num > high {
		return
	}
	if num >= low {
		sequence = append(sequence, num)
	}
	if digit < 9 {
		sequentialDigitsExec(num, digit+1, low, high)
	}
}

// 另外一种递归写法
// num: 表示当前的顺次数
// digit: 表示接下来要添加到 当前的顺次数 末尾的数字
// low,high: 结果的上下界[low, high]
func sequentialDigitsExec(num int, digit int, low int, high int) {
	if num >= low && num<=high{
		sequence = append(sequence,num)
	}
	// 剪枝
	if num > high {
		return
	}
	if digit<=9{
		sequentialDigitsExec(num*10+digit,digit+1,low,high)
	}
}


// 迭代解法
func sequentialDigits(low int, high int) []int {
	sequence := make([]int, 0)
	for i := 1; i <= 9; i++ {
		num := i
		for t:=i+1;t<=9;t++{
			num = num*10 + t
			if num>=low && num<=high{
				sequence = append(sequence,num)
			}
		}
	}
	sort.Ints(sequence)
	return sequence
}



/*
	题目链接:
		https://leetcode-cn.com/problems/sequential-digits/			顺次数
*/
/*
	总结:
	1. 这种题型之前也遇到过，但是没有去认真的求解，今天遇到了，然后还是不会写...于是看了答案之后自己写出了上面的代码。
	2. 个人感觉这种类型的题目，如果规则简单的话，迭代很容易做出，但是规则比较复杂，那么最好使用递归求解。
	3. 官方还有人用滑动窗口做，就是先穷举出 123456789,再在这个数用滑动窗口。
*/
