package 套模板

var subsetSequence [][]int    // 结果集

// 返回结果集的函数
func subsets(nums []int) [][]int {
	/* 1. 预处理 */
	subsetSequence = make([][]int, 0)

	/* 2. 调用回溯函数 */
	subsetsExec(nums, []int{})

	/* 5. 返回结果集 */
	return subsetSequence
}

func subsetsExec(nums []int, sequence []int) {
	/* 3. 判断是否需要加入结果集以及进行剪枝 */
	subsetSequence = append(subsetSequence, newSlice(sequence))

	for i := 0; i < len(nums); i++ {
		/* 4. 继续调用回溯函数 */
		// 因为题目要求的是组合数且不能重复选取，所以下一层处理的是 nums[i+1:]
		subsetsExec(nums[i+1:], append(sequence, nums[i]))
	}
}

func newSlice(slice []int) []int {
	s := make([]int, len(slice))
	copy(s, slice)
	return s
}

/*
	题目链接:
		https://leetcode-cn.com/problems/subsets/	子集
*/

/*
	总结
	1. 这题求的是数组中的子集，相当于求数组中的所有组合。
*/
