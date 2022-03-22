package 链表

func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 1. 空返回。
	if head == nil {
		return nil
	}

	// 2. 获取链表长度。
	length := getListLen(head)

	// 3. 消除循环节。
	k = (k-1)%length + 1

	// 4. 先行指针走 k-1 步。
	goFirstPointer := head
	for i := 1; i <= k-1; i++ {
		goFirstPointer = goFirstPointer.Next
	}

	// 5. 先行指针、后行指针一起走，直到先行指针到达链表尾。
	goSecondPointer := head
	for goFirstPointer.Next != nil {
		goFirstPointer = goFirstPointer.Next
		goSecondPointer = goSecondPointer.Next
	}

	// 6. 此时后行指针指向链表倒数第 k 个节点。
	return goSecondPointer
}

// getListLen 获取链表长度。
func getListLen(head *ListNode) int {
	length := 0
	for cur := head; cur != nil; {
		length++
		cur = cur.Next
	}
	return length
}

// 问题: 为什么走 k-1 步?
// 证明: 假如先行指针位置为 b，后行指针位置为 a。
// 		当 b 是倒数第 1 个节点时，此时 a 为倒数第 (b-a)+1 个节点。
//		为了保证 a 是倒数第 k 个节点，此时需要保证 (b-a)+1 = k，即 b-a = k-1，而 b-a 即为 先行指针与后行指针的距离。
//		所以 先行指针需要先走 k - 1 步。