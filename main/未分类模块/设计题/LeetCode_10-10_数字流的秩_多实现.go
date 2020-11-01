package main

import (
	"fmt"
	"reflect"
)

//	----------- StreamRank -----------
type StreamRank struct {
	streamRanker StreamRanker
}

func Constructor() StreamRank {
	return StreamRank{NewLineTree(0, 50000)}
}

func (sr *StreamRank) Track(x int) {
	sr.streamRanker.Track(x)
}
func (sr *StreamRank) GetRankOfNumber(x int) int {
	return sr.streamRanker.GetRankOfNumber(x)
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

// 相当于C++的 upperBound 函数
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
type TreeArray struct {
	minimumInDataStream int
	maximumInDataStream int

	// theQuantityOfNumbersInInterval[intervalIdentity] 表示: 区间[intervalIdentity-getLowBit(intervalIdentity), intervalIdentity] 内的元素个数
	theQuantityOfNumbersInInterval []int
}

func NewTreeArray(minimumInDataStream, maximumInDataStream int) *TreeArray {
	buffer := make([]int, maximumInDataStream-minimumInDataStream+1)
	return &TreeArray{minimumInDataStream, maximumInDataStream, buffer}
}

func (ta *TreeArray) Track(number int) {
	for intervalIdentity := ta.getIntervalIdentity(number); ta.intervalIdentityIsValid(intervalIdentity); intervalIdentity += getLowBit(intervalIdentity) {
		ta.theQuantityOfNumbersInInterval[intervalIdentity] ++
	}
}
func (ta *TreeArray) GetRankOfNumber(number int) int {
	rank := 0

	for intervalIdentity := ta.getIntervalIdentity(number); ta.intervalIdentityIsValid(intervalIdentity); intervalIdentity -= getLowBit(intervalIdentity) {
		rank += ta.theQuantityOfNumbersInInterval[intervalIdentity]
	}
	return rank
}
func (ta *TreeArray) intervalIdentityIsValid(intervalIdentity int) bool {
	return intervalIdentity > 0 && intervalIdentity <= ta.getIntervalIdentity(ta.maximumInDataStream)
}
func (ta *TreeArray) getIntervalIdentity(number int) int {
	return number - ta.minimumInDataStream + 1
}

// 由于 number 可能不是正数，所以需要将它转换为正数，这样才能通过 buffer 存储
// 如当 minNumber = -10, maxNumber = 100时
// number == -10，会被存储到 1
// number == 0，会被存储到 11
// number == 100，会被存储到 111

// ------------ LineTree -----------
type LineTree struct {
	root *LineTreeNode
	minimumInDataStream int
	maximumInDataStream int
}
func NewLineTree(minimumInDataStream, maximumInDataStream int) *LineTree {
	root := NewLineTreeNode(minimumInDataStream, maximumInDataStream)
	return &LineTree{root,minimumInDataStream,maximumInDataStream}
}
func (lt *LineTree) Track(number int) {
	lt.root.insertNumber(number)
}
func (lt *LineTree) GetRankOfNumber(number int) int {
	return lt.root. getTheQuantityOfNumberBetweenIntervalLeftBoundaryAndReferenceBoundary(number)
}

// ------------ LineTreeNode -----------
// TODO: 重构
type LineTreeNode struct {
	intervalLeftBoundary          int
	intervalRightBoundary          int
	theQuantityOfNumbersInInterval int
	leftSon             *LineTreeNode
	rightSon            *LineTreeNode
}

func NewLineTreeNode(intervalLeftBoundary, intervalRightBoundary int) *LineTreeNode {
	if intervalLeftBoundary > intervalRightBoundary {
		return nil
	}
	return &LineTreeNode{
		intervalLeftBoundary,
		intervalRightBoundary,
		0,
		nil,
		nil,
	}
}

func (ltn *LineTreeNode) insertNumber(number int) {
	if !ltn.numberIsInInterval(number) {
		return
	}
	ltn.theQuantityOfNumbersInInterval++
	if ltn.isLeafNode() {
		return
	}
	ltn.generateSon()
	if ltn.numberIsInLeftPartOfInterval(number) {
		ltn.leftSon.insertNumber(number)
	}else{
		ltn.rightSon.insertNumber(number)
	}
}
func (ltn *LineTreeNode) getTheQuantityOfNumberBetweenIntervalLeftBoundaryAndReferenceBoundary(referenceBoundary int) int {
	if !ltn.numberIsInInterval(referenceBoundary){
		return 0
	}
	if ltn.isLeafNode() {
		return ltn.theQuantityOfNumbersInInterval
	}
	ltn.generateSon()
	if ltn.numberIsInLeftPartOfInterval(referenceBoundary){
		return ltn.leftSon.getTheQuantityOfNumberBetweenIntervalLeftBoundaryAndReferenceBoundary(referenceBoundary)
	}else{
		return ltn.leftSon.theQuantityOfNumbersInInterval + ltn.rightSon.getTheQuantityOfNumberBetweenIntervalLeftBoundaryAndReferenceBoundary(referenceBoundary)
	}
}

func (ltn *LineTreeNode) numberIsInInterval(number int) bool{
	return  number >= ltn.intervalLeftBoundary && number <= ltn.intervalRightBoundary
}
func (ltn *LineTreeNode) numberIsInLeftPartOfInterval(number int) bool{
	return  number<=ltn.getIntervalMidPoint()
}
func (ltn *LineTreeNode) getIntervalMidPoint() int{
	return (ltn.intervalLeftBoundary + ltn.intervalRightBoundary) / 2
}
func (ltn *LineTreeNode) isLeafNode() bool{
	return ltn.intervalLeftBoundary == ltn.intervalRightBoundary
}

// 命名有待商榷
func (ltn *LineTreeNode) generateSon() {
	if ltn.leftSon == nil {
		ltn.leftSon = NewLineTreeNode(ltn.intervalLeftBoundary, ltn.getIntervalMidPoint())
	}
	if ltn.rightSon == nil {
		ltn.rightSon = NewLineTreeNode(ltn.getIntervalMidPoint()+1, ltn.intervalRightBoundary)
	}
}
// ------------- 公共函数 -------------
const INF = 100000000

func getMinNumber(array []int) int {
	minNumber := INF
	for _, number := range array {
		minNumber = min(minNumber, number)
	}
	return minNumber
}
func getMaxNumber(array []int) int {
	maxNumber := -INF
	for _, number := range array {
		maxNumber = max(maxNumber, number)
	}
	return maxNumber
}
func getLowBit(positiveNumber int) int {
	return positiveNumber & -positiveNumber
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
