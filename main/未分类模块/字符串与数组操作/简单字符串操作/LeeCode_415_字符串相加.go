package main

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
	1. 这是看了评论区大佬后写出来的代码，这个模板很优雅，而且还可以用到链表相加中。
*/
