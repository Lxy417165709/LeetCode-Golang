package Geek

// -------------------- 遍历法 --------------------
// 执行用时：8 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：5.3 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度: O(n * log_n)
func findSpecialInteger(arr []int) int {
	minIndexSpanOfSpecialNum := len(arr) / 4
	for i := 0; i+minIndexSpanOfSpecialNum < len(arr); i++ {
		if arr[i] == arr[i+minIndexSpanOfSpecialNum] {
			return arr[i]
		}
	}
	panic("题目说一定存在，所以这句代码一定不会执行")
}

// -------------------- 二分法统计出现次数 --------------------
// 执行用时：12 ms,  在所有 Go 提交中击败了 81.54% 的用户
// 内存消耗：5.3 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度: O(n * log_n)
func findSpecialInteger(arr []int) int {
	for i := 0; i < len(arr); i++ {
		if getCountOfNum(arr, arr[i]) >= len(arr)/4+1 {
			return arr[i]
		}
	}
	panic("不可能到达这")
}

func getCountOfNum(arr []int, ref int) int {
	indexOfFirstEqual := getIndexOfFirstEqual(arr, ref)
	indexOfLastEqual := getIndexOfLastEqual(arr, ref)
	if indexOfFirstEqual == -1 {
		return 0
	}
	return indexOfLastEqual - indexOfFirstEqual + 1
}

func getIndexOfFirstEqual(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	isExist := false
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == ref {
			isExist = true
			right = mid - 1
		} else {
			if arr[mid] > ref {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	if isExist {
		return left
	}
	return -1
}

func getIndexOfLastEqual(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	isExist := false
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == ref {
			isExist = true
			left = mid + 1
		} else {
			if arr[mid] > ref {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	if isExist {
		return right
	}
	return -1
}


// -------------------- 哈希法统计出现次数(无代码) --------------------
// 略

/*
	题目链接: https://leetcode-cn.com/problems/element-appearing-more-than-25-in-sorted-array/
	总结:
		1. 因为数组是有序的，所以如果A[i]超过数组元素总数的1/4，假设其在数组的开始索引是 first,
			则其结束索引一定大于等于 first+len(arr)/4。 通过这个就可以来得出这个数了。
			(这里的除法是整除)
*/
