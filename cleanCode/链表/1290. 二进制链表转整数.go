package 链表

func getDecimalValue(head *ListNode) int {
	cur := head
	value := 0
	for cur != nil {
		value = value*2 + cur.Val
		cur = cur.Next
	}
	return value
}

/*
	题目链接: https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/
	总结:
		1. "水"只一个字，我只说一次。
*/
