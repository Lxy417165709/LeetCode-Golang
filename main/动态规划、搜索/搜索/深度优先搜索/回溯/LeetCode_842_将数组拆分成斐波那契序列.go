package main

import "strconv"

func splitIntoFibonacci(S string) []int {
	// 第一个数字: [0, i-1]
	// 第二个数字: [i, t]
	// 思路是: 先划分出斐波那契数列的前2项，之后对余下字符串进行递归
	for i := 1; i < len(S); i++ {
		for t := i; t < len(S); t++ {
			ans := splitIntoFibonacciExec(S[t+1:], []string{S[:i], S[i : t+1]})
			// 表示形成斐波那契序列了
			if len(ans) >= 3 {
				return toIntSlice(ans)
			}
		}
	}
	return []int{}
}
func toIntSlice(strSlice []string) []int {
	intSlice := []int{}
	for i := 0; i < len(strSlice); i++ {
		x, _ := strconv.Atoi(strSlice[i])
		intSlice = append(intSlice, x)
	}
	return intSlice
}

// 递归函数
func splitIntoFibonacciExec(S string, sequence []string) []string {
	if len(S) == 0 {
		return sequence
	}
	firstNumber, secondNumber := sequence[len(sequence)-2], sequence[len(sequence)-1]
	if isValid(firstNumber) == false || isValid(secondNumber) == false {
		return []string{}
	}
	sum := addStrings(firstNumber, secondNumber)
	if isValid(sum) == false {
		return []string{}
	}

	if len(S) >= len(sum) && sum == S[:len(sum)] {
		return splitIntoFibonacciExec(S[len(sum):], append(sequence, sum))
	}
	return []string{}
}

// 判断作为数值的字符串是否合法
func isValid(s string) bool {
	if len(s) == 0 || len(s) > 10 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	x, _ := strconv.Atoi(s)
	return x <= (1<<31)-1 && s[0] != '0'
}

// 字符串相加
func addStrings(num1 string, num2 string) string {
	sum := make([]byte, 0)
	i, j, carry := len(num1)-1, len(num2)-1, 0
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(num2[j] - '0')
			j--
		}
		sum = append(sum, byte(carry%10+'0'))
		carry /= 10
	}
	reverse(sum, 0, len(sum)-1)
	return string(sum)
}

func reverse(bytes []byte, l, r int) {
	for i := 0; i < (r-l+1)/2; i++ {
		bytes[l+i], bytes[r-i] = bytes[r-i], bytes[l+i]
	}
}

/*
	总结
	1. 这题和 306. 累加数 是一样的
	2. 这个代码更有普遍性，适用于更长的字符串。 (因为我上面采用了字符串相加，而不是简单的整数相加)
*/
