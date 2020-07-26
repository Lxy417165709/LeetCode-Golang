package 未归类

var diagonalsOwningSameHash map[float64][]*Diagonal

const INF = 100000000000000.0

func minAreaFreeRect(points [][]int) float64 {
	diagonalsOwningSameHash = make(map[float64][]*Diagonal)
	coordinates := make([]*Coordinate, 0)
	diagonals := make([]*Diagonal, 0)
	minArea := INF
	for _, point := range points {
		coordinates = append(coordinates, NewCoordinate(float64(point[0]), float64(point[1])))
	}
	for i := 0; i < len(coordinates); i++ {
		for t := i + 1; t < len(coordinates); t++ {
			ci, ct := coordinates[i], coordinates[t]
			diagonals = append(diagonals, NewDiagonal(ci, ct))
		}
	}
	for _, diagonal := range diagonals {
		sameHashDiagonals := diagonalsOwningSameHash[getHash(diagonal)]
		diagonalsOwningSameHash[getHash(diagonal)] = append(diagonalsOwningSameHash[getHash(diagonal)], diagonal)
		for _, another := range sameHashDiagonals {
			minArea = min(minArea, getMatrixSquare(diagonal, another))
		}
	}
	if minArea == INF {
		return 0
	}
	return minArea
}

func getHash(d *Diagonal) float64 {
	const maxX = 40000
	const maxY = 40000
	mc := getMiddleCoordinate(d.Coordinate1, d.Coordinate2)
	dis := getDistance(d.Coordinate1, d.Coordinate2)
	hash := mc.X + mc.Y*(maxX+1) + dis*(maxY*maxX+1)
	return hash
}

func getMiddleCoordinate(c1, c2 *Coordinate) *Coordinate {
	return NewCoordinate((c1.X+c2.X)/2, (c1.Y+c2.Y)/2)
}

func getMatrixSquare(d1, d2 *Diagonal) float64 {
	return getDistance(d1.Coordinate1, d2.Coordinate1) * getDistance(d1.Coordinate1, d2.Coordinate2)
}

func getDistance(c1, c2 *Coordinate) float64 {
	return math.Sqrt((c1.X-c2.X)*(c1.X-c2.X) + (c1.Y-c2.Y)*(c1.Y-c2.Y))
}



// ----------- Diagonal ----------
type Diagonal struct {
	Coordinate1 *Coordinate
	Coordinate2 *Coordinate
}

func NewDiagonal(c1, c2 *Coordinate) *Diagonal {
	return &Diagonal{c1, c2}
}

// ----------- Coordinate ----------
type Coordinate struct {
	X float64
	Y float64
}

func NewCoordinate(x, y float64) *Coordinate {
	return &Coordinate{x, y}
}

// ----------- Utils ----------
func min(arr ...float64) float64 {
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

/*

	题目链接: https://leetcode-cn.com/problems/minimum-area-rectangle-ii/
	总结:
		1. 这题采用 对角线 来唯一标识矩形。将属于同一矩形的对角线找出，就能得出矩形的面积了。
*/
