package main



func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
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
