package main

import (
	"fmt"
	"strings"
)

type MyDummyDbList struct {
	dummyHead, dummyTail *MyDummyDbListNode
}

func NewMyDummyDbList() *MyDummyDbList {
	dummyHead, dummyTail := &MyDummyDbListNode{}, &MyDummyDbListNode{}
	dummyHead.Next = dummyTail
	dummyTail.Pre = dummyHead
	return &MyDummyDbList{
		dummyHead: dummyHead,
		dummyTail: dummyTail,
	}
}

func (m *MyDummyDbList) RemoveNode(node *MyDummyDbListNode) {
	node.Next.Pre = node.Pre
	node.Pre.Next = node.Next

	node.Next = nil
	node.Pre = nil
}

func (m *MyDummyDbList) InsertNode(preNode *MyDummyDbListNode, node *MyDummyDbListNode) {
	preNode.Next.Pre = node
	node.Next = preNode.Next
	preNode.Next = node
	node.Pre = preNode
}

type MyDummyDbListNode struct {
	Key       int
	Value     int
	Freq      int
	Pre, Next *MyDummyDbListNode
}

type MyLFUCache struct {
	FreqToNodeMap map[int]*MyDummyDbListNode
	KeyToNodeMap  map[int]*MyDummyDbListNode
	DummyDbList   *MyDummyDbList
	Capacity      int
}

func NewMyLFUCache(capacity int) *MyLFUCache {
	return &MyLFUCache{
		FreqToNodeMap: make(map[int]*MyDummyDbListNode),
		KeyToNodeMap:  make(map[int]*MyDummyDbListNode),
		DummyDbList:   NewMyDummyDbList(),
		Capacity:      capacity,
	}
}

func (m *MyLFUCache) Put(key int, value int) {
	// 1. 键存在、空间未满时。
	_, exist := m.KeyToNodeMap[key]
	if exist || len(m.KeyToNodeMap) < m.Capacity {
		m.Adjust(key, value)
		return
	}

	// 2. 键不存在且空间满时。(题目限制了容量不可能为0，所以删除的一定是内部元素)
	m.RemoveLFUNode()
	m.Adjust(key, value)
}

func (m *MyLFUCache) RemoveLFUNode() {
	removeNode := m.DummyDbList.dummyHead.Next
	if m.FreqToNodeMap[removeNode.Freq] == removeNode {
		delete(m.FreqToNodeMap, removeNode.Freq)
	}
	delete(m.KeyToNodeMap, removeNode.Key)
	m.DummyDbList.RemoveNode(m.DummyDbList.dummyHead.Next)
}

func (m *MyLFUCache) Get(key int) int {
	// 1. 键不存在时。
	node, exist := m.KeyToNodeMap[key]
	if !exist {
		return -1
	}

	// 2. 调整。
	m.Adjust(node.Key, node.Value)

	// 3. 返回。
	return node.Value
}

func (m *MyLFUCache) Adjust(key, value int) {
	if _, exist := m.KeyToNodeMap[key]; exist {
		m.AdjustWhenKeyExist(key, value)
	} else {
		m.AdjustWhenKeyNotExist(key, value, 1)
	}
}

func (m *MyLFUCache) AdjustWhenKeyExist(key, value int) {
	// 1. 将对应的 Key 删除。
	node := m.KeyToNodeMap[key]
	curFreq := node.Freq
	if node.Pre.Freq == curFreq {
		m.FreqToNodeMap[curFreq] = node.Pre
	} else {
		delete(m.FreqToNodeMap, curFreq)
	}
	delete(m.KeyToNodeMap, key)
	m.DummyDbList.RemoveNode(node)

	// 2. 调整。
	nextFreq := curFreq + 1
	m.AdjustWhenKeyNotExist(key, value, nextFreq)
}

func (m *MyLFUCache) AdjustWhenKeyNotExist(key, value, freq int) {
	// 1. 初始化。
	nextFreq := freq
	newNode := &MyDummyDbListNode{
		Key:   key,
		Value: value,
		Freq:  nextFreq,
	}

	// 2. 处理。
	if nextFreqNode, exist := m.FreqToNodeMap[nextFreq]; exist {
		// 2.1 如果 nextFreq 存在，将新节点插入对应的节点尾部，调整 FreqToNodeMap。
		m.DummyDbList.InsertNode(nextFreqNode, newNode)
		m.FreqToNodeMap[nextFreq] = newNode
	} else {
		// 2.2 如果 nextFreq 不存在，将新节点插入双向链表头部，调整 FreqToNodeMap。
		m.DummyDbList.InsertNode(m.DummyDbList.dummyHead, newNode)
		m.FreqToNodeMap[nextFreq] = newNode

		// todo: 这里有问题， 当链表元素频次是  1->1->1 时，get第一个元素后，实际上是要插到链表尾部，变为 1->1->2
	}
	m.KeyToNodeMap[key] = newNode
}

func main() {
	lfu := NewMyLFUCache(3)
	lfu.Put(1, 10)
	lfu.Put(2, 20)
	lfu.Put(3, 30)

	showList(lfu.DummyDbList.dummyHead)
	fmt.Println("get 1", lfu.Get(1))
	showList(lfu.DummyDbList.dummyHead)

	lfu.Put(4, 40)
	showList(lfu.DummyDbList.dummyHead)
	fmt.Println(lfu.Get(2))
	//fmt.Printf("%+v %+v",lfu.FreqToNodeMap[1],lfu.FreqToNodeMap[2])
}

func showList(listNode *MyDummyDbListNode) {
	values := make([]string, 0)
	cur := listNode
	for cur != nil {
		values = append(values,fmt.Sprintf("%d",cur.Value))
		cur = cur.Next
	}
	fmt.Println(strings.Join(values[1:len(values)-1],"->"))
}

// 这题还没解决