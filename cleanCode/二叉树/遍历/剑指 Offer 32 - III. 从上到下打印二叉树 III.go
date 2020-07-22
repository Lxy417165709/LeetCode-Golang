package 遍历

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
	curLay := 0
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
		if len(sequenceOfTCurrentLay) == 0 {
			continue
		}
		if curLay%2 == 0 {
			sequences = append(sequences, sequenceOfTCurrentLay)
		} else {
			sequences = append(sequences, reverse(sequenceOfTCurrentLay))
		}
		curLay++
	}
	return sequences
}

func reverse(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
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
	题目链接: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
	总结:
		1. 送分题。
*/
