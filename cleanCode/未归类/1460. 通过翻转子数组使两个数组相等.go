package 未归类

import "sort"

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(arr)
	sort.Ints(target)
	return areArraysSame(arr, target)
}

func areArraysSame(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

/*
	题目链接: https://leetcode-cn.com/problems/make-two-arrays-equal-by-reversing-sub-arrays/submissions/
	总结
	1. 这题其实可以采用哈希，虽然时间复杂度变为了O(n)，但空间复杂度也变为O(n)了
	2. 刚刚开始没有注意到arr其实不是升序的，所以WA了。 ---> 要注意题目条件
*/
