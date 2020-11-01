package 链表

func isPalindrome(head *ListNode) bool {
	if head==nil{
		return true
	}
	leftMiddleNode := getLeftMiddleNode(head)
	rightPartList := reverseList(leftMiddleNode.Next)
	return isPrefix(head,rightPartList)
}

func isPrefix(patternList *ListNode,sampleList*ListNode) bool{
	pCur := patternList
	sCur := sampleList
	for pCur!=nil && sCur!=nil{
		if pCur.Val!=sCur.Val{
			return false
		}
		pCur=pCur.Next
		sCur=sCur.Next
	}
	return sCur==nil
}

func reverseList(list *ListNode) *ListNode{
	var next,pre *ListNode
	cur := list
	for cur!=nil{
		next=cur.Next
		cur.Next=pre
		pre=cur
		cur = next

	}
	return pre
}

func getLeftMiddleNode(list *ListNode) *ListNode{
	slow,fast := list,list.Next
	for fast!=nil && fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
	}
	return slow
}

/*
	题目链接: https://leetcode-cn.com/problems/palindrome-linked-list-lcci/comments/
	总结:
		1. 这题和 _234. 回文链表_ 是一样的。
*/
