package 层序遍历

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := NewMyQueue()
	queue.Push(root)
	nums := make([]int, 0)
	for !queue.IsEmpty() {
		size := queue.GetSize()
		for i := 0; i < size; i++ {
			top := queue.Pop()
			nums = append(nums, top.Val)
			if top.Left != nil {
				queue.Push(top.Left)
			}
			if top.Right != nil {
				queue.Push(top.Right)
			}
		}
	}
	return nums
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
	题目链接: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
	总结:
		1. 可以说是层序遍历的裸题了。
*/
