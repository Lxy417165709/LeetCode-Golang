package main

// ------------------- 解法1 (面向过程，采用2个哈希表存储 "0" 的信息) --------------------
func setZeroes(matrix [][]int) {
	isExistedZeroInRow := make(map[int]bool)
	isExistedZeroInColumn := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[i]); t++ {
			if matrix[i][t] == 0 {
				isExistedZeroInRow[i] = true
				isExistedZeroInColumn[t] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[i]); t++ {
			if isExistedZeroInRow[i] || isExistedZeroInColumn[t] {
				matrix[i][t] = 0
			}
		}
	}
}

// ------------------- 解法2 (面向对象，采用第一行、第一列存储 "0" 的信息) --------------------
func setZeroes(matrix [][]int) {
	m := NewMatrix(matrix)
	m.SetZeroes()
}

// ------------------ Matrix ---------------
type Matrix struct {
	data                       [][]int
	isExistedZeroInFirstRow    bool
	isExistedZeroInFirstColumn bool
	firstRow                   int
	firstColumn                int
}

func NewMatrix(data [][]int) *Matrix {
	return &Matrix{data: data, firstRow: 0, firstColumn: 0}
}

func (m *Matrix) SetZeroes() {
	for row := m.firstRow; row < m.getRows(); row++ {
		for column := m.firstColumn; column < m.getColumns(); column++ {
			if m.data[row][column] == 0 {
				m.storeZeroInformation(row, column)
			}
		}
	}
	for row := m.firstRow + 1; row < m.getRows(); row++ {
		if m.isExistedZeroInRow(row) {
			m.clearMatrixRow(row)
		}
	}
	for column := m.firstColumn + 1; column < m.getColumns(); column++ {
		if m.isExistedZeroInColumn(column) {
			m.clearMatrixColumn(column)
		}
	}
	if m.isExistedZeroInRow(m.firstRow) {
		m.clearMatrixRow(m.firstRow)
	}
	if m.isExistedZeroInColumn(m.firstColumn) {
		m.clearMatrixColumn(m.firstColumn)
	}
}

func (m *Matrix) storeZeroInformation(row, column int) {
	m.data[row][m.firstColumn] = 0
	m.data[m.firstRow][column] = 0

	if row == m.firstRow {
		m.isExistedZeroInFirstRow = true
	}
	if column == m.firstColumn {
		m.isExistedZeroInFirstColumn = true
	}
}

func (m *Matrix) isExistedZeroInRow(row int) bool {
	if row == m.firstRow {
		return m.isExistedZeroInFirstRow
	}
	return m.data[row][0] == 0
}

func (m *Matrix) isExistedZeroInColumn(column int) bool {
	if column == m.firstColumn {
		return m.isExistedZeroInFirstColumn
	}
	return m.data[0][column] == 0
}

func (m *Matrix) clearMatrixRow(row int) {
	for column := m.firstColumn; column < m.getColumns(); column++ {
		m.data[row][column] = 0
	}
}

func (m *Matrix) clearMatrixColumn(column int) {
	for row := m.firstRow; row < m.getRows(); row++ {
		m.data[row][column] = 0
	}
}

func (m *Matrix) getRows() int {
	return len(m.data)
}

func (m *Matrix) getColumns() int {
	if m.getRows() == 0 {
		return 0
	}
	return len(m.data[0])
}

/*
	总结
	1. 这题解题方法有：
			O(1) 空间: 用原矩阵的第一行、第一列来存储 "0" 的信息，之后第一行、第一列特殊处理。
			O(1) 空间 且 存在不可能出现在矩阵中的数A: 用A进行标记清零的位置。
			O(m+n)空间: 采用2个哈希表来存储 "0" 的信息。
			O(m*n)空间: 开辟新矩阵。

*/
