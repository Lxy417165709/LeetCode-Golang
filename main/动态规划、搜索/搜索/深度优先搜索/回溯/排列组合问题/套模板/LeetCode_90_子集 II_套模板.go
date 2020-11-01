package 套模板

import "sort"


var subsetSequence [][]int	// 结果集

// 返回结果集的函数
func subsetsWithDup(nums []int) [][]int {
	/* 1. 预处理 */
	subsetSequence = make([][]int, 0)
	sort.Ints(nums) 	// 有重复元素的组合问题就要排序。为什么要排序呢？后面会说

	/* 2. 调用回溯函数 */
	subsetsExec(nums, []int{})

	/* 5. 返回结果集 */
	return subsetSequence
}
// 回溯函数
func subsetsExec(nums []int, sequence []int) {
	/* 3. 判断是否需要加入结果集以及进行剪枝 */
	subsetSequence = append(subsetSequence, newSlice(sequence))

	isVisit := make(map[int]bool) // 记录数字是否使用过，防止出现重复的结果
	for i := 0; i < len(nums); i++ {
		if isVisit[nums[i]] {
			continue
		}
		isVisit[nums[i]] = true

		/* 4. 继续调用回溯函数 */
		// 因为题目要求的是组合数且不能重复选取，所以下一层处理的是 nums[i+1:]
		subsetsExec(nums[i+1:], append(sequence, nums[i]))
	}
}

// 深拷贝
func newSlice(slice []int) []int {
	s := make([]int, len(slice))
	copy(s, slice)
	return s
}

/*
	题目链接:
		https://leetcode-cn.com/problems/subsets-ii/submissions/	子集2
*/

/*
	总结
	1. 这题求的是数组中的子集，相当于求数组中的所有组合。
	2. 这题要求的数组是有重复元素的，之前求的是没有重复元素的，上面已经标识出了增加了代码的部分。
*/
