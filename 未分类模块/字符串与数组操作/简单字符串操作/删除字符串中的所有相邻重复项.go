package main

/*
	给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
	在 S 上反复执行重复项删除操作，直到无法继续删除。
	在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。
*/

func removeDuplicates(S string) string {
	stack := make([]byte, 0, len(S))
	for i := 0; i < len(S); i++ {
		// 遇到相同的则出栈，否则入栈
		if len(stack) != 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}

/*
	题目链接:
		https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/		删除字符串中的所有相邻重复项
*/

/*
	总结
	1. 这题是看了别人的答案后，自己采用栈AC的。 (之前采用了2个指针的方式，很不优美)
*/
