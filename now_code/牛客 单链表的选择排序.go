package main
import (
	. "nc_tools"
)

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类 the head node
 * @return ListNode类
 */

func sortInList( head *ListNode ) *ListNode {
	for head != nil{
		minValNode :=getMinValNode(head)
		head.Val,minValNode.Val = minValNode.Val,head.Val
		head = head.Next
	}
	return head
}

func getMinValNode( head *ListNode) *ListNode{
	minValNode := head
	for head!=nil{
		if minValNode.Val > head.Val{
			minValNode = head
		}
		head = head.Next
	}
	return  minValNode
}

/*
	题目链接：https://www.nowcoder.com/practice/f23604257af94d939848729b1a5cda08?tpId=188&&tqId=36172&rp=1&ru=/ta/job-code-high-week&qru=/ta/job-code-high-week/question-ranking
	总结:
		1. 这题居然超时...已经按照要求写的了。
*/