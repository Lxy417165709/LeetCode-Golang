package main

func longestWord(words []string) string {
	isLiving := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		isLiving[words[i]] = true
	}
	return longestWordExec(isLiving, "")
}

func longestWordExec(isLiving map[string]bool, nowStr string) string {
	ans := nowStr
	for i := 'a'; i <= 'z'; i++ {
		str := nowStr + string(i)
		if isLiving[str] {
			str = longestWordExec(isLiving, str)
			if len(ans) < len(str) {
				ans = str
			}
		}
	}
	return ans
}
/*
	总结
	1. 该题我采用了搜索的方法AC。
	2. 官方有人用字典树的方法做这个题
	3. 官方有用排序+set做这道题的
*/
