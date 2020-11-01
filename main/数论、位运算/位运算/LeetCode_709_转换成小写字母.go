package main

func toLowerCase(str string) string {
	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		bytes[i] |= 32

	}
	return string(bytes)
}
/*
	总结
	1. 这里利用大小写字符的阿斯克码的特点，通过位运算实现字符串(仅包含字母)大小写的转换。
	2. 		大写变小写、小写变大写 : 字符 ^= 32;
			大写变小写、小写变小写 : 字符 |= 32;
			小写变大写、大写变大写 : 字符 &= -33;
*/
