package 排列组合问题

// 回溯 + 外部变量 实现全排列
var permuteSequence [][]int
func permute(nums []int) [][]int {
	permuteSequence = make([][]int, 0,100)
	permuteUniqueExec(nums, 0, len(nums)-1)
	return permuteSequence
}

func permuteUniqueExec(nums []int, l int, r int) {
	if l == r {
		// 一定要加入深拷贝nums，否则之后对nums的修改会影响permuteSequence内部的切片
		permuteSequence = append(permuteSequence, newSlice(nums))
		return
	}
	for i := l; i <= r; i++ {
		nums[i], nums[l] = nums[l], nums[i]
		permuteUniqueExec(nums, l+1, r)
		nums[i], nums[l] = nums[l], nums[i]
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