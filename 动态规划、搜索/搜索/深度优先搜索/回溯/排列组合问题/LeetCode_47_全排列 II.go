package 排列组合问题

// 回溯 + 外部变量 实现全排列
var permuteSequence [][]int
func permuteUnique(nums []int) [][]int {
	permuteSequence = make([][]int, 0)
	permuteUniqueExec(nums, 0, len(nums)-1)
	return permuteSequence
}

func permuteUniqueExec(nums []int, l int, r int) {
	if l == r {
		// 一定要加入深拷贝nums，否则之后对nums的修改会影响permuteSequence内部的切片
		permuteSequence = append(permuteSequence, newSlice(nums))
		return
	}
	isVisited := make(map[int]bool)
	for i := l; i <= r; i++ {
		// 因为nums中有重复数字，所以要标识数字是否已被使用
		if isVisited[nums[i]] == true {
			continue
		}
		isVisited[nums[i]] = true
		nums[i], nums[l] = nums[l], nums[i]
		permuteUniqueExec(nums, l+1, r)
		// 由于切片是引用传递(浅拷贝)，所以这里必须再交换回来，这样才能保证不会影响上层。
		// 如果每次传给下层的切片都是新切片(深拷贝)，那么就不需要交换，因为这样不会影响上层。
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
	题目链接: https://leetcode-cn.com/problems/permutations-ii/submissions/
*/

/*
	总结
	1. 这题目的本质是: 给定一个数组，求数组的全排列。 (数组可以有重复元素)
*/