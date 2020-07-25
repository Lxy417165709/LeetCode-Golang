package 动态规划

// ---------------------------- 动态规划，非原地 ----------------------------
func countSquares(matrix [][]int) int {
	if len(matrix)==0{
		return 0
	}
	rows, cols := getRowsAndCols(matrix)
	maxSquareHeight := get2DSlice(rows, cols)
	maxSquareHeight[0][0] = matrix[0][0]
	for i:=1;i<rows;i++{
		maxSquareHeight[i][0]=matrix[i][0]
	}
	for t:=1;t<cols;t++{
		maxSquareHeight[0][t]=matrix[0][t]
	}
	for i:=1;i<rows;i++{
		for t:=1;t<cols;t++{
			if matrix[i][t]==1{
				maxSquareHeight[i][t] = min(
					maxSquareHeight[i-1][t-1],
					maxSquareHeight[i][t-1],
					maxSquareHeight[i-1][t],
				)+1
			}
		}
	}
	countOfSquare := 0
	for i:=0;i<rows;i++{
		for t:=0;t<cols;t++{
			countOfSquare+=maxSquareHeight[i][t]
		}
	}
	return countOfSquare
}


func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

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
	a,b := arr[0],min(arr[1:]...)
	if a>b{
		return b
	}
	return a
}



// ---------------------------- 动态规划，原地 ----------------------------
func countSquares(matrix [][]int) int {
	rows,cols := getRowsAndCols(matrix)
	for i:=1;i<rows;i++{
		for t:=1;t<cols;t++{
			if matrix[i][t]==1{
				matrix[i][t] = min(
					matrix[i-1][t-1],
					matrix[i][t-1],
					matrix[i-1][t],
				)+1
			}
		}
	}
	countOfSquare := 0
	for i:=0;i<rows;i++{
		for t:=0;t<cols;t++{
			countOfSquare+=matrix[i][t]
		}
	}
	return countOfSquare
}


func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

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
	a,b := arr[0],min(arr[1:]...)
	if a>b{
		return b
	}
	return a
}



/*
	总结:
		1. 这题和 _1504. 统计全 1 子矩形_ 有些类似，但是一个求的是矩形数量、一个是正方形数量，二者解法天差地别。
*/
