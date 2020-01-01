package 设计题

type MyHashMap struct {
	arr [100009]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	hashMap := MyHashMap{[100009]int{}}
	for i := 0; i < len(hashMap.arr); i++ {
		hashMap.arr[i] = -1
	}
	return hashMap
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.arr[key%100007] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.arr[key%100007]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.arr[key%100007] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

 /*
 	总结
 	1. 这里采用除留余数法来进行哈希，使数组空间不用开辟1E6个，而只开辟1E5个。 (之所以这样哈希，是因为操作只有1E4个，采用1E5个空间的哈希冲突已经很小了)
 	2. 这题还可以使用链表解决哈希冲突
 */
