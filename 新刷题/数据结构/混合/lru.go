package 混合

import (
	"fmt"
	"strings"
)

// --------------------------------------------------  双向链表(开始) --------------------------------------------------

type MyDbList struct {
	Head, Tail *MyDbListNode
}

// Append 追加节点。
func (l *MyDbList) Append(node *MyDbListNode) *MyDbListNode {
	// 1. 清空节点原连接。
	node.Next = nil
	node.Pre = nil

	// 2. 链表为空。
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return node
	}

	// 3. 链表非空。
	l.Tail.Next = node
	node.Pre = l.Tail
	l.Tail = l.Tail.Next

	// 4. 返回。
	return node
}

// Remove 删除节点。
func (l *MyDbList) Remove(node *MyDbListNode) *MyDbListNode {
	// 1. 节点为空。
	if node == nil {
		return nil
	}

	// 2. 节点为头节点，
	if node == l.Head {
		// 2.1 只有一个节点。
		if node.Next == nil {
			l.Head = nil
			l.Tail = nil
			return node
		}

		// 2.2 多个节点。
		l.Head.Next.Pre = nil
		l.Head = l.Head.Next
		return node
	}

	// 3. 节点为尾节点。 (此时长度一定大于1)
	if node == l.Tail {
		l.Tail.Pre.Next = nil
		l.Tail = l.Tail.Pre
		return node
	}

	// 4. 节点为中间节点。
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	return node
}

// RangeNodes 遍历出所有节点。
func (l *MyDbList) RangeNodes() []*MyDbListNode {
	nodes := make([]*MyDbListNode, 0)
	cur := l.Head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}
	return nodes
}

// Stringify 序列化链表。
func (l *MyDbList) Stringify() string {
	nodes := l.RangeNodes()
	parts := make([]string, 0)
	for _, node := range nodes {
		parts = append(parts, fmt.Sprintf("(key:%+v, value:%+v)", node.Key, node.Value))
	}
	return strings.Join(parts, "->")
}

type MyDbListNode struct {
	Key       int
	Value     int
	Pre, Next *MyDbListNode
}

// -------------------------------------------------- 双向链表(结束) --------------------------------------------------

// --------------------------------------------------  LRU(开始) --------------------------------------------------
type MyLRU struct {
	KeyToNodeMap map[int]*MyDbListNode
	DbList       *MyDbList
	Capacity     int
}

func NewMyLRU(capacity int) *MyLRU {
	return &MyLRU{
		KeyToNodeMap: make(map[int]*MyDbListNode),
		DbList:       &MyDbList{},
		Capacity:     capacity,
	}
}

func (l *MyLRU) Put(key, value int) {
	// 1. 键已存在，更新节点的值。
	node := l.KeyToNodeMap[key]
	if node != nil {
		node.Value = value
		l.flashNode(node)
		return
	}

	// 2. 容量是否已满，满了则清理最近最少的键。
	if len(l.KeyToNodeMap) == l.Capacity {
		l.removeLruKey()
	}

	// 3. 加入新键。
	l.insertKey(key, value)
}

func (l *MyLRU) Get(key int) int {
	// 1. 键不存在，返回。
	node := l.KeyToNodeMap[key]
	if node == nil {
		return -1
	}

	// 2. 键已存在，更新节点使用序。
	l.flashNode(node)

	// 3. 返回。
	return node.Value
}

func (l *MyLRU) removeLruKey() {
	removedNode := l.DbList.Remove(l.DbList.Head)
	if removedNode != nil {
		delete(l.KeyToNodeMap, removedNode.Key)
	}
}

func (l *MyLRU) insertKey(key, value int) {
	l.insertNode(&MyDbListNode{
		Key:   key,
		Value: value,
	})
}

func (l *MyLRU) insertNode(node *MyDbListNode) {
	l.KeyToNodeMap[node.Key] = l.DbList.Append(node)
}

func (l *MyLRU) flashNode(node *MyDbListNode) {
	l.DbList.Remove(node)
	l.insertNode(node)
}

// --------------------------------------------------  LRU(结束) --------------------------------------------------

// 优化: 其实双向链表可以搞个伪头、伪尾，这样便于编码。
