package 未归类

var charIndexsOfS [][]int // 原先采用哈希，这次试用二维数组，速度快了1倍
func numMatchingSubseq(S string, words []string) int {
	charIndexsOfS = make([][]int, 200)
	for i := 0; i < len(S); i++ {
		charIndexsOfS[S[i]] = append(charIndexsOfS[S[i]], i)
	}
	countOfSubSequence := 0
	for _, word := range words {
		if isSubSequence(word) {
			countOfSubSequence++
		}
	}
	return countOfSubSequence
}

func isSubSequence(word string) bool {
	reachIndex := -1
	for i := 0; i < len(word); i++ {
		charIndexs := charIndexsOfS[word[i]]
		index := getIndexOfFirstGreater(charIndexs, reachIndex)
		if index == len(charIndexs) {
			return false
		}
		reachIndex = charIndexs[index]
	}
	return true
}

func getIndexOfFirstGreater(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if ref == arr[mid] {
			left = mid + 1
		} else {
			if ref > arr[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return left
}
