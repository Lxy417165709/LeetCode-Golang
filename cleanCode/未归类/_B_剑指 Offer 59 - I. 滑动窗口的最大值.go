package 未归类

func maxSlidingWindow(nums []int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	windowQueue := NewWindowQueue()
	return windowQueue.GetMaxValuesOfEachWindow(nums, k)
}

// ------------ WindowQueue ------------
type WindowQueue struct {
	monotonicallyDecreasingQueue *Queue

	handledNums           []int
	windowSize            int
	maxValuesInEachWindow []int
	readingIndex          int
}

func NewWindowQueue() *WindowQueue {
	return &WindowQueue{}
}

func (wq *WindowQueue) GetMaxValuesOfEachWindow(nums []int, windowSize int) []int {
	wq.handledNums, wq.windowSize, wq.monotonicallyDecreasingQueue = nums, windowSize, NewQueue()
	for wq.isReadingOver() {
		wq.removeNumberThatOverWindow()
		wq.removeImpossibleMaxValueInWindow()
		wq.pushReadingNumberIntoWindow()
		if wq.hasWindowFormed() {
			wq.storeWindowMaxValue()
		}
		wq.readNextOne()
	}
	return wq.maxValuesInEachWindow
}
//
func (wq *WindowQueue) isReadingOver() bool {
	return wq.readingIndex < len(wq.handledNums)
}

func (wq *WindowQueue) removeNumberThatOverWindow() {
	if wq.isQueueFirstNumberOverWindow() {
		wq.monotonicallyDecreasingQueue.RemoveFirstNumber()
	}
}

func (wq *WindowQueue) removeImpossibleMaxValueInWindow() {
	for wq.isReadingNumberGreaterThanQueueLastNumber() {
		wq.monotonicallyDecreasingQueue.RemoveLastNumber()
	}
}

func (wq *WindowQueue) pushReadingNumberIntoWindow() {
	wq.monotonicallyDecreasingQueue.PushBack(wq.getReadingNumber())
}

func (wq *WindowQueue) hasWindowFormed() bool {
	return wq.readingIndex >= wq.windowSize-1
}

func (wq *WindowQueue) storeWindowMaxValue() {
	wq.maxValuesInEachWindow = append(wq.maxValuesInEachWindow, wq.monotonicallyDecreasingQueue.GetFirstNumber())
}

func (wq *WindowQueue) readNextOne() {
	wq.readingIndex++
}
//
func (wq *WindowQueue) isQueueFirstNumberOverWindow() bool {
	return wq.isExistedNumberThatOverWindow() && wq.getNumberThatOverWindow() == wq.monotonicallyDecreasingQueue.GetFirstNumber()
}

func (wq *WindowQueue) isReadingNumberGreaterThanQueueLastNumber() bool {
	return !wq.monotonicallyDecreasingQueue.IsEmpty() &&
		wq.getReadingNumber() > wq.monotonicallyDecreasingQueue.GetLastNumber()
}
//
func (wq *WindowQueue) getNumberThatOverWindow() int {
	return wq.handledNums[wq.readingIndex-wq.windowSize]
}

func (wq *WindowQueue) isExistedNumberThatOverWindow() bool {
	return wq.readingIndex >= wq.windowSize
}

func (wq *WindowQueue) getReadingNumber() int {
	return wq.handledNums[wq.readingIndex]
}


// ------------ Queue ------------
type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{nums: make([]int, 0)}
}

func (q *Queue) PushBack(number int) {
	q.nums = append(q.nums, number)
}

func (q *Queue) RemoveFirstNumber() {
	q.nums = q.nums[1:]
}

func (q *Queue) RemoveLastNumber() {
	q.nums = q.nums[:q.getSize()-1]
}

func (q *Queue) GetFirstNumber() int {
	return q.nums[0]
}

func (q *Queue) GetLastNumber() int {
	return q.nums[q.getSize()-1]
}

func (q *Queue) IsEmpty() bool {
	return q.getSize() == 0
}
//
func (q *Queue) getSize() int {
	return len(q.nums)
}


/*
	题目链接: https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
	总结：
		1. 这题的面向对象使得代码增长了 2 倍。
	存在问题：
		1. window和queue的概念，以及它们的关系没有用代码的形式说明清楚。 （感觉可以划分出一个单调递减队列对象进行解决）
		2. windowQueue名字感觉不太好

*/
