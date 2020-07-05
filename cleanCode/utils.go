package main


// ------------ About Forming ---------
func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
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
func min(arr ...int) int{
	if len(arr)==1{
		return arr[0]
	}
	a := arr[0]
	b := min(arr[1:]...)
	if a>b{
		return b
	}
	return a
}

// ------------ About Searching ---------
func getIndexOfFirstGreater(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if ref == arr[mid] {
			left = mid + 1
		} else {
			if ref > arr[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return left
}

// 简化写法
func getIndexOfFirstGreater(arr []int,ref int) int{
	left,right := 0,len(arr)-1
	for left<=right{
		mid := left+(right-left)/2
		if ref >= arr[mid]{
			left = mid + 1
		}else{
			right = mid-1
		}
	}
	return left
}

func getIndexOfFirstGreaterOrEqual(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if ref == arr[mid] {
			right = mid - 1
		} else {
			if ref > arr[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return left
}

// 简化写法
func getIndexOfFirstGreaterOrEqual(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if ref > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
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
