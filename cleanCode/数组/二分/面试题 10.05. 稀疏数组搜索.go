package 二分

func findString(words []string, s string) int {
	return getIndex(words, 0, len(words)-1, s)
}

func getIndex(words []string, left, right int, s string) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if words[mid] == s {
		return mid
	}
	if words[mid] == "" {
		leftSearchResult := getIndex(words, left, mid-1, s)
		rightSearchResult := getIndex(words, mid+1, right, s)
		if leftSearchResult != -1 {
			return leftSearchResult
		}
		if rightSearchResult != -1 {
			return rightSearchResult
		}
		return -1
	}
	if words[mid] > s {
		return getIndex(words, left, mid-1, s)
	} else {
		return getIndex(words, mid+1, right, s)
	}
}
