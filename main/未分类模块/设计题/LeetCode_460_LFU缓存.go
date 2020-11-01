package main

import "fmt"

// -------------- LFUCache ----------------
type LFUCache struct {
	length                  int
	capacity                int
	frequencyMaintainer     *FrequencyMaintainer // 作用： 维护 frequency -> node 的映射关系
	keysMaintainer          *KeyMaintainer       // 作用： 维护 key -> node 的映射关系
	frequencyDecreasingList *DoubleList          // 作用： 存储底层数据，数据类型为 Node
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		0,
		capacity,
		NewFrequencyMaintainer(),
		NewKeyMaintainer(),
		NewDoubleList(),
	}
}

func (lfu *LFUCache) Get(key int) int {
	if !lfu.keyIsExist(key) {
		return -1
	}
	value := lfu.getValue(key)
	lfu.flashKey(key, value)
	return value
}
func (lfu *LFUCache) Put(key int, value int) {
	if !lfu.capacityIsValid() {
		return
	}
	if lfu.keyIsExist(key) {
		lfu.flashKey(key, value)
		return
	}
	if lfu.isFull() {
		lfu.deleteLastKey()
	}
	lfu.insertNewKey(key, value)
}

func (lfu *LFUCache) insertNewKey(key, value int) {
	newNode := NewNode(key, value, 0)
	lfu.maintainNodeInsertOperation(newNode)
	lfu.increaseLength()
}
func (lfu *LFUCache) deleteLastKey() {
	lastNode := lfu.frequencyDecreasingList.GetLastNode()
	lfu.maintainNodeDeleteOperation(lastNode)
	lfu.decreaseLength()
}
func (lfu *LFUCache) flashKey(key int, value int) {
	node := lfu.keysMaintainer.GetNode(key)
	nextFrequency := node.Frequency + 1
	nextNode := NewNode(key, value, nextFrequency)
	lfu.maintainNodeInsertOperation(nextNode)
	lfu.maintainNodeDeleteOperation(node)
}
func (lfu *LFUCache) getValue(key int) int {
	return lfu.keysMaintainer.GetValue(key)
}
func (lfu *LFUCache) keyIsExist(key int) bool {
	return lfu.keysMaintainer.KeyIsExist(key)
}

func (lfu *LFUCache) capacityIsValid() bool {
	return lfu.capacity > 0
}
func (lfu *LFUCache) isFull() bool {
	return lfu.length == lfu.capacity
}
func (lfu *LFUCache) increaseLength() {
	lfu.length++
}
func (lfu *LFUCache) decreaseLength() {
	lfu.length--
}

func (lfu *LFUCache) maintainNodeInsertOperation(node *Node) {
	insertReferenceNode := lfu.getReferenceNodeOfInsertOperation(node.Frequency)
	lfu.frequencyDecreasingList.InsertAfter(insertReferenceNode, node)
	lfu.frequencyMaintainer.AddFrequencyMappingRelation(node)
	lfu.keysMaintainer.AddKeyMappingRelation(node)
}
func (lfu *LFUCache) maintainNodeDeleteOperation(node *Node) {
	lfu.frequencyMaintainer.DeleteFrequencyMappingRelation(node)
	lfu.frequencyDecreasingList.DeleteNode(node)
	lfu.keysMaintainer.DeleteKeyMappingRelation(node)
}
func (lfu *LFUCache) getReferenceNodeOfInsertOperation(frequency int) *Node {
	if lfu.frequencyMaintainer.FrequencyIsExist(frequency) {
		return lfu.frequencyMaintainer.GetFirstNode(frequency).Pre
	} else {
		if frequency == 0 {
			return lfu.frequencyDecreasingList.GetLastNode()
		}
		return lfu.frequencyMaintainer.GetFirstNode(frequency - 1).Pre
	}
}

// -------------- FrequencyMaintainer ----------------
type FrequencyMaintainer struct {
	frequencyToFirstNode map[int]*Node
}

