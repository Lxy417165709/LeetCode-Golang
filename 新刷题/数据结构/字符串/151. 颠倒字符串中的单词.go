package 字符串

// reverseWords	颠倒字符串中的单词。
func reverseWords(s string) string {
	// 1. 转换。
	bytes := []byte(s)

	// 2. 翻转所有单词的位置。
	reverseAllWordsPosition(bytes)

	// 3. 移除多余空格，返回最终索引。
	end := removeExtraBlank(bytes)

	// 4. 返回。
	return string(bytes[:end+1])
}

// removeExtraBlank 移除多余空格，返回最终索引。
func removeExtraBlank(bytes []byte) int {
	// 1. 消除单词之间的多余空格。 (每次循环: 写单词 + 写一个空格(reader 未越界))
	writer, reader := 0, 0
	for reader < len(bytes) {
		// 1.1 reader 指向第一个非空格字符。
		for reader < len(bytes) && bytes[reader] == ' ' {
			reader++
		}

		// 1.2 将 reader 写入 writer 指向的位置，直到 reader 指向第一个空格字符。
		for reader < len(bytes) && bytes[reader] != ' ' {
			bytes[writer] = bytes[reader]
			writer++
			reader++
		}

		// 1.3 如果还未越界，添加单个空格分隔单词。
		if reader < len(bytes) {
			bytes[writer] = bytes[reader] // 此时 reader 指向空格。
			writer++
			reader++
		}
	}

	// 2. 消除 步骤3 遗留在尾部的空格。 (这种场景出现的原因: bytes 尾部具有空格。)
	end := writer - 1
	if bytes[end] == ' ' {
		end--
	}

	// 3. 返回。
	return end
}

// reverseAllWordsPosition 翻转所有单词的位置。
func reverseAllWordsPosition(bytes []byte) {
	// 1. 翻转整个字符串。
	reverseBytes(bytes)

	// 2. 翻转字符串的每个单词。
	left, right := 0, 0 // 单词位于 bytes[left, right)
	for right < len(bytes) {
		// 2.1 左右指针指向非空位置。
		for left < len(bytes) && bytes[left] == ' ' {
			left++
		}
		right = left

		// 2.2 右指针指向下一个空位置。
		for right < len(bytes) && bytes[right] != ' ' {
			right++
		}

		// 2.3 翻转左右指针之间的单词。
		reverseBytes(bytes[left:right])

		// 2.4 处理下一个单词。
		left = right
	}
}

// reverseBytes 翻转字符串。
func reverseBytes(bytes []byte) {
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
	}
}
