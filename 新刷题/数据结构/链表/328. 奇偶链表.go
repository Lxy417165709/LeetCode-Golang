package 链表

// oddEvenList 重排链表为 L0 → L2 → L4  → ... → L(2n-5) → L(2n-3) → L(2n-1) 顺序。
func oddEvenList(head *ListNode) *ListNode {
	// 1. 空返回。
	if head == nil {
		return nil
	}

	// 2. 分离链表。
	oddListDummyHead, evenListDummyHead := spiltListByPos(head)

	// 3. 连接奇偶链表。
	getTailNode(oddListDummyHead).Next = evenListDummyHead.Next

	// 4. 返回。
	return oddListDummyHead.Next
}

// spiltListByPos 根据序号的奇偶性拆分链表，返回具有伪头的奇偶序号链表。
func spiltListByPos(head *ListNode) (*ListNode, *ListNode) {
	// 1. 初始化。
	oddListDummyHead := &ListNode{}
	evenListDummyHead := &ListNode{}

	// 2. 根据序号的奇偶性拆分链表。
	cur := head
	curOfOddList := oddListDummyHead
	curOfEvenList := evenListDummyHead
	for curPos := 1; cur != nil; curPos++ {
		if curPos&1 == 1 {
			curOfOddList.Next = cur
			curOfOddList = curOfOddList.Next
		} else {
			curOfEvenList.Next = cur
			curOfEvenList = curOfEvenList.Next
		}
		cur = cur.Next
	}

	// 3. 断开连接。
	curOfOddList.Next = nil
	curOfEvenList.Next = nil

	// 4. 返回。
	return oddListDummyHead, evenListDummyHead
}

// getTailNode 获取尾节点。
func getTailNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}
