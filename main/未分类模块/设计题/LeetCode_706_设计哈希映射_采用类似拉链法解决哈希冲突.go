package 设计题

type Node struct {
	Key   int
	Value int
}

const m = 10003

type MyHashMap struct {
	arr [10005][]Node
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{[10005][]Node{}}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	hashKey := key % m
	if len(this.arr[hashKey]) == 0 {
		this.arr[hashKey] = make([]Node, 0)
	}
	// 这里要注意，不能直接append,而是要检测该键值是否存在，存在则修改，否则则append
	for i := 0; i < len(this.arr[hashKey]); i++ {
		if this.arr[hashKey][i].Key == key {
			this.arr[hashKey][i].Value = value
			return
		}
	}

	this.arr[hashKey] = append(this.arr[hashKey], Node{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hashKey := key % m
	for i := 0; i < len(this.arr[hashKey]); i++ {
		if this.arr[hashKey][i].Key == key {
			return this.arr[hashKey][i].Value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	hashKey := key % m
	for i := 0; i < len(this.arr[hashKey]); i++ {
		if this.arr[hashKey][i].Key == key {
			// 这里只是为了该键值对，我把Key改-1，表示该键值对不存在了。
			this.arr[hashKey][i].Key = -1
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
