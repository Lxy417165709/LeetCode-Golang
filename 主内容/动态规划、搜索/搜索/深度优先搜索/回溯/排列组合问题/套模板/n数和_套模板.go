package 套模板

var combinationSequence [][]int // 结果集

func combinationSum3(k int, n int) [][]int {
	/* 1. 进行一些预处理 */
	candidates := make([]int, 9)
	combinationSequence = make([][]int, 0)
	for i := 1; i <= 9; i++ {
		candidates[i-1] = i
	}

	/* 2. 调用回溯函数 */
	combinationSumExec(candidates, n, k, make([]int, 0, 10))

	/* 5. 返回结果集 */
	return combinationSequence
}

// 有重复、不可以重复选、数值可能小于0
// 在上层需要先排序, k是总和, n是个数
func combinationSumExec(candidates []int, n int, k int, sequence []int) {

	/* 3. 判断是否需要加入结果集以及进行剪枝 */
	if n == 0 && k == 0 {
		combinationSequence = append(combinationSequence, newSlice(sequence))
		return
	}
	if n == 0{
		return
	}
	isVisit := make(map[int]bool) // 记录数字是否使用过，防止出现重复的结果
	for i := 0; i < len(candidates); i++ {
		if isVisit[candidates[i]] == true {
			continue
		}
		isVisit[candidates[i]] = true
		/* 4. 继续调用回溯函数 */
		// 因为题目要求的是组合数且不能重复选取，所以下一层处理的是 candidates[i+1:]
		combinationSumExec(candidates[i+1:], n-candidates[i], k-1, append(sequence, candidates[i]))
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
		https://leetcode-cn.com/problems/combination-sum-iii/submissions/			组合求和3
*/

/*
	总结
	1. 该题题意是:
			在一个不可重复选取、元素大于0、元素不存在重复的数组中，选出k个数，使它们的值等于n，输出所有的组合。
*/
