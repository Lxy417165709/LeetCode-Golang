package 层序遍历

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	queue := NewMyQueue()
	queue.Push(root)
	averages := make([]float64, 0)
	for !queue.IsEmpty() {
		size := queue.GetSize()
		sum, countOfNode := 0, 0
		for i := 0; i < size; i++ {
			top := queue.Pop()
			sum += top.Val
			countOfNode++
			if top.Left != nil {
				queue.Push(top.Left)
			}
			if top.Right != nil {
				queue.Push(top.Right)
			}
		}
		averages = append(averages, float64(sum)/float64(countOfNode))
	}
	return averages
}

// ------------------ MyQueue ---------------------
type MyQueue struct {
	data []*TreeNode
}

func NewMyQueue() *MyQueue {
	return &MyQueue{}
}

func (mq *MyQueue) Push(val *TreeNode) {
	mq.data = append(mq.data, val)
}

func (mq *MyQueue) IsEmpty() bool {
	return mq.GetSize() == 0
}

func (mq *MyQueue) GetSize() int {
	return len(mq.data)
}

func (mq *MyQueue) Pop() *TreeNode {
	top := mq.data[0]
	mq.data = mq.data[1:]
	return top
}

func (mq *MyQueue) GetTop() *TreeNode {
	return mq.data[0]
}

/*
	题目链接: https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
	总结:
		1. 可以说是层序遍历的裸题了。
*/
