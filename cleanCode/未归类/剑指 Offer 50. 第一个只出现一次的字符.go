package 未归类

func firstUniqChar(s string) byte {
	charStorage := NewCharStorage()

	return byte(charStorage.GetFirstUniqueChar(s))
}

// --------- CharStorage ---------
type CharStorage struct {
	charAppearingTimes                 map[rune]int
	returnCharWhenUniqueCharDoesExist  rune
	sample string
}

func NewCharStorage() *CharStorage {
	return &CharStorage{
		returnCharWhenUniqueCharDoesExist: ' ',
	}
}

func (sv *CharStorage) GetFirstUniqueChar(s string) rune {
	sv.sample = s
	sv.formCharAppearingTimes()
	return sv.findOutFirstUniqueChar()
}

func (sv *CharStorage) formCharAppearingTimes() {
	sv.charAppearingTimes = make(map[rune]int)
	for _, char := range sv.sample {
		sv.charAppearingTimes[char]++
	}
}

func (sv *CharStorage) findOutFirstUniqueChar() rune {
	for _, char := range sv.sample {
		if sv.isCharUnique(char) {
			return char
		}
	}
	return sv.returnCharWhenUniqueCharDoesExist
}

func (sv *CharStorage) isCharUnique(char rune) bool {
	return sv.charAppearingTimes[char] == 1
}




/*
	题目链接: https://leetcode-cn.com/problemset/all/
	总结：
		1. 个人感觉这个代码不需要重构，因为原来的代码就很短了...这一重构，虽然职责分明了，
		   但是代码量增长了4倍...
		2. 重构使得代码可重用性增强了
		3. 当功能单一的时候，重构的必要性不是那么强。
			(比如从命名方面来说，CharStorage其实不够准确，准确的应该是 StringHandler,而它们的关系是：StringHandler包含 CharStorage 了)
*/
