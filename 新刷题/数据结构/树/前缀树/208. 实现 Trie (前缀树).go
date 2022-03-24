package 前缀树

type Trie struct {
	MyTrie *MyTrie
}

func Constructor() Trie {
	return Trie{MyTrie: NewMyTrie()}
}

func (t *Trie) Insert(word string) {
	t.MyTrie.Insert(word)
}

func (t *Trie) Search(word string) bool {
	return t.MyTrie.SearchWord(word)
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.MyTrie.SearchPrefix(prefix)
}
