package 数学

import "sort"

func numEquivDominoPairs(dominoes [][]int) int {
	makeAllDominoesWidthGreaterThanOrEqualHeight(dominoes)
	countOfEquivalenceDomino := getCountOfEquivalenceDomino(dominoes)
	return getCountOfEquivalenceDominoPair(countOfEquivalenceDomino)
}

func makeAllDominoesWidthGreaterThanOrEqualHeight(dominoes [][]int) {
	for i := 0; i < len(dominoes); i++ {
		makeWidthGreaterThanOrEqualHeight(dominoes[i])
	}
}

func getCountOfEquivalenceDomino(dominoes [][]int) map[int]int {
	countOfEquivalenceDomino := make(map[int]int)
	for i := 0; i < len(dominoes); i++ {
		countOfEquivalenceDomino[getHashCode(dominoes[i])]++
	}
	return countOfEquivalenceDomino
}

func getCountOfEquivalenceDominoPair(countOfEquivalenceDominoe map[int]int) int {
	result := 0
	for _, count := range countOfEquivalenceDominoe {
		result += count * (count - 1) / 2
	}
	return result
}

func makeWidthGreaterThanOrEqualHeight(domino []int) {
	sort.Ints(domino)
}

func getHashCode(domino []int) int {
	const overLength = 10
	return domino[0]*overLength + domino[1]
}

/*
	题目链接: https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/
*/
