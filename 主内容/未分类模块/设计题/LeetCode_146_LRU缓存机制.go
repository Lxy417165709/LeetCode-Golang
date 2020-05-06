package main

type Node struct {
	Key       int
	Val       int
	Pre, Next *Node
}

func NewNode(key, value int) *Node {
	return &Node{key, value, nil, nil}
}

type DoubleList struct {
	dummyHead, dummyTail *Node
}

func NewDoubleList() *DoubleList {
	dummyHead := new(Node)
	dummyTail := new(Node)
	dummyHead.Next = dummyTail
	dummyTail.Pre = dummyHead
	return &DoubleList{dummyHead, dummyTail}
}

// supposed that node is not nil
func (d *DoubleList) InsertToTail(node *Node) {
	d.dummyTail.Pre.Next = node
	node.Pre = d.dummyTail.Pre
	d.dummyTail.Pre = node
	node.Next = d.dummyTail
}
func (d *DoubleList) InsertToHead(node *Node) {
	d.dummyHead.Next.Pre = node
	node.Next = d.dummyHead.Next
	d.dummyHead.Next = node
	node.Pre = d.dummyHead
}
func (d *DoubleList) Remove(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre

	// To help GC
	node.Pre = nil
	node.Next = nil
}
func (d *DoubleList) GetTailNode() *Node {
	return d.dummyTail.Pre
}

type LRUCache struct {
	doubleList *DoubleList
	keyToNode  map[int]*Node
	length     int
	capacity   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{NewDoubleList(), make(map[int]*Node), 0, capacity}
}

func (lru *LRUCache) Get(key int) int {
	if !lru.isLiving(key) {
		// the job requires that when the key is not exist we should return -1
		return -1
	}
	node := lru.getNode(key)
	lru.flash(key, node.Val)
	return node.Val
}

func (lru *LRUCache) Put(key int, value int) {
	if lru.isLiving(key) {
		lru.flash(key, value)
		return
	}
	if lru.isFull() {
		lru.removeOldestKey()
	}
	lru.insertNewKey(key, value)
}

func (lru *LRUCache) isLiving(key int) bool {
	return lru.keyToNode[key] != nil
}
func (lru *LRUCache) isFull() bool {
	return lru.length == lru.capacity
}

func (lru *LRUCache) getNode(key int) *Node {
	return lru.keyToNode[key]
}

// replace the key's old value with newValue
// and adjust the position of key in doubleList
func (lru *LRUCache) flash(key int, newValue int) {
	node := lru.getNode(key)
	node.Val = newValue
	lru.doubleList.Remove(node)
	lru.doubleList.InsertToHead(node)
}

func (lru *LRUCache) removeOldestKey() {
	lastNode := lru.doubleList.GetTailNode()
	lru.keyToNode[lastNode.Key] = nil
	lru.doubleList.Remove(lastNode)
	lru.length--
}
func (lru *LRUCache) insertNewKey(key int, value int) {
	newNode := NewNode(key, value)
	lru.keyToNode[key] = newNode
	lru.doubleList.InsertToHead(newNode)
	lru.length++
}

func main() {

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
