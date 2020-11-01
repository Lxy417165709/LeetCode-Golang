package 整数问题

import (
	"fmt"
	"sort"
)

/*
	如果一个整数上的每一位数字与其相邻位上的数字的绝对差都是 1，那么这个数就是一个「步进数」。
	给你两个整数，low 和 high，请你找出在 [low, high] 范围内的所有步进数，并返回 排序后 的结果。
*/

// 外部变量 + 回溯解法
var sequence []int
func countSteppingNumbers(low int, high int) []int {
	sequence = make([]int, 0)
	// 枚举，从第一位为 1 开始枚举到第一位为 9
	for i := 1; i <= 9; i++ {
		countSteppingNumbersExec(0, i, low, high)
	}

	// 官方认为0也是步进数
	if low == 0{
		sequence = append(sequence, 0)
	}
	sort.Ints(sequence)
	return sequence
}

// num: 表示先前的顺次数
// digit: 表示接下来要添加到 先前的顺次数 末尾的数字
// low,high: 结果的上下界[low, high]
func countSteppingNumbersExec(num int, digit int, low int, high int) {
	num = num*10 + digit
	// 剪枝
	if num > high {
		return
	}
	if num >= low {
		sequence = append(sequence, num)
	}
	if digit >= 0 && digit <= 9 {
		countSteppingNumbersExec(num, digit+1, low, high)
		countSteppingNumbersExec(num, digit-1, low, high) // 和顺次数相比就加了这句
	}
}

func main() {
	fmt.Println(countSteppingNumbers(0, 10000000000000))
}

/*
	题目链接:
		https://leetcode-cn.com/problems/stepping-numbers/			步进数(被锁了)
*/

/*
	总结
	1. 由于这题被锁了，所以我不能验证以上程序的正确性，但是输出结果是符合步进数要求的。
	2. 小伙伴可以把这题和顺次数进行比较，其实就加了一句话。
	3. 由于步进数规则比顺次数复杂，所以我只写出了递归的解法。
	4. 由于题目说0也是步进数，所以如果0在[low,high]范围内，也要把0加入到结果集中。(顺次数的输入下界是10，所以不用考虑)
*/
