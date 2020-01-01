package main

func hasGroupsSizeX(deck []int) bool {
	count := make(map[int]int)
	for i := 0; i < len(deck); i++ {
		count[deck[i]]++
	}
	g := 0
	for _, v := range count {
		g = gcd(g, v)
	}

	return g >= 2
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

/*
	总结
	1. 这题 最大公约数 很关键。。。。
	2. 思路就是求出所有 卡牌 出现次数的最大公约数，判断这个最大公约数是否>=2，大于则true，否则则false
*/
