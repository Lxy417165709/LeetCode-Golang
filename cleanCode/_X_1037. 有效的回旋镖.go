package main

func isBoomerang(points [][]int) bool {
	v1 := NewVector(points[0][0]-points[1][0],points[0][1]-points[1][1])
	v2 := NewVector(points[0][0]-points[2][0],points[0][1]-points[2][1])
	return !isParallel(v1,v2)
}



// ----------- Vector ----------
type Vector struct {
	X int
	Y int
}

func NewVector(x, y int) *Vector {
	return &Vector{x, y}
}


func isParallel(v1, v2 *Vector) bool{
	return getCross(v1,v2)==0
}

func getCross(v1, v2 *Vector) int {
	return v1.X*v2.Y - v1.Y*v2.X
}