func NewFrequencyMaintainer() *FrequencyMaintainer {
	return &FrequencyMaintainer{make(map[int]*Node)}
}

func (fm *FrequencyMaintainer) AddFrequencyMappingRelation(node *Node) {
	if fm.FrequencyIsExist(node.Frequency) {
		if !fm.frequencyIsUnique(node.Frequency) {
			delete(fm.frequencyToFirstNode, node.Frequency)
		}
	}
	fm.frequencyToFirstNode[node.Frequency] = node
}
func (fm *FrequencyMaintainer) DeleteFrequencyMappingRelation(node *Node) {
	if !node.IsFirstNode() {
		return
	}
	if fm.frequencyIsUnique(node.Frequency) {
		fm.frequencyToFirstNode[node.Frequency] = node.Next
	} else {
		delete(fm.frequencyToFirstNode, node.Frequency)
	}
}

func (fm *FrequencyMaintainer) FrequencyIsExist(frequency int) bool {
	return fm.frequencyToFirstNode[frequency] != nil
}
func (fm *FrequencyMaintainer) GetFirstNode(frequency int) *Node {
	return fm.frequencyToFirstNode[frequency]
}

func (fm *FrequencyMaintainer) frequencyIsUnique(frequency int) bool {
	return fm.frequencyToFirstNode[frequency].Next.Frequency == frequency
}

// -------------- KeyMaintainer ----------------
type KeyMaintainer struct {
	keyToNode map[int]*Node
}

func NewKeyMaintainer() *KeyMaintainer {
	return &KeyMaintainer{make(map[int]*Node)}
}

func (km *KeyMaintainer) GetValue(key int) int {
	return km.GetNode(key).Value
}
func (km *KeyMaintainer) KeyIsExist(key int) bool {
	return km.GetNode(key) != nil
}
func (km *KeyMaintainer) GetNode(key int) *Node {
	return km.keyToNode[key]
}
func (km *KeyMaintainer) AddKeyMappingRelation(node *Node) {
	km.keyToNode[node.Key] = node
}
func (km *KeyMaintainer) DeleteKeyMappingRelation(node *Node) {
	if !km.keyMappingRelationIsNeedToDelete(node) {
		return
	}
	delete(km.keyToNode, node.Key)
}

func (km *KeyMaintainer) keyMappingRelationIsNeedToDelete(node *Node) bool {
	return km.keyToNode[node.Key] == node
}

// -------------- DoubleList ----------------
type DoubleList struct {
	dummyHead *Node
	dummyTail *Node
}

func NewDoubleList() *DoubleList {
	const invalidNumber = -100
	dummyHead := NewNode(invalidNumber, invalidNumber, invalidNumber)
	dummyTail := NewNode(invalidNumber, invalidNumber, invalidNumber)
	dummyHead.Next = dummyTail
	dummyTail.Pre = dummyHead
	return &DoubleList{dummyHead, dummyTail}
}

func (dl *DoubleList) GetLastNode() *Node {
	return dl.dummyTail.Pre
}
func (dl *DoubleList) InsertAfter(referenceNode, node *Node) {
	node.LinkAfter(referenceNode)
}
func (dl *DoubleList) DeleteNode(node *Node) {
	node.BreakLink()
}

// -------------- Node ----------------
type Node struct {
	Key       int
	Value     int
	Frequency int
	Pre, Next *Node
}

func NewNode(key, value, frequency int) *Node {
	return &Node{key, value, frequency, nil, nil}
}

func (n *Node) BreakLink() {
	n.Next.Pre = n.Pre
	n.Pre.Next = n.Next
	n.Next = nil
	n.Pre = nil
}
func (n *Node) LinkAfter(referenceNode *Node) {
	n.Pre = referenceNode
	n.Next = referenceNode.Next
	referenceNode.Next.Pre = n
	referenceNode.Next = n
}
func (n *Node) IsFirstNode() bool {
	return n.Pre == nil || n.Pre.Frequency != n.Frequency
}
