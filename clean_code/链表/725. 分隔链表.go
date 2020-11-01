package 链表

func splitListToParts(root *ListNode, k int) []*ListNode {
	length := getLength(root)
	cutSize, countOfLengthGreaterOneList := length/k, length%k
	countOfLengthEqualList := k - countOfLengthGreaterOneList

	graterOneLists, curHead := repeatCut(root, cutSize+1, countOfLengthGreaterOneList)
	equalLists, _ := repeatCut(curHead, cutSize, countOfLengthEqualList)

	result := append(graterOneLists, equalLists...)
	return result
}

func getLength(list *ListNode) int {
	cur := list
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	return length
}

func repeatCut(head *ListNode, size int, cutTimes int) ([]*ListNode, *ListNode) {
	var curHead, nextHead *ListNode
	var result []*ListNode
	curHead, nextHead = head, head
	for i := 1; i <= cutTimes; i++ {
		if size == 0 {
			result = append(result, nil)
			continue
		}
		nextHead = cut(curHead, size)
		result = append(result, curHead)
		curHead = nextHead
	}
	return result, nextHead
}

func cut(head *ListNode, size int) *ListNode {
	curtail := getKthNode(head, size)
	nextHead := curtail.Next
	curtail.Next = nil
	return nextHead
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

/*
	题目链接: https://leetcode-cn.com/problems/split-linked-list-in-parts/submissions/
*/
