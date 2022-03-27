package 字符串

import (
	"strconv"
	"strings"
)

// restoreIpAddresses 获取字符串中合法的IP地址。
func restoreIpAddresses(s string) []string {
	result := parse(s, 1)
	addresses := make([]string, 0)
	for _, addressParts := range result {
		addresses = append(addresses, strings.Join(addressParts, "."))
	}
	return addresses
}

// parse 解析出所有合法的地址部分集。
func parse(s string, addressPartNum int) [][]string {
	// 1. 空返回。
	if len(s) == 0 {
		return nil
	}

	// 2. 如果是最后一级，此时字符串合法时返回，否则返回空。
	if addressPartNum == 4 && isValidAddressPart(s) {
		if isValidAddressPart(s) {
			return [][]string{{s}}
		}
		return nil
	}

	// 3. 获取地址部分集。
	var result [][]string
	for i := 1; i <= 3 && i <= len(s); i++ {
		// 3.1 本段非法则推出。 (由于本段不合法，添加字符后还是不合法，所以这里break)
		if !isValidAddressPart(s[:i]) {
			break
		}

		// 3.2 递归。
		for _, pre := range parse(s[i:], addressPartNum+1) {
			result = append(result, append([]string{s[:i]}, pre...))
		}
	}

	// 4. 返回。
	return result
}

// isValidAddressPart 是否是合法的IP地址部分。
func isValidAddressPart(s string) bool {
	// 1. 空字符串不合法。
	if len(s) == 0 {
		return false
	}

	// 2. 只有一位，合法。
	if len(s) == 1 {
		return true
	}

	// 3. 有前导零，不合法。
	if s[0] == '0' {
		return false
	}

	// 4. 判断大小。
	num, _ := strconv.Atoi(s)
	return num <= 255
}
