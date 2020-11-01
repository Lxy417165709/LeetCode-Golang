package 链表

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

/*
	题目链接: https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
	总结：
		1. 上面的代码可以删除链表中多个值为val的节点。
		2. 这题和 _203. 移除链表元素_ 差不多。
*/
