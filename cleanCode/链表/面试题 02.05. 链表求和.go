package main

type ListNode struct {
	Val  int
	Next *ListNode
}
// ---------------------- 方法1: 时空复杂度 O(n),O(1) ----------------------
// 概述: 将链表相加的结果存到较长的链表中。
func addTwoNumbers(listA *ListNode, listB *ListNode) *ListNode {
	listA, listB = getLongAndShortList(listA, listB)
	curA, curB := listA, listB
	for curA != nil {
		valA, valB := getValue(curA), getValue(curB)
		setValue(curA, valA+valB)
		curA, curB = getNext(curA), getNext(curB)
	}
	return listA
}

func getLongAndShortList(listA, listB *ListNode) (*ListNode, *ListNode) {
	if getLength(listA) < getLength(listB) {
		listA, listB = listB, listA
	}
	return listA, listB
}

func getValue(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}

func setValue(list *ListNode, value int) {
	standard, carry := getStandardAndCarry(value)
	list.Val = standard
	if carry != 0 {
		if list.Next == nil {
			list.Next = &ListNode{}
		}
		list.Next.Val += carry
	}
}

func getNext(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return nil
	}
	return node.Next
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

func getStandardAndCarry(sum int) (int, int) {
	standard, carry := sum%10, sum/10
	return standard, carry
}


/*
	题目链接: https://leetcode-cn.com/problemset/all/
*/
