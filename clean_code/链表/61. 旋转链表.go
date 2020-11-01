package 链表

type ListNode struct {
	Val  int
	Next *ListNode
}
// --------------------- 方法1: 时空复杂度 O(n),O(1) ---------------------
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := getLength(head)
	listPart1, listPart2 := cutToTwoParts(head, k%length)
	return link(listPart2, listPart1)
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

func cutToTwoParts(head *ListNode, lengthOfSecondPart int) (*ListNode, *ListNode) {
	// lengthOfSecondPart 必须小于链表长度。
	node := getLastKthNode(head, lengthOfSecondPart+1)
	firstPart, secondPart := head, node.Next
	node.Next = nil
	return firstPart, secondPart
}

func link(listA *ListNode, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	}
	tailNodeOfListA := getTailNode(listA)
	tailNodeOfListA.Next = listB
	return listA
}

func getLastKthNode(head *ListNode, k int) *ListNode {
	aft := getKthNode(head, k)
	cur := head
	for aft.Next != nil {
		cur = cur.Next
		aft = aft.Next
	}
	return cur
}

func getTailNode(notNilList *ListNode) *ListNode {
	cur := notNilList
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
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
	题目链接: https://leetcode-cn.com/problems/rotate-list/
*/
