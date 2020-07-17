package main

func movingCount(m int, n int, k int) int {
	board := NewBoard()
	return board.GetCountOfReachingCoordinate(m, n, k, NewCoordinate(0, 0), []*Vector{
		NewVector(0, 1),
		NewVector(0, -1),
		NewVector(1, 0),
		NewVector(-1, 0),
	})
}

// ----------- Board ----------
type Board struct {
	rows                         int
	columns                      int
	maxValidDigitSumOfCoordinate int
	hasVisit                     map[Coordinate]bool
	moveVectors                  []*Vector
}

func NewBoard() *Board {
	return &Board{}
}

func (bd *Board) GetCountOfReachingCoordinate(rows int, columns int, maxValidDigitSumOfCoordinate int, startCoordinate *Coordinate, moveVectors []*Vector) int {
	bd.rows = rows
	bd.columns = columns
	bd.maxValidDigitSumOfCoordinate = maxValidDigitSumOfCoordinate
	bd.moveVectors = moveVectors
	bd.hasVisit = make(map[Coordinate]bool)
	return bd.getCountOfReachingCoordinateThere(startCoordinate)
}

func (bd *Board) getCountOfReachingCoordinateThere(nowCoordinate *Coordinate) int {
	if bd.isCoordinateOutOfBound(nowCoordinate) {
		return 0
	}
	if bd.isDigitSumOfCoordinateValid(nowCoordinate) {
		return 0
	}
	if bd.hasCoordinateVisit(nowCoordinate) {
		return 0
	}
	countOfReachingCoordinate := 1
	bd.markCoordinateHasVisit(nowCoordinate)
	for i := 0; i < len(bd.moveVectors); i++ {
		nextCoordinate := getNextCoordinate(nowCoordinate, bd.moveVectors[i])
		countOfReachingCoordinate += bd.getCountOfReachingCoordinateThere(nextCoordinate)
	}
	return countOfReachingCoordinate
}

func (bd *Board) isCoordinateOutOfBound(c *Coordinate) bool {
	return c.X < 0 || c.X >= bd.columns || c.Y < 0 || c.Y >= bd.rows
}

func (bd *Board) isDigitSumOfCoordinateValid(c *Coordinate) bool {
	return getDigitSumOfCoordinate(c) > bd.maxValidDigitSumOfCoordinate
}

func (bd *Board) markCoordinateHasVisit(c *Coordinate) {
	bd.hasVisit[*c] = true
}

func (bd *Board) hasCoordinateVisit(c *Coordinate) bool {
	return bd.hasVisit[*c]
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

// ----------- Utils ----------
func getNextCoordinate(c *Coordinate, v *Vector) *Coordinate {
	return NewCoordinate(c.X+v.X, c.Y+v.Y)
}

func getDigitSumOfCoordinate(c *Coordinate) int {
	return getDigitSumOfNumber(c.X) + getDigitSumOfNumber(c.Y)
}

func getDigitSumOfNumber(num int) int {
	digitSum := 0
	for num != 0 {
		digitSum += num % 10
		num /= 10
	}
	return digitSum
}

/*
	题目链接： https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/submissions/
*/
