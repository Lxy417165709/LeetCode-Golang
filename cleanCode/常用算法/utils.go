package main

const INF = 1000000000000

// ------------ About Array ---------
func getMaxSumOfSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSumOfSubArrayEndPreNum := make([]int, len(nums)+1)
	maxSumOfSubArrayEndPreNum[0] = -INF
	for i := 1; i <= len(nums); i++ {
		maxSumOfSubArrayEndPreNum[i] = max(maxSumOfSubArrayEndPreNum[i-1]+nums[i-1], nums[i-1])
	}
	return getMax(maxSumOfSubArrayEndPreNum)
}

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
