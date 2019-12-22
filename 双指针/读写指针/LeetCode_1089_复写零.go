package main

func duplicateZeros(arr []int) {
	beforeLength, afterLength := len(arr), 0

	// 获取拓展后的长度
	for i := 0; i < beforeLength; i++ {
		if arr[i] == 0 {
			afterLength += 2
		} else {
			afterLength += 1
		}
	}
	writer, reader := afterLength-1, beforeLength-1

	for reader >= 0 {
		// 遇到0则复写
		if arr[reader] == 0 {
			// 因为超出beforeLength的字符不需要写，所以这里要判断一下是否需要写入
			if writer < beforeLength {
				arr[writer] = 0
			}
			writer--
			// 因为超出beforeLength的字符不需要写，所以这里要判断一下是否需要写入
			if writer < beforeLength {
				arr[writer] = 0
			}
			writer--
		} else {
			// 因为超出beforeLength的字符不需要写，所以这里要判断一下是否需要写入
			if writer < beforeLength {
				arr[writer] = arr[reader]
			}
			writer--
		}
		reader--
	}
}

/*
	总结
	1. 这题目和剑指offer的空格替换很像，都是要求原地替换字符，而且做法也是类似的。
*/
