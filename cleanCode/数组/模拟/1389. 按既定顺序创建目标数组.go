package 模拟

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(nums))
	for i := 0; i < len(index); i++ {
		target = insert(target, index[i], nums[i])
	}
	return target
}

func insert(array []int, index, num int) []int {
	for t := len(array) - 1; t >= index+1; t-- {
		array[t] = array[t-1]
	}
	array[index] = num
	return array
}

/*
	题目链接: https://leetcode-cn.com/problems/create-target-array-in-the-given-order/
	总结:
		1. 这题类似插入排序。
*/
