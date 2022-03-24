package 链表

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 1. 初始化。
	carry := 0
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}

	// 2. 计算。
	for resultNode := dummyHead; l1 != nil || l2 != nil || carry != 0; {
		// 2.1 获取总和。
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// 2.2 保存相加结果。
		resultNode.Next = &ListNode{
			Val: sum % 10,
		}

		// 2.3 进入下次循环。
		resultNode = resultNode.Next
		carry = sum / 10
	}

	// 3. 返回。
	return dummyHead.Next
}
