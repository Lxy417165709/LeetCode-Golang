package main

const INF = 100000000

func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	treeArray := NewTreeArray(getMinNumber(nums), getMaxNumber(nums))

	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = treeArray.GetRankOfNumber(nums[i] - 1)
		treeArray.Track(nums[i])
	}
	return result
}

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

// ----------- TreeArray ----------
type TreeArray struct {
	minNumber         int
	maxNumber         int
	offset            int
	theCountOfNumbers []int // buffer[x] 表示 区间[x-getLowBit(x), x] 内的元素个数
}

func NewTreeArray(minNumber, maxNumber int) *TreeArray {
	offset := -minNumber + 1
	buffer := make([]int, maxNumber+offset+1)
	return &TreeArray{minNumber, maxNumber, offset, buffer}
}

func (ta *TreeArray) Track(number int) {
	for nextNumber := ta.convertToPositiveNumber(number); ta.isInBoundary(nextNumber); nextNumber += ta.getLowBit(nextNumber) {
		ta.theCountOfNumbers[nextNumber] ++
	}
}
func (ta *TreeArray) GetRankOfNumber(number int) int {
	rank := 0
	for nextNumber := ta.convertToPositiveNumber(number); ta.isInBoundary(nextNumber); nextNumber -= ta.getLowBit(nextNumber) {
		rank += ta.theCountOfNumbers[nextNumber]
	}
	return rank
}

func (ta *TreeArray) convertToPositiveNumber(number int) int {
	return number + ta.offset
}
func (ta *TreeArray) isInBoundary(positiveNumber int) bool {
	return positiveNumber > 0 && positiveNumber <= ta.convertToPositiveNumber(ta.maxNumber)
}
func (ta *TreeArray) getLowBit(positiveNumber int) int {
	return positiveNumber & -positiveNumber
}


// 这题与 https://leetcode-cn.com/problems/rank-from-stream-lcci/ 的解法是一样的，不过这题更难一些
