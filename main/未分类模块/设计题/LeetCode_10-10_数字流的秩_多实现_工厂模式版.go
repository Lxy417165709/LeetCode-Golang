package main

import (
	"fmt"
	"reflect"
)

// 下面的代码用到了：单例模式、工厂模式、策略模式

// ------------ StreamRankerFlag -----------
type StreamRankerFlag int

const (
	OrderedArrayFlag StreamRankerFlag = iota
	TreeArrayFlag
	LineTreeFlag
)

func init() {
	streamRankerFactory := GetUniqueStreamRankerFactory()
	streamRankerFactory.Register(OrderedArrayFlag, NewOrderedArray)
	streamRankerFactory.Register(TreeArrayFlag, NewTreeArray)
	streamRankerFactory.Register(LineTreeFlag, NewLineTree)
}

//	----------- StreamRank -----------
type StreamRank struct {
	streamRanker StreamRanker
}

func Constructor() StreamRank {
	streamRankerFactory := GetUniqueStreamRankerFactory()
	return StreamRank{streamRankerFactory.GetStreamRanker(TreeArrayFlag)}
}

func (sr *StreamRank) Track(x int) {
	sr.streamRanker.Track(x)
}
func (sr *StreamRank) GetRankOfNumber(x int) int {
	return sr.streamRanker.GetRankOfNumber(x)
}

// ------------ StreamRankerFactory -----------
// 通过反射，这个工厂可以符合开闭原则
// TODO: 重构
type StreamRankerFactory struct {
	flagToStreamRankerGenerateFunction map[StreamRankerFlag]interface{}
}

func newStreamRankerFactory() *StreamRankerFactory {
	return &StreamRankerFactory{
		make(map[StreamRankerFlag]interface{}),
	}
}

var streamRankerFactory = newStreamRankerFactory()

func GetUniqueStreamRankerFactory() *StreamRankerFactory {
	return streamRankerFactory
}
func (srf *StreamRankerFactory) GetStreamRanker(flag StreamRankerFlag) StreamRanker {
	ref := reflect.ValueOf(srf.flagToStreamRankerGenerateFunction[flag])
	value := ref.Call([]reflect.Value{})[0]
	return value.Interface().(StreamRanker)
}

func (srf *StreamRankerFactory) Register(flag StreamRankerFlag, generateFunction interface{}) {
	srf.flagToStreamRankerGenerateFunction[flag] = generateFunction
}

//	----------- StreamRanker -----------
type StreamRanker interface {
	Track(element int)
	GetRankOfNumber(element int) int
}

//	----------- OrderedArray -----------
type OrderedArray struct {
	numbers []int
}

func NewOrderedArray() *OrderedArray {
	return &OrderedArray{make([]int, 0)}
}

func (oa *OrderedArray) Track(x int) {
	oa.insertNumber(x)
	oa.insertSort()
}
func (oa *OrderedArray) GetRankOfNumber(x int) int {
	return oa.getTheIndexOfFirstNumberGreaterThanReference(x)
}

func (oa *OrderedArray) insertNumber(element int) {
	oa.numbers = append(oa.numbers, element)
}
func (oa *OrderedArray) insertSort() {
	for i := len(oa.numbers) - 1; i >= 1; i-- {
		if oa.numbers[i-1] > oa.numbers[i] {
			oa.numbers[i-1], oa.numbers[i] = oa.numbers[i], oa.numbers[i-1]
		} else {
			break
		}
	}
}

// 相当于C++的 upperbound 函数
func (oa *OrderedArray) getTheIndexOfFirstNumberGreaterThanReference(reference int) int {
	leftBoard, rightBoard := 0, len(oa.numbers)-1
	for leftBoard <= rightBoard {
		readingIndex := (leftBoard + rightBoard) / 2
		readingNumber := oa.numbers[readingIndex]
		if readingNumber == reference {
			leftBoard = readingIndex + 1
		} else {
			if readingNumber > reference {
				rightBoard = readingIndex - 1
			} else {
				leftBoard = readingIndex + 1
			}
		}
	}
	return leftBoard
}

// ----------- TreeArray -----------
// TODO: 重构
type TreeArray struct {
	minNumber         int
	maxNumber         int
	offset            int
	theQuantityOfNumbersInInterval []int // buffer[x] 表示 区间[x-getLowBit(x), x] 内的元素个数
}

