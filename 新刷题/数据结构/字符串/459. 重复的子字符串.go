package 字符串

import "strings"

// repeatedSubstringPattern 判断 s 内是否存在循环节。
// 题解: https://leetcode-cn.com/problems/repeated-substring-pattern/solution/jian-dan-ming-liao-guan-yu-javaliang-xing-dai-ma-s/
func repeatedSubstringPattern(s string) bool {
	ss := s + s // 使用 s + s，即可得出 s 所有移位情况。
	return strings.Contains(ss[1:len(ss)-1], s)	// ss[1:len(ss)-1] 排除了循环节是自身的情况。
}
