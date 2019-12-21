package main

import (
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	// [0,i-1] 表示第一个数字
	// [i,t]表示第二个数字
	for i := 1; i < len(num); i++ {
		for t := i; t < len(num); t++ {
			if isAdditiveNumberExec(num[t+1:], num[:i], num[i:t+1]) == true {
				return true
			}
		}
	}
	return false
}

// 没有优化的递归搜索
func isAdditiveNumberExec(remainDigits string, first string, second string) bool {
	if !isValid(first) || !isValid(second) {
		return false
	}

	a, _ := strconv.Atoi(first)
	b, _ := strconv.Atoi(second)
	if isValid(remainDigits) {
		c, _ := strconv.Atoi(remainDigits)
		if a+b == c {
			return true
		}
	}
	ans := false
	// 这里其实可以不用for循环，这是因为由于我们知道了a+b的结果，那么我们可以从结果出发限定下一个搜索的字符串长度
	// 接下来写一份优化
	for i := 0; i <= len(remainDigits); i++ {
		if !isValid(remainDigits[:i]) {
			continue
		}
		c, _ := strconv.Atoi(remainDigits[:i])
		if a+b == c {
			ans = ans || isAdditiveNumberExec(remainDigits[i:], second, remainDigits[:i])
		}
	}
	return ans
}

// 判断数字是否合法
func isValid(a string) bool {
	if len(a) == 0 {
		return false
	}
	if len(a) == 1 {
		return true
	}
	return a[0] != '0'
}

// 优化的递归搜索
func isAdditiveNumberExec(remainDigits string, first string, second string) bool {
	if !isValid(first) || !isValid(second) {
		return false
	}

	a, _ := strconv.Atoi(first)
	b, _ := strconv.Atoi(second)
	resultStr := strconv.Itoa(a + b)
	ans := false
	// 如果相等，表示已经组成累加序列了，于是结果为true
	if resultStr == remainDigits {
		ans = true
	} else {
		// 如果 resultStr 不是 remainDigits 的前缀，那么肯定不可能组成 累加序列，于是结果为false
		if strings.Index(remainDigits, resultStr) != 0 {
			ans = false
		} else {
			// 如果 resultStr 是 remainDigits 的前缀，那么就继续往下找
			ans = isAdditiveNumberExec(remainDigits[len(resultStr):], second, resultStr)
		}
	}
	return ans
}

/*
	题目链接:
		https://leetcode-cn.com/problems/additive-number/			累加数
*/

/*
	总结
	1. 这题只需要枚举出前2个数，之后剩下的字符串再进行递归搜索就可以了。
	2. 回溯有时候可以通过当前的结果来指明方向，对于这道题，当前结果就是 first + second 的总和，而指明的方向就是
	   接下来需要回溯的字符串位于remainDigits的哪一部分。
	3. 本来这题还使用记忆化的，但是其实这题没有重叠子，于是便删除了记忆化了。
*/
