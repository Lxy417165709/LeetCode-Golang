package 前缀树

// -------------------------------------------------- 1. 前缀树(开始) --------------------------------------------------

// MyTrie 前缀树。
type MyTrie struct {
	DummyRoot *MyTrieNode
}

// NewMyTrie 构建前缀树。
func NewMyTrie() *MyTrie {
	return &MyTrie{DummyRoot: &MyTrieNode{
		CanBeWord:   false,
		keyToSonMap: make(map[string]*MyTrieNode),
	}}
}

// Insert 向前缀树插入单词。
func (m *MyTrie) Insert(word string) {
	curNode := m.DummyRoot
	for _, char := range word {
		nextNode, ok := curNode.keyToSonMap[string(char)]
		if !ok {
			nextNode = &MyTrieNode{keyToSonMap: make(map[string]*MyTrieNode)}
			curNode.keyToSonMap[string(char)] = nextNode
		}
		curNode = nextNode
	}
	curNode.CanBeWord = true
}

// SearchWord 搜索单词，单词存在返回 true，否则返回 false。
func (m *MyTrie) SearchWord(word string) bool {
	node := m.locateNode(word)
	return node != nil && node.CanBeWord
}

// SearchPrefix 搜索前缀，前缀存在返回 true，否则返回 false。
func (m *MyTrie) SearchPrefix(prefix string) bool {
	node := m.locateNode(prefix)
	return node != nil
}

// locateNode 定位节点。
func (m *MyTrie) locateNode(str string) *MyTrieNode {
	curNode := m.DummyRoot
	for _, char := range str {
		nextNode, ok := curNode.keyToSonMap[string(char)]
		if !ok {
			return nil
		}
		curNode = nextNode
	}
	return curNode
}

// -------------------------------------------------- 1. 前缀树(结束) --------------------------------------------------

// -------------------------------------------------- 2. 树节点(开始) --------------------------------------------------

// MyTrieNode 前缀树节点。
type MyTrieNode struct {
	CanBeWord   bool
	keyToSonMap map[string]*MyTrieNode
}

// -------------------------------------------------- 2. 树节点(结束) --------------------------------------------------
