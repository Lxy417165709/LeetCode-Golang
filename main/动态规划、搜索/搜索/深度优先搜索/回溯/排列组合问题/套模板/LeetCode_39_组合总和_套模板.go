package 套模板


var combinationSequence [][]int    // 结果集

// 返回结果集的函数
func combinationSum(candidates []int, target int) [][]int {
	/* 1. 进行一些预处理 */
	combinationSequence = make([][]int, 0)

	/* 2. 调用回溯函数 */
	combinationSumExec(candidates, target, make([]int, 0, 100))

	/* 5. 返回结果集 */
	return combinationSequence
}

// 回溯函数
func combinationSumExec(candidates []int, target int, sequence []int) {
	/* 3. 判断是否需要加入结果集以及进行剪枝 */
	// 剪枝
	if target < 0 {
		return
	}
	if target == 0 {
		combinationSequence = append(combinationSequence, newSlice(sequence))
		return
	}

	for i := 0; i < len(candidates); i++ {
		/* 4. 继续调用回溯函数 */
		// 因为题目要求的是组合数且能重复选取，所以下一层处理的是 candidates[i:]
		combinationSumExec(candidates[i:], target-candidates[i], append(sequence, candidates[i]))
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
		https://leetcode-cn.com/problems/combination-sum/			组合求和
		https://leetcode-cn.com/problems/coin-change-2/				零钱兑换2
*/

/*
	总结
	1. 切片可以初始化容量，这样可以避免当切片长度大于容量时，切片扩容导致的额外代价。
	2. 该题和零钱兑换2几乎一样，只是零钱兑换2求的是种数，而这里需要输出所有的组合。 (零钱兑换那使用了动态规划，这里使用了回溯)
	3. 该题题意是: 在一个可重复选取、元素大于0的数组中，选出一些数，使它们的值等于target，输出所有的组合。
*/
