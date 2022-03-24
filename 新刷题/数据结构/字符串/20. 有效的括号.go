package 字符串

// rightSymbolToLeftSymbolMap 右符号到左符号的映射。
var rightSymbolToLeftSymbolMap = map[byte]byte{
	'}': '{',
	']': '[',
	')': '(',
}

// isValid 判断是否是有效的括号字符串。
func isValid(s string) bool {
	// 1. 新建栈。
	stack := make([]byte, 0)

	// 2. 迭代处理。
	for i := 0; i < len(s); i++ {
		// 2.1 判断是否是右符号，获取合适的左符号。
		symbol := s[i]
		fitLeftSymbol, isRightSymbol := rightSymbolToLeftSymbolMap[symbol]

		// 2.2 如果是右符号。
		if isRightSymbol {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != fitLeftSymbol {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}

		// 2.3 如果是左符号。
		stack = append(stack, symbol)
	}

	// 3. 返回。
	return len(stack) == 0
}

// errorFunctionOfIsValid 错误解法。 案例: "([)]"。
func errorFunctionOfIsValid(s string) bool {
	symbolCountMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		char := s[i]
		if leftSymbol, ok := rightSymbolToLeftSymbolMap[char]; ok {
			// 说明是右括号。
			symbolCountMap[leftSymbol]--
			if symbolCountMap[leftSymbol] < 0 {
				return false
			}
			continue
		}
		// 说明是左括号。
		symbolCountMap[char]++
	}

	for _, count := range symbolCountMap {
		if count != 0 {
			return false
		}
	}
	return true
}
