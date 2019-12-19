package main

func numEquivDominoPairs(dominoes [][]int) int {
	count := make(map[int]int)
	ans := 0
	for i := 0; i < len(dominoes); i++ {
		hashNumber := hash(dominoes[i])
		ans += count[hashNumber]
		count[hashNumber]++
	}
	return ans
}

func hash(dominoe []int) int {
	a, b := dominoe[0], dominoe[1]
	if a > b {
		a, b = b, a
	}
	return (b << 5) | a
}

/*
	题目链接:
		https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/		等价多米诺骨牌对的数量
*/

/*
	总结
	1. 这题目我是使用哈希解决的，思路就是把一个骨牌哈希为一个数，如果两个骨牌哈希值相同，那么他们就是等价的。
	2. 这个还用到了排列组合的知识: 加入现在有n个骨牌是等价的，那么再加入1个等价骨牌，那么等价骨牌对数增加n。
*/
