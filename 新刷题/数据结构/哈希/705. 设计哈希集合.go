package 哈希

const foo = 1

type MyHashSet struct {
	Hash *MyHash
}

func Constructor() MyHashSet {
	return MyHashSet{Hash: NewHashTable(100000)}
}

func (m *MyHashSet) Add(key int) {
	m.Hash.Set(key, foo)
}

func (m *MyHashSet) Remove(key int) {
	m.Hash.Remove(key)
}

func (m *MyHashSet) Contains(key int) bool {
	return m.Hash.Get(key) != nil
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
