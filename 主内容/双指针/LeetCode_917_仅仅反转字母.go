package 双指针

func reverseOnlyLetters(S string) string {
	s := []byte(S)
	l, r := 0, len(s)-1
	for l <= r {
		for l <= r && !judge(s[l]) {
			l++
		}
		for l <= r && !judge(s[r]) {
			r--
		}
		if l <= r {
			s[l], s[r] = s[r], s[l]
			r--
			l++
		}
	}
	return string(s)
}

func judge(a uint8) bool {
	return a >= 'a' && a <= 'z' || a <= 'Z' && a >= 'A'
}
/*
	总结
	1. 水题，采用左右指针直接AC。
*/
