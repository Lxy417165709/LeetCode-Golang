package 字符串

func replaceSpace(s string) string {
	return string(replaceSpaceOfChars([]byte(s)))
}

func replaceSpaceOfChars(chars []byte) []byte {
	// 1. 获取空格数。
	countOfSpace := 0
	for _, char := range chars {
		if char == ' ' {
			countOfSpace++
		}
	}

	// 2. 计算新字符集长度。
	lenOfNewChars := len(chars) - countOfSpace + 3*countOfSpace

	// 3. 构建新字符集。
	readIndex := len(chars) - 1
	writeIndex := lenOfNewChars - 1
	chars = append(chars, make([]byte, lenOfNewChars-len(chars))...)
	for readIndex >= 0 {
		if chars[readIndex] == ' ' {
			chars[writeIndex] = '0'
			writeIndex--
			chars[writeIndex] = '2'
			writeIndex--
			chars[writeIndex] = '%'
			writeIndex--
		} else {
			chars[writeIndex] = chars[readIndex]
			writeIndex--
		}
		readIndex--
	}

	// 4. 返回。
	return chars
}