func NewTreeArray() *TreeArray {
	minNumber := -10000
	maxNumber := 50000
	offset := -minNumber + 1
	buffer := make([]int, maxNumber+offset+1)
	return &TreeArray{minNumber, maxNumber, offset, buffer}
}

func (ta *TreeArray) Track(number int) {
	for nextNumber := ta.convertToPositiveNumber(number); ta.isInBoundary(nextNumber); nextNumber += getLowBit(nextNumber) {
		ta.theQuantityOfNumbersInInterval[nextNumber] ++
	}
}
func (ta *TreeArray) GetRankOfNumber(number int) int {
	rank := 0
	for nextNumber := ta.convertToPositiveNumber(number); ta.isInBoundary(nextNumber); nextNumber -= getLowBit(nextNumber) {
		rank += ta.theQuantityOfNumbersInInterval[nextNumber]
	}
	return rank
}
func (ta *TreeArray) convertToPositiveNumber(number int) int {
	return number + ta.offset
}
func (ta *TreeArray) isInBoundary(positiveNumber int) bool {
	return positiveNumber > 0 && positiveNumber <= ta.convertToPositiveNumber(ta.maxNumber)
}


// 由于 number 可能不是正数，所以需要将它转换为正数，这样才能通过 buffer 存储
// 如当 minNumber = -10, maxNumber = 100时
// number == -10，会被存储到 1
// number == 0，会被存储到 11
// number == 100，会被存储到 111

// ------------ LineTree -----------
type LineTree struct {
	root *LineTreeNode
}

func NewLineTree() *LineTree {
	return &LineTree{NewLineTreeNode(0, 50000)}
}
func (lt *LineTree) Track(number int) {
	lt.root.insertNumber(number)
}
func (lt *LineTree) GetRankOfNumber(number int) int {
	return lt.root.getTheQuantityOfNumberInBoard(number)
}

// ------------ LineTreeNode -----------
// TODO: 重构
type LineTreeNode struct {
	leftBoard        int
	rightBoard       int
	theQuantityOfNumber int
	leftSon          *LineTreeNode
	rightSon         *LineTreeNode
}

func NewLineTreeNode(leftBoard, rightBoard int) *LineTreeNode {
	if leftBoard > rightBoard {
		return nil
	}
	return &LineTreeNode{leftBoard, rightBoard, 0, nil, nil}
}

func (ltn *LineTreeNode) insertNumber(number int) {

	if number < ltn.leftBoard || number > ltn.rightBoard {
		return
	}
	ltn.theQuantityOfNumber++
	if ltn.rightBoard == ltn.leftBoard {
		return
	}

	mid := (ltn.leftBoard + ltn.rightBoard) / 2
	if ltn.leftSon == nil {
		ltn.leftSon = NewLineTreeNode(ltn.leftBoard, mid)
	}
	ltn.leftSon.insertNumber(number)
	if ltn.rightSon == nil {
		ltn.rightSon = NewLineTreeNode(mid+1, ltn.rightBoard)
	}
	ltn.rightSon.insertNumber(number)
}
func (ltn *LineTreeNode) getTheQuantityOfNumberInBoard(rightBoard int) int {
	if ltn == nil {
		return 0
	}
	switch {
	case ltn.leftBoard > rightBoard:
		return 0
	case ltn.rightBoard <= rightBoard:
		return ltn.theQuantityOfNumber
	case ltn.rightBoard > rightBoard:
		return ltn.leftSon.getTheQuantityOfNumberInBoard(rightBoard) + ltn.rightSon.getTheQuantityOfNumberInBoard(rightBoard)
	}
	return 0
}










// ------------- 公共函数 -------------
const INF = 100000000
func getMinNumber(array []int) int{
	minNumber := INF
	for _,number := range array{
		minNumber = min(minNumber,number)
	}
	return minNumber
}
func getMaxNumber(array []int) int{
	maxNumber := -INF
	for _,number := range array{
		maxNumber = max(maxNumber,number)
	}
	return maxNumber
}
func getLowBit(positiveNumber int) int {
	return positiveNumber & -positiveNumber
}
func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
