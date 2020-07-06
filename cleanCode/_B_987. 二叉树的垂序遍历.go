package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var intermediateResult [][]int
const beginSpecialValue = 1000
const maxSpecialValue = beginSpecialValue + 1000
const minSpecialValue = beginSpecialValue - 1000
const beginLay = 0
const weightOfLay = 10000

func verticalTraversal(root *TreeNode) [][]int {
	intermediateResult = make([][]int, maxSpecialValue-minSpecialValue+1)
	formIntermediateResult(root, beginSpecialValue, beginLay)
	result := make([][]int, 0)
	for i := 0; i < len(intermediateResult); i++ {
		if len(intermediateResult[i]) != 0 {
			sort.Ints(intermediateResult[i])
			balanceAllNumsInArray(intermediateResult[i], weightOfLay)
			result = append(result, intermediateResult[i])
		}
	}
	return result
}

func formIntermediateResult(root *TreeNode, nowSpecialValue int, lay int) {
	if root == nil {
		return
	}
	intermediateResult[nowSpecialValue] = append(intermediateResult[nowSpecialValue], lay*weightOfLay+root.Val)
	formIntermediateResult(root.Left, nowSpecialValue-1, lay+1)
	formIntermediateResult(root.Right, nowSpecialValue+1, lay+1)
}

func balanceAllNumsInArray(array []int, mod int) {
	for i := 0; i < len(array); i++ {
		array[i] %= mod
	}
}



/*
	题目链接： https://leetcode-cn.com/problems/vertical-order-traversal-of-a-binary-tree/
	总结
		1. 由于要求层数低的先输出，所以使用递归的遍历时，麻烦之处在于：如何记录节点层的信息。
			这题采用了比较geek的技巧，就是将层的信息*10000，记录到 intermediateResult 中。
		2. 题解还有人使用层序遍历AC的，当时WA后也想到了层序遍历。
*/
