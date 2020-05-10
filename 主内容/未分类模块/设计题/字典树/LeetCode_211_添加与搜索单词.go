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
	if readingChar == '.' {
		for char := byte('a'); char <= 'z'; char++ {
			// replace the readingChar with char(a->z), and rematch
			if tn.getNextNode(char).SearchWordFromThis(nextSearchWord) {
				return true
			}
		}
		return false
	}
	return tn.getNextNode(readingChar).SearchWordFromThis(nextSearchWord)
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

type WordDictionary struct {
	trie *TrieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{NewTrieNode()}
}

/** Adds a word into the data structure. */
func (wd *WordDictionary) AddWord(word string) {
	wd.trie.AddWordFromThis(word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (wd *WordDictionary) Search(word string) bool {
	return wd.trie.SearchWordFromThis(word)
}
