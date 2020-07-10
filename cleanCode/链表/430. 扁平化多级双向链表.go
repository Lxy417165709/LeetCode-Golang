package 链表

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	beList(root)
	return root
}

// 当root == nil时，返回 (nil,nil), 否则返回扁平化链表的头尾节点。
func beList(root *Node) (*Node, *Node) {
	if root == nil {
		return nil, nil
	}
	childHead, childTail := beList(root.Child)
	nextHead, nextTail := beList(root.Next)
	root.Child = nil
	switch {
	case nextHead == nil && childHead == nil:
		return root, root
	case nextHead != nil && childHead == nil:
		return root, nextTail
	case nextHead == nil && childHead != nil:
		root.Next, childHead.Prev = childHead, root
		return root, childTail
	case nextHead != nil && childHead != nil:
		root.Next, childHead.Prev = childHead, root
		childTail.Next, nextHead.Prev = nextHead, childTail
		return root, nextTail
	}
	panic("不可能到这")
}

/*
	题目链接: https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/comments/
	总结：
		1. 第一眼我就看出这其实是棵树。
		2. 这题和 _114. 二叉树展开为链表_ 很像，只是多了个 prev 指针。
*/
