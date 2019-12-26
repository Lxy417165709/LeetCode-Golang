package main

// 简单二分查找: 从递增数组中查找目标值的索引，不存在则返回-1。
// 例子: binarySearch([]int{1,2,3,4,5}, 1) -> return 0
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		// mid := (l + r)/2 这样写也可以，但是l + r很大时可能造成溢出。
		// 所以我采用 mid := l + (r-l)/2 这种写法防止溢出。
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else {
			if nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

/*
	题目链接:
		https://leetcode-cn.com/problems/binary-search/

*/
