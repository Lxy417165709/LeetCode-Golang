package 链表

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(headOfOldList *Node) *Node {
	oldNodeToNewNode := getNodeRelationMap(headOfOldList)
	return getNewList(headOfOldList, oldNodeToNewNode)
}

func getNewList(headOfOldList *Node, oldNodeToNewNode map[*Node]*Node) *Node {
	headOfNewList := oldNodeToNewNode[headOfOldList]
	curNew := headOfNewList
	curOld := headOfOldList
	for curNew != nil && curOld != nil {
		curNew.Next = oldNodeToNewNode[curOld.Next]
		curNew.Random = oldNodeToNewNode[curOld.Random]
		curNew.Val = curOld.Val

		curNew = curNew.Next
		curOld = curOld.Next
	}
	return headOfNewList
}

func getNodeRelationMap(head *Node) map[*Node]*Node {
	curOld := head
	oldNodeToNewNode := make(map[*Node]*Node)
	for curOld != nil {
		oldNodeToNewNode[curOld] = &Node{}
		curOld = curOld.Next
	}
	return oldNodeToNewNode
}

/*
	题目链接: https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
	总结
		1. 上面是采用哈希表的方法AC的代码，官方题解还有更好的算法！它的思想是在原链表的基础上，
			记录新链表的信息，从而省去了哈希表，将空间复杂度从 O(n) 优化为 O(1)。 ---> 这种方法比较Geek~
		2. 这题和 _138. 复制带随机指针的链表_ 是一样的。
*/
