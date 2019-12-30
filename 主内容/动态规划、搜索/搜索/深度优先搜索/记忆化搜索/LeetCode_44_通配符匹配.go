package main

func isMatch(s string, p string) bool {
	isVisit = make(map[int]bool)
	return isMatchExec(s, p, len(s), len(p))
}

var isVisit map[int]bool

func hash(ends, endp int) int {
	return (ends << 20) | endp
}

func isMatchExec(s string, p string, ends, endp int) bool {
	if ends == 0 && endp == 0 {
		return true
	}
	if endp == 0 {
		return false
	}
	if ends == 0 {
		for i := 0; i < endp; i++ {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}
	hashNumber := hash(ends, endp)
	if x, ok := isVisit[hashNumber]; ok {
		return x
	}

	ans := false
	if s[ends-1] == p[endp-1] || p[endp-1] == '?' {
		ans = isMatchExec(s, p, ends-1, endp-1)
	} else {
		if p[endp-1] == '*' {
			for i := ends; i >= 0 && ans == false; i-- {
				ans = isMatchExec(s, p, i, endp-1)
			}
		}
	}

	isVisit[hashNumber] = ans
	return ans
}

/*
	总结
	1. 第一次写想到了记忆化搜索，但是哈希函数采用的是字符串哈希，所以超时了。
	2. 第二次就用数字哈希，然后就AC了。
	3. 所以尽量不要采用字符串哈希..很慢..
	4. 这题和编辑距离类似，都是匹配问题。
	5. 这题官方还有使用暴力法、DP做的。
*/
