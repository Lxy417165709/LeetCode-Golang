package 哈希

// ------------------------------------------------ 1. 哈希表(开始) ------------------------------------------------

// MyHash 哈希表。
type MyHash struct {
	Buckets []*MyList // 桶。
}

func NewHashTable(bucketCount int) *MyHash {
	buckets := make([]*MyList, bucketCount)
	for i := 0; i < len(buckets); i++ {
		buckets[i] = &MyList{}
	}
	return &MyHash{
		Buckets: buckets,
	}
}

func (h *MyHash) Set(key int, value interface{}) {
	index := h.getBucketIndex(key)
	h.Buckets[index].Insert(key, value)
}

func (h *MyHash) Get(key int) interface{} {
	index := h.getBucketIndex(key)
	node := h.Buckets[index].Find(key)
	if node == nil {
		return nil
	}
	return node.Value
}

func (h *MyHash) Remove(key int) {
	index := h.getBucketIndex(key)
	h.Buckets[index].Remove(key)
}

func (h *MyHash) getBucketIndex(key int) int {
	return key % len(h.Buckets)
}

// ------------------------------------------------ 1. 哈希表(结束) ------------------------------------------------

// ------------------------------------------------ 2. 链表(开始) ------------------------------------------------

// MyList 链表。
type MyList struct {
	Head *MyListNode
	Tail *MyListNode
}

// Insert 将元素插入链表。
func (l *MyList) Insert(key int, value interface{}) {
	// 1. 初始化插入节点。
	insertedNode := &MyListNode{
		Key:   key,
		Value: value,
	}

	// 2. 链表为空时。
	if l.Head == nil {
		l.Head = insertedNode
		l.Tail = insertedNode
		return
	}

	// 3. 链表不为空时。
	targetNode := l.Find(key)
	if targetNode == nil {
		// 3.1 不存在相同的键时，插入新节点。
		l.Tail.Next = insertedNode
		insertedNode.Pre = l.Tail
		l.Tail = l.Tail.Next
	} else {
		// 3.1 存在相同的键时，覆盖节点值。
		targetNode.Value = value
	}
}

// Find 寻找对应key的节点。
func (l *MyList) Find(key int) *MyListNode {
	node := l.Head
	for node != nil {
		if node.Key == key {
			return node
		}
		node = node.Next
	}
	return nil
}

// Remove 删除对应key的节点。
func (l *MyList) Remove(key int) {
	// 1. 找到目标节点。
	targetNode := l.Find(key)

	// 2. 目标节点为空时。
	if targetNode == nil {
		return
	}

	// 3. 目标节点非空时。
	if targetNode == l.Head {
		// 3.1 目标节点为头节点。
		l.Head = l.Head.Next
		if l.Head != nil {
			l.Head.Pre = nil
		}
	} else if targetNode == l.Tail {
		// 3.2 目标节点为尾节点。
		l.Tail = l.Tail.Pre
		if l.Tail != nil {
			l.Tail.Next = nil
		}
	} else {
		// 3.3 目标节点为中间节点。
		targetNode.Pre.Next = targetNode.Next
		targetNode.Next.Pre = targetNode.Pre
	}
}

// ------------------------------------------------ 2. 链表(结束) ------------------------------------------------

// ------------------------------------------------ 3. 链表节点(开始) ------------------------------------------------

// MyListNode 链表节点。
type MyListNode struct {
	Key   int
	Value interface{}
	Next  *MyListNode
	Pre   *MyListNode
}

// ------------------------------------------------ 3. 链表节点(结束) ------------------------------------------------
