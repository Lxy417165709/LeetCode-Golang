package main

// 进阶二分查找1: 从递增数组中找到第一个 ＞ 目标值的数的索引, 当目标值≥数组中的最大值时，返回数组的长度。
// 例子: firstGreater([]int{1,2,3,4,5}, 4) -> return 4
func firstGreater(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			// 相等时，说明数组 mid 右边的值都 ≥ target，于是为了找到第一个大于target的数的索引，
			// 我们需要将左指针右移，继续到数组右边找。
			l = mid + 1
		} else {
			if arr[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	// 循环体出来后，必然有 r + 1 == l
	// 此时 l 指向了第一个大于目标值的数，r 指向了最后一个小于等于目标值的数。
	// 所以此时我们返回 l。当然，如果需要的是最后一个小于等于目标值的数的索引，此时应该返回 r。
	return l
}

// 进阶二分查找2: 从递增数组中找到第一个 ≥ 目标值的数的索引, 当目标值＞数组中的最大值时，返回数组的长度。
// 例子: firstGreaterOrEqual([]int{1,2,3,4,5}, 4) -> return 3
func firstGreaterOrEqual(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			// 相等时，说明数组 mid 右边的值都 ≥ target，于是为了找到第一个 ≥ target 的数的索引，
			// 我们需要将右指针指针左移，继续到数组左边找。
			r = mid - 1
		} else {
			if arr[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	// 循环体出来后，必然有 r + 1 == l
	// 此时 l 指向了第一个大于等于目标值的数，r 指向了最后一个小于目标值的数。
	// 所以此时我们返回 l。当然，如果需要的是最后一个小于目标值的数的索引，此时应该返回 r。
	return l
}

// 进阶二分查找3: 最后一个小于等于
// 例子: lastLessOrEqual([]int{1,2,3,4,5}, 4) -> return 3
func lastLessOrEqual(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			l = mid + 1
		} else {
			if arr[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	// firstGreater 返回 l, 而 lastLessOrEqual 返回 r。
	return r
}

// 进阶二分查找4: 最后一个小于
// 例子: lastLess([]int{1,2,3,4,5}, 4) -> return 2
func lastLess(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			r = mid - 1
		} else {
			if arr[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	// firstGreaterOrEqual 返回 l, 而 lastLess 返回 r。
	return r
}



/*
	题目链接:
		https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
		https://leetcode-cn.com/problems/search-a-2d-matrix/		搜索二维矩阵
*/



/*
	总结
	1. 对于二分搜索来说: 我们要找的是 "第一个大于"、"第一个大于等于"、"最后一个小于"、"最后一个小于等于" 目标值的数的索引。
		(当寻找的是 "最后一个大于"、"最后一个大于等于" 目标值的数的索引时，我们只需要将数组的最后一个元素与目标值做比较)
		(当寻找的是 "第一个小于"、"第一个小于等于" 目标值的数的索引时，我们只需要将数组的第一个元素与目标值做比较)
	2. 进阶的二分搜索与简单二分搜索的代码几乎一样，但也有以下区别。
			- [ ] 进阶二分搜索在 arr[mid] == target 时并不立即返回，而是根据条件继续向左或向右查找。
			- [ ] 进阶二分搜索退出循环后并不是返回 -1，而是返回对应的左指针或右指针。
			- [ ] 进阶二分搜索返回左指针时，此时我们寻找的就是"第一个大于"或"第一个大于等于"。
			- [ ] 进阶二分搜索返回右指针时，此时我们寻找的就是"最后一个小于"或"最后一个小于等于"。
*/
