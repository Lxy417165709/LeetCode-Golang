package main


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
