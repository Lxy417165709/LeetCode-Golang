package 读写指针

func defangIPaddr(address string) string {
	dotCount := 0
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			dotCount++
		}
	}
	bytes := make([]byte, dotCount*2+len(address))
	copy(bytes, []byte(address))
	reader, writer := len(address)-1, len(bytes)-1
	for reader >= 0 {
		if bytes[reader] == '.' {
			bytes[writer] = ']'
			writer--
			bytes[writer] = '.'
			writer--
			bytes[writer] = '['
			writer--
		} else {
			bytes[writer] = bytes[reader]
			writer--
		}
		reader--
	}
	return string(bytes)

}

/*
	总结
	1. 这题和 1089. 复写0 是类似的..
	2. 还有更方便的方法就是正向扫描 address, 再用一个额外的 string 存储结果。
*/
