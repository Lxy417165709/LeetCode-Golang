package main

func reverseLeftWords(s string, n int) string {
	stringHandler := NewStringHandler()
	return stringHandler.ReverseLeftChars(s, n)
}

// ------------ StringHandler ---------
type StringHandler struct {
}

func NewStringHandler() *StringHandler {
	return &StringHandler{}
}

func (sh *StringHandler) ReverseLeftChars(s string, countOfReversing int) string {
	runes := []rune(s)
	reversePartOfRunes(runes, NewInterval(0, countOfReversing-1))
	reversePartOfRunes(runes, NewInterval(countOfReversing, len(runes)-1))
	reversePartOfRunes(runes, NewInterval(0, len(runes)-1))
	return string(runes)
}

func reversePartOfRunes(runes []rune, reverseInterval *Interval) {
	sumOfLeftAndRightIndex := reverseInterval.Left + reverseInterval.Right
	middleIndex := sumOfLeftAndRightIndex / 2
	for i := reverseInterval.Left; i <= middleIndex; i++ {
		runes[i], runes[sumOfLeftAndRightIndex-i] = runes[sumOfLeftAndRightIndex-i], runes[i]
	}
}

// ------------ Interval ---------
type Interval struct {
	Left  int
	Right int
}

func NewInterval(left, right int) *Interval {
	return &Interval{Left: left, Right: right}
}


/*
	题目链接：https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
*/
