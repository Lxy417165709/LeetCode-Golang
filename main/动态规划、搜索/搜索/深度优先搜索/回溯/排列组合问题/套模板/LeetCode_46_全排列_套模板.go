package 排列组合问题

var permuteSequence [][]int // 结果集

// 返回结果集的函数
func permute(nums []int) [][]int {
	/* 1. 进行一些预处理 */
	permuteSequence = make([][]int, 0)

	/* 2. 调用回溯函数 */
	permuteUniqueExec(nums, []int{})

	/* 5. 返回结果集 */
	return permuteSequence
}

// 回溯函数
func permuteUniqueExec(nums []int, sequence []int) {
	/* 3. 判断是否需要加入结果集以及进行剪枝 */
	if len(nums) == 0 {
		permuteSequence = append(permuteSequence, newSlice(sequence))
		return
	}


	for i := 0; i < len(nums); i++ {
		/* 4. 继续调用回溯函数，这里会有以下几种情况。*/
		// 因为题目要求的是排列数，且不可重复选取，所以处理如下。
		nums[0], nums[i] = nums[i], nums[0]
		permuteUniqueExec(nums[1:], append(sequence, nums[0]))
		nums[0], nums[i] = nums[i], nums[0]
	}
}

// 深拷贝
func newSlice(oldSlice []int) []int {
	slice := make([]int, len(oldSlice))
	copy(slice, oldSlice)
	return slice
}

/*
	题目链接:
		https://leetcode-cn.com/problems/permutations/			全排列
*/

/*
	总结
	1. 这题目的本质是: 给定一个数组，求数组的全排列。 (数组没有重复元素)
*/
