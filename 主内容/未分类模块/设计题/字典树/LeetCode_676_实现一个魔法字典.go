package main


type TrieNode struct {
	endOfWordFlag bool
	charToNode          [26]*TrieNode
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

func (tn *TrieNode) SearchWordFromThis(word string, mismatchChance int) bool {
	if tn == nil || mismatchChance < 0 {
		return false
	}
	if word == "" {
		return tn.isEndOfWord() && mismatchChance == 0
	}

	readingChar := word[0]
	nextSearchWord := word[1:]
	for char := byte('a'); char <= 'z'; char++ {
		if char == readingChar {
			continue
		}
		// replace the readingChar with char(a->z), and rematch
		if tn.getNextNode(char).SearchWordFromThis(nextSearchWord, mismatchChance-1) {
			return true
		}
	}
	return tn.getNextNode(readingChar).SearchWordFromThis(nextSearchWord, mismatchChance)
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

type MagicDictionary struct {
	trieRoot       *TrieNode
	mismatchChance int
}



/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{NewTrieNode(), 1}
}

/** Build a dictionary through a list of words */
func (md *MagicDictionary) BuildDict(dict []string) {
	for _, word := range dict {
		md.trieRoot.AddWordFromThis(word)
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (md *MagicDictionary) Search(word string) bool {
	return md.trieRoot.SearchWordFromThis(word, md.mismatchChance)
}

/*
	总结
	1. 这题我通过构建字典树解决了
	2. 在字典树和字典树节点的职责划分方面，我花了很多时间
*/
