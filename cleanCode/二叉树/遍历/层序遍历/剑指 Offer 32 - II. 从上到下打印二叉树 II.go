package 层序遍历

import "bytes"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	sequences := make([][]int, 0)
	queue := NewTreeNodeQueue()
	queue.Push(root)
	for queue.IsNotEmpty() {
		sequenceOfTCurrentLay := make([]int, 0)
		for countOfUnRecordNodes := queue.GetSize(); countOfUnRecordNodes > 0; countOfUnRecordNodes-- {
			top := queue.Pop()
			if top == nil {
				continue
			}
			sequenceOfTCurrentLay = append(sequenceOfTCurrentLay, top.Val)
			queue.Push(top.Left, top.Right)
		}
		if len(sequenceOfTCurrentLay) != 0 {
			sequences = append(sequences, sequenceOfTCurrentLay)
		}
	}
	return sequences
}

// --------------------------- TreeNodeQueue ---------------------------
type TreeNodeQueue struct {
	data []*TreeNode
}

func NewTreeNodeQueue() *TreeNodeQueue {
	return &TreeNodeQueue{}
}

func (tnq *TreeNodeQueue) Push(treeNodes ...*TreeNode) {
	tnq.data = append(tnq.data, treeNodes...)
}

func (tnq *TreeNodeQueue) Pop() *TreeNode {
	top := tnq.data[0]
	tnq.data = tnq.data[1:]
	return top
}

func (tnq *TreeNodeQueue) IsNotEmpty() bool {
	return tnq.GetSize() != 0
}

func (tnq *TreeNodeQueue) GetSize() int {
	return len(tnq.data)
}

/*
	题目链接: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/submissions/
	总结:
		1. 这题和 _102. 二叉树的层序遍历_ 一样。
*/
