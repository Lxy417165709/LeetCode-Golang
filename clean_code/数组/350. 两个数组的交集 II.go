package 数组

func intersect(nums1 []int, nums2 []int) []int {
	countOfNums2 := getCountOfNum(nums2)
	intersection := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		if countOfNums2[nums1[i]] > 0 {
			intersection = append(intersection, nums1[i])
			countOfNums2[nums1[i]]--
		}
	}
	return intersection
}

func getCountOfNum(array []int) map[int]int {
	count := make(map[int]int)
	for i := 0; i < len(array); i++ {
		count[array[i]]++
	}
	return count
}

/*
	题目链接: https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/solution/liang-ge-shu-zu-de-jiao-ji-ii-by-leetcode-solution/
	总结:
		1. 其实可以把结果放在nums1中，这样就可以将时间复杂度优化为O(1)
		2. 也可以通过排序AC这个题
*/
