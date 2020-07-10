package 链表




func reverseKGroup(head *ListNode, k int) *ListNode {

}

func reverse(head *ListNode) (*ListNode,*ListNode) {
	var pre, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre,head
}

func cut(head *ListNode, size int) (*ListNode,*ListNode) {
	curTail := getKthNode(head, size)
	if curTail==nil{
		return nil,nil
	}
	nextHead := curTail.Next
	curTail.Next = nil
	return curTail,nextHead
}

func getKthNode(head *ListNode, k int) *ListNode {
	cur := head
	moveTimes := k - 1
	for cur != nil && moveTimes > 0 {
		moveTimes--
		cur = cur.Next
	}
	return cur
}
