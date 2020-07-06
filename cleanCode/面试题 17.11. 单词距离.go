package main

const INF = 100000000

func findClosest(words []string, word1 string, word2 string) int {
	indexs := make(map[string][]int)
	for index, word := range words {
		indexs[word] = append(indexs[word], index)
	}
	return getMinAbsDiffBetweenTwoSortedArrays(indexs[word1], indexs[word2])
}

func getMinAbsDiffBetweenTwoSortedArrays(arr1, arr2 []int) int {
	if len(arr1) == 0 || len(arr2) == 0 {
		panic("不存在最小绝对距离")
	}
	p1, p2 := 0, 0
	minAbsDiff := INF
	for p1 < len(arr1) && p2 < len(arr2) {
		absDiff := abs(arr1[p1] - arr2[p2])
		minAbsDiff = min(minAbsDiff, absDiff)
		if arr1[p1] > arr2[p2] {
			p2++
		} else {
			p1++
		}
	}
	return minAbsDiff
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}


/*
	总结
		1. 这题还能使用 空间O(1) 的解法，详情看题解，很简单。
		2. 这题的抽象模型就是: 指定2个数，找出它们在数组中的最小距离。
*/
