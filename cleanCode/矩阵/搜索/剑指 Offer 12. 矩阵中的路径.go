package 搜索

func exist(charMatrix [][]byte, word string) bool {
	board := NewBoard(charMatrix, []*Vector{
		NewVector(1, 0),
		NewVector(-1, 0),
		NewVector(0, -1),
		NewVector(0, 1),
	})
	return board.IsExistRoadToFormSpecificWord(word)
}

// ----------- Board ----------
type Board struct {
	charMatrix [][]byte
	hasVisit   map[Coordinate]bool
	vectors    []*Vector
}

func NewBoard(charMatrix [][]byte, vectors []*Vector) *Board {
	return &Board{
		charMatrix: charMatrix,
		vectors:    vectors,
	}
}

func (bd *Board) IsExistRoadToFormSpecificWord(specificWord string) bool {
	bd.initHasVisit()
	for i := 0; i < bd.getCharMatrixRows(); i++ {
		for j := 0; j < bd.getCharMatrixColumns(); j++ {
			if bd.isExistRoadToFormSpecificWordThere(NewCoordinate(i, j), specificWord) {
				return true
			}
		}
	}
	return false
}

func (bd *Board) initHasVisit() {
	bd.hasVisit = make(map[Coordinate]bool)
}

func (bd *Board) isExistRoadToFormSpecificWordThere(c *Coordinate, specificWord string) bool {
	if isWordBlank(specificWord) {
		return true
	}
	if bd.isCoordinateNotNeedToVisit(c) || bd.isCharOfCoordinateNotFix(c, specificWord[0]) {
		return false
	}
	bd.markCoordinateHasVisit(c)
	for i := 0; i < bd.getVectorRows(); i++ {
		nextX := c.X + bd.vectors[i].X
		nextY := c.Y + bd.vectors[i].Y
		nextCoordinate := NewCoordinate(nextX, nextY)
		nextWord := specificWord[1:]
		if bd.isExistRoadToFormSpecificWordThere(nextCoordinate, nextWord) {
			return true
		}
	}
	bd.clearCoordinateHasVisit(c)
	return false
}

func (bd *Board) isCoordinateNotNeedToVisit(c *Coordinate) bool {
	return bd.isCoordinateInvalid(c) || bd.isCoordinateHasVisit(c)
}

func (bd *Board) isCoordinateInvalid(c *Coordinate) bool {
	return c.X < 0 || c.X >= bd.getCharMatrixRows() || c.Y < 0 || c.Y >= bd.getCharMatrixColumns()
}

func (bd *Board) isCoordinateHasVisit(c *Coordinate) bool {
	return bd.hasVisit[*c]
}

func (bd *Board) isCharOfCoordinateNotFix(c *Coordinate, comparedChar byte) bool {
	return bd.charMatrix[c.X][c.Y] != comparedChar
}

func (bd *Board) markCoordinateHasVisit(c *Coordinate) {
	bd.hasVisit[*c] = true
}

func (bd *Board) clearCoordinateHasVisit(c *Coordinate) {
	bd.hasVisit[*c] = false
}

func (bd *Board) getCharMatrixRows() int {
	return len(bd.charMatrix)
}

func (bd *Board) getCharMatrixColumns() int {
	return len(bd.charMatrix[0])
}

func (bd *Board) getVectorRows() int {
	return len(bd.vectors)
}

func isWordBlank(word string) bool {
	return word == ""
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

/*
	题目链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/submissions/

	提升代码可读性策略：
		1. 分包
		2. 将 (x, y) 封装为 Coordinate 结构体
		3. 将 vector 封装为结构体
*/
