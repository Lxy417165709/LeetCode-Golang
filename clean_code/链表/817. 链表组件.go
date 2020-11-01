package 链表

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// ----------------------方法1: 时空复杂度 O(n^2),O(1) -----------------
// 概述: 暴力查找。
func numComponents(head *ListNode, G []int) int {
	cur := head
	countOfComponents := 0
	curLength := 0
	for cur != nil {
		if isInArray(G, cur.Val) {
			if curLength == 0 {
				countOfComponents++
			}
			curLength++
		} else {
			curLength = 0
		}
		cur = cur.Next
	}
	return countOfComponents
}

func isInArray(array []int, val int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			return true
		}
	}
	return false
}

// ----------------------方法2: 时空复杂度 O(n * log_n),O(1) -----------------
// 概述: 二分查找。
func numComponents(head *ListNode, G []int) int {
	sort.Ints(G)
	cur := head
	countOfComponents := 0
	curLength := 0
	for cur != nil {
		if isInArray(G, cur.Val) {
			if curLength == 0 {
				countOfComponents++
			}
			curLength++
		} else {
			curLength = 0
		}
		cur = cur.Next
	}
	return countOfComponents
}

func isInArray(sortedArray []int, ref int) bool {
	left, right := 0, len(sortedArray)-1
	for left <= right {
		mid := (left + right) / 2
		if sortedArray[mid] == ref {
			return true
		}
		if sortedArray[mid] > ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// ----------------------方法3: 时空复杂度 O(n),O(n) -----------------
// 概述: 通过哈希表查找。
var isInG map[int]bool

func numComponents(head *ListNode, G []int) int {
	isInG = make(map[int]bool)
	for i := 0; i < len(G); i++ {
		isInG[G[i]] = true
	}
	cur := head
	countOfComponents := 0
	curLength := 0
	for cur != nil {
		if isInG[cur.Val] {
			if curLength == 0 {
				countOfComponents++
			}
			curLength++
		} else {
			curLength = 0
		}
		cur = cur.Next
	}
	return countOfComponents
}

/*
	题目链接: https://leetcode-cn.com/problems/linked-list-components/comments/
*/
