package 遍历

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	rightTreeLeftIndex, rightTreeRightIndex := getIndexOfFirstNumGreaterThanLastNum(postorder), len(postorder)-2
	leftTreeLeftIndex, leftTreeRightIndex := 0, rightTreeLeftIndex-1
	rootVal := postorder[len(postorder)-1]
	if isExistGreaterNum(postorder[leftTreeLeftIndex:leftTreeRightIndex+1], rootVal) {
		return false
	}
	return verifyPostorder(postorder[leftTreeLeftIndex:leftTreeRightIndex+1]) &&
		verifyPostorder(postorder[rightTreeLeftIndex:rightTreeRightIndex+1])
}

func isExistGreaterNum(arr []int, ref int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] > ref {
			return true
		}
	}
	return false
}

func getIndexOfFirstNumGreaterThanLastNum(arr []int) int {
	lastNum := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] < lastNum {
			return i + 1
		}
	}
	return 0
}

/*
	题目链接: https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
	总结:
		1. 官方还有人用单调栈做这题的！
*/
