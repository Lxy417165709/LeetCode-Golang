package 数组

const INF = 100000000000
var hasExisted map[Coordinate]bool


func minAreaRect(points [][]int) int {
	coordinates := make([]*Coordinate, 0)
	minArea := INF
	hasExisted = make(map[Coordinate]bool)

	for _, point := range points {
		coordinate := NewCoordinate(point[0], point[1])
		hasExisted[*coordinate] = true
		coordinates = append(coordinates, coordinate)
	}

	for i := 0; i < len(coordinates); i++ {
		for t := i + 1; t < len(coordinates); t++ {
			ci, ct := coordinates[i], coordinates[t]
			if canBeMatrix(ci, ct) && haveOtherTwoCoordinatesBothExisted(ci, ct) {
				minArea = min(minArea, getSquareOfMatrix(ci, ct))
			}
		}
	}
	if minArea == INF {
		return 0
	}
	return minArea
}

func canBeMatrix(c1 *Coordinate, c2 *Coordinate) bool {
	if c1.X == c2.X || c1.Y == c2.Y {
		return false
	}
	return true
}

func haveOtherTwoCoordinatesBothExisted(c1 *Coordinate, c2 *Coordinate) bool {
	coordinates := getOtherTwoCoordinates(c1, c2)
	return hasExisted[*coordinates[0]] && hasExisted[*coordinates[1]]
}

func getSquareOfMatrix(c1 *Coordinate, c2 *Coordinate) int {
	return abs(c1.X-c2.X) * abs(c1.Y-c2.Y)
}

func getOtherTwoCoordinates(c1 *Coordinate, c2 *Coordinate) []*Coordinate {
	return []*Coordinate{
		NewCoordinate(c1.X, c2.Y),
		NewCoordinate(c2.X, c1.Y),
	}
}


// ----------- Coordinate ----------
type Coordinate struct {
	X int
	Y int
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{x, y}
}

// ----------- Utils ----------
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*
	题目: https://leetcode-cn.com/problems/minimum-area-rectangle-ii/
	总结:
		1. 我觉得其实可以通过2个点就推出另外2个点的，即使是在长、宽不平行与X、Y轴的情况。 (可以通过向量运算，但我没算出)
*/
