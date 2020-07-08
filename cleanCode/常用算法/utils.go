package main

import "math/rand"

const INF = 1000000000000

// ------------ About Array ---------
func getCountOfGreaterElement(nums []int, ref int) int {
	return len(nums) - getIndexOfFirstGreater(nums, ref)
}

func getLengthOfLAS(nums []int) int {
	LAS := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		index := getIndexOfFirstGreaterOrEqual(LAS, nums[i])
		if index == len(LAS) {
			LAS = append(LAS, nums[i])
		}
		LAS[index] = nums[i]
	}
	return len(LAS)
}

func getMax(nums []int) int {
	result := -INF
	for i := 0; i < len(nums); i++ {
		result = max(nums[i], result)
	}
	return result
}

func getMin(nums []int) int {
	result := INF
	for i := 0; i < len(nums); i++ {
		result = min(nums[i], result)
	}
	return result
}


func getCountOfCombination(arrayOfElementGreaterThanZero []int, sumOfGreaterThanOrEqualZero int) int {
	countOfCombinationWithSpecificSum := make([]int, sumOfGreaterThanOrEqualZero+1)
	countOfCombinationWithSpecificSum[0] = 1
	for i := 0; i < len(arrayOfElementGreaterThanZero); i++ {
		for t := 1; t <= sumOfGreaterThanOrEqualZero; t++ {
			if t-arrayOfElementGreaterThanZero[i] >= 0 {
				countOfCombinationWithSpecificSum[t] += countOfCombinationWithSpecificSum[t-arrayOfElementGreaterThanZero[i]]
			}
		}
	}
	return countOfCombinationWithSpecificSum[sumOfGreaterThanOrEqualZero]
}
func getCountOfPermutation(arrayOfElementGreaterThanZero []int, sumOfGreaterThanOrEqualZero int) int {
	countOfPermutationWithSpecificSum := make([]int, sumOfGreaterThanOrEqualZero+1)
	countOfPermutationWithSpecificSum[0] = 1
	for i := 1; i <= sumOfGreaterThanOrEqualZero; i++ {
		for t := 0; t < len(arrayOfElementGreaterThanZero); t++ {
			if i-arrayOfElementGreaterThanZero[t] >= 0 {
				countOfPermutationWithSpecificSum[i] += countOfPermutationWithSpecificSum[i-arrayOfElementGreaterThanZero[t]]
			}
		}
	}
	return countOfPermutationWithSpecificSum[sumOfGreaterThanOrEqualZero]
}

// ------------ About Judging ---------
func isLetter(char byte) bool {
	return char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z'
}

// ------------ About Forming ---------
func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

func makeSliceINF(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = INF
	}
}

// ------------ About Comparing ---------
func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := max(arr[1:]...)
	if a < b {
		return b
	}
	return a
}
func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := min(arr[1:]...)
	if a > b {
		return b
	}
	return a
}

// ------------ About Searching ---------
func getIndexOfFirstGreater(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func getIndexOfFirstGreaterOrEqual(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func getIndexOfLastSmall(arr []int, ref int) int {
	return getIndexOfFirstGreaterOrEqual(arr, ref) - 1
}
func getIndexOfLastSmallOrEqual(arr []int, ref int) int {
	return getIndexOfFirstGreater(arr, ref) - 1
}

func makeKthSmallToRightPlace(nums []int, KthSmall int) {
	// 例子: A 的 rightPlace 就是：A处于某个位置，这个位置左边的元素都小于等于它，右边的元素都大于等于它。

	l, r := 0, len(nums)-1
	for l <= r {
		index := randomPartition(nums, l, r)
		XthSmall := index + 1
		if XthSmall == KthSmall {
			return
		} else {
			if XthSmall > KthSmall {
				r = index - 1
			} else {
				l = index + 1
			}
		}
	}
	// 到达这里表示:  原数组不存在第K小
}
func makeKthBigToRightPlace(nums []int, KthBig int) {
	makeKthSmallToRightPlace(nums, len(nums)-KthBig+1)
}

func randomPartition(nums []int, l int, r int) int {
	// 返回值是一个索引，假如记做 index
	// 则 nums[index] 位于 rightPlace
	// 同 partition函数

	upsetArrayRandomly(nums, l, r)
	return partition(nums, l, r)
}
func upsetArrayRandomly(nums []int, l int, r int) {
	randomIndex := rand.Intn(r-l+1) + l
	nums[randomIndex], nums[l] = nums[l], nums[randomIndex]
}
func partition(nums []int, l int, r int) int {
	refIndex := l // 这个索引取 (l+r)/2 也可以
	for l <= r {
		for l <= r && nums[l] <= nums[refIndex] {
			l++
		}
		for l <= r && nums[r] >= nums[refIndex] {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[refIndex], nums[r] = nums[r], nums[refIndex]
	return r
}

// ------------ Interval ---------
type Interval struct {
	Left  int
	Right int
}

func NewInterval(left, right int) *Interval {
	return &Interval{Left: left, Right: right}
}

// ----------- Coordinate ----------
type Coordinate struct {
	X int
	Y int
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{x, y}
}

// ----------- Vector ----------
type Vector struct {
	X int
	Y int
}

func NewVector(x, y int) *Vector {
	return &Vector{x, y}
}
