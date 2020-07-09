package 链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	leftMiddleNode := getLeftMiddleNode(head)
	rightPartList := getReverseList(leftMiddleNode.Next)
	return isPrefix(head, rightPartList)
}

func getLeftMiddleNode(list *ListNode) *ListNode {
	slow, fast := list, list.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func getReverseList(list *ListNode) *ListNode {
	// 这个函数会改变 list 的内部结构
	var next, pre *ListNode
	cur := list
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func isPrefix(patternList *ListNode, sampleList *ListNode) bool {
	pCur := patternList
	sCur := sampleList
	for pCur != nil && sCur != nil {
		if pCur.Val != sCur.Val {
			return false
		}
		pCur = pCur.Next
		sCur = sCur.Next
	}
	return sCur == nil
}

/*
	题目链接: https://leetcode-cn.com/problems/palindrome-linked-list/
	总结:
		1. 上面的这种方法虽然空间复杂度是O(1)，但是有改变链表结构的副作用。
		2. 可以在判断之后恢复链表，这样就不会有副作用了。
*/
