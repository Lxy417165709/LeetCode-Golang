package 链表

func reversePrint(head *ListNode) []int {
	return getReverseVals(head)
}

func getReverseVals(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(getReverseVals(head.Next), head.Val)
}

/*
	题目链接:
	总结:
		1. 有多种方法可以AC这题，比如先反转再记录；比如先开数组，再把元素放进去后，反转数组....
			比如使用栈...但我就是要用递归！
*/
