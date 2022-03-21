package 哈希

type MyHashMap struct {
	Hash *MyHash
}

func Constructor() MyHashMap {
	return MyHashMap{Hash: NewHashTable(100000)}
}

func (m *MyHashMap) Put(key int, value int) {
	m.Hash.Set(key, value)
}

func (m *MyHashMap) Get(key int) int {
	value := m.Hash.Get(key)
	if value == nil {
		return -1
	}
	return value.(int)
}

func (m *MyHashMap) Remove(key int) {
	m.Hash.Remove(key)
}
