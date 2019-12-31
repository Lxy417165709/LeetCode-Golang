package main

func numRabbits(answers []int) int {
	count := make(map[int]int)	// 用于计次，记录有value只兔子说了自己和key只兔子颜色相同
	for i := 0; i < len(answers); i++ {
		count[answers[i]]++
	}
	ans := 0
	for k, v := range count {
		ans += (floor(v, k+1)) * (k + 1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 取 a/b 的长界
func floor(a, b int) int {
	if a%b == 0 {
		return a / b
	}
	return a/b + 1
}

/*
    总结
    1. 这道题要自己在草稿纸上模拟后，思路清晰了再写...
    2. 做这个题的时候是有思路的，就是因为没有理清，所以搞了好久...
	3. 这道题思路需要利用贪心思维。
*/
