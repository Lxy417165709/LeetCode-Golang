package main

func reverseWords(s string) string {
	bytes := []byte(s)
	bytes = trimBlank(bytes)
	// 翻转bytes数组中的每个单词
	reverseSingleWords(bytes)

	// 对bytes数组全体进行翻转
	reverse(bytes, 0, len(bytes)-1)
	return string(bytes)
}

// 翻转字符数组中的所有单词
// 单词 "abc" 翻转后变 "cba"
func reverseSingleWords(bytes []byte) {
	begin, end := 0, 0
	for end <= len(bytes) {
		if end == len(bytes) || bytes[end] == ' ' {
			reverse(bytes, begin, end-1)
			begin = end + 1
		}
		end++
	}
}

// 翻转字符数组的 [l,r] 区域
func reverse(bytes []byte, l, r int) {
	for i := 0; i < (r-l+1)/2; i++ {
		bytes[l+i], bytes[r-i] = bytes[r-i], bytes[l+i]
	}
}

// 修剪空格，使空格满足题目要求
// 返回修剪后的字符切片
func trimBlank(bytes []byte) []byte {
	reader, writer := 0, 0
	banFlag := true // 禁止空格标识: true表示禁止
	for reader < len(bytes) {
		if bytes[reader] != ' ' {
			banFlag = false
			bytes[writer] = bytes[reader]
			writer++
		} else {
			if banFlag == false {
				bytes[writer] = bytes[reader]
				writer++
			}
			banFlag = true
		}
		reader++
	}
	// 之所以这里要判断一下，是因为可能原字符串是"abc   d    "
	// 那么最终结果是 "abc d " (因为我们无法判断d后面是否还有单词，所以d后面会加上空格)
	if writer >= 1 && bytes[writer-1] == ' ' {
		writer--
	}
	return bytes[:writer]
}

/*
	总结
	1. 这道题总体思路是:
			(1) 先把字符串中的多余空格去除。 (原地算法，采用读写指针，对应的函数是: trimBlank)
			(2) 之后翻转字符串中的每个单词。 (原地算法，对应函数是: reverseSingleWords)
			(3) 最后把字符串整体进行翻转。	 (原地算法, 对应函数是: reverse)
		其中(2)(3)步骤可以对换。
*/
