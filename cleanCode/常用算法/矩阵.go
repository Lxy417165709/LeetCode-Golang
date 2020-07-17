package main


func getRowsAndCols(matrix [][]int) (int,int){
	if len(matrix)==0{
		return 0,0
	}
	return len(matrix),len(matrix[0])
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}
