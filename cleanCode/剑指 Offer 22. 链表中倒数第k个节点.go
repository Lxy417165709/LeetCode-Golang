package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getKthFromEnd(head *ListNode, k int) *ListNode {
	nodeStack := NewNodeStack()
	nodeStack.PushAllListNodes(head)
	nodeStack.PopNTimes(k - 1)
	return nodeStack.GetTop()
}

type NodeStack struct {
	nodes []*ListNode
}

func NewNodeStack() *NodeStack {
	return &NodeStack{}
}

func (ns *NodeStack) PushAllListNodes(head *ListNode) {
	for head != nil {
		ns.push(head)
		head = head.Next
	}
}

func (ns *NodeStack) PopNTimes(n int) {
	for n > 0 {
		ns.pop()
		n--
	}
}

func (ns *NodeStack) GetTop() *ListNode {
	return ns.nodes[ns.getSize()-1]
}

func (ns *NodeStack) push(node *ListNode) {
	ns.nodes = append(ns.nodes, node)
}

func (ns *NodeStack) pop() {
	ns.nodes = ns.nodes[:ns.getSize()-1]
}

func (ns *NodeStack) getSize() int {
	return len(ns.nodes)
}

/*
	题目链接: https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

	解法：
		O(1)空间复杂度:
			1次扫描: 采用前后指针
			2次扫描: 第一次先获取长度，第二次就能到达准确的位置
		O(n)空间复杂度:
			1次扫描: 采用数组存储
			2次扫描: 采用栈
*/
