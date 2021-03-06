package 链表


func hasCycle(head *ListNode) bool {
	slow,fast := head,head
	for fast!=nil && fast.Next!=nil{
		slow = slow.Next
		fast=fast.Next.Next
		if slow==fast{
			return true
		}
	}
	return false
}
/*
	题目链接: https://leetcode-cn.com/problems/linked-list-cycle/comments/
*/
