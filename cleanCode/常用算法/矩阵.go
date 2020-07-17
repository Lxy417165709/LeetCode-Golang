package main


func getRowsAndCols(matrix [][]int) (int,int){
	if len(matrix)==0{
		return 0,0
	}
	return len(matrix),len(matrix[0])
}
