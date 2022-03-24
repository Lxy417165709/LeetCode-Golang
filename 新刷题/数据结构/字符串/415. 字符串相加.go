package 字符串

import "fmt"

func addStrings(num1 string, num2 string) string {
	readIndexOfNum1 := len(num1) - 1
	readIndexOfNum2 := len(num2) - 1
	result := ""
	carry := 0
	for readIndexOfNum1 >= 0 || readIndexOfNum2 >= 0 || carry != 0 {
		sum := carry
		if readIndexOfNum1 >= 0 {
			sum += int(num1[readIndexOfNum1] - '0')
			readIndexOfNum1--
		}
		if readIndexOfNum2 >= 0 {
			sum += int(num2[readIndexOfNum2] - '0')
			readIndexOfNum2--
		}
		digit := sum % 10
		result = fmt.Sprintf("%d", digit) + result
		carry = sum / 10
	}
	return result
}
