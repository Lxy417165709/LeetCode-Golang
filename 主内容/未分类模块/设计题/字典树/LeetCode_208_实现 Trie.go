package main

type TrieNode struct {
	endOfWordFlag bool
	charToNode    [26]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{false, [26]*TrieNode{}}
}
func (tn *TrieNode) AddWordFromThis(word string) {
	trieNode := tn
	for i := 0; i < len(word); i++ {
		char := word[i]
		if !trieNode.charIsExist(char) {
			trieNode.addChar(char)
		}
		trieNode = trieNode.getNextNode(char)
	}
	trieNode.endOfWordMark()
}

func (tn *TrieNode) SearchWordFromThis(word string) bool {
	if tn == nil {
		return false
	}
	if word == "" {
		return tn.isEndOfWord()
	}

	readingChar := word[0]
	nextSearchWord := word[1:]
	return tn.getNextNode(readingChar).SearchWordFromThis(nextSearchWord)
}

func (tn *TrieNode) SearchPrefixFromThis(prefix string) bool{
	if tn == nil {
		return false
	}
	if prefix == "" {
		return true
	}
	readingChar := prefix[0]
	nextSearchPrefix := prefix[1:]
	return tn.getNextNode(readingChar).SearchPrefixFromThis(nextSearchPrefix)
}
func (tn *TrieNode) charIsExist(char byte) bool {
	return tn.charToNode[getCharIdentity(char)] != nil
}

func (tn *TrieNode) addChar(char byte) {
	tn.charToNode[getCharIdentity(char)] = NewTrieNode()
}

func (tn *TrieNode) getNextNode(char byte) *TrieNode {
	return tn.charToNode[getCharIdentity(char)]
}

func (tn *TrieNode) endOfWordMark() {
	tn.endOfWordFlag = true
}

func (tn *TrieNode) isEndOfWord() bool {
	return tn.endOfWordFlag
}

func getCharIdentity(char byte) int {
	return int(char - 'a')
}

type Trie struct {
	root *TrieNode
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{NewTrieNode()}
}


/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
	t.root.AddWordFromThis(word)
}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	return t.root.SearchWordFromThis(word)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.root.SearchPrefixFromThis(prefix)
}
