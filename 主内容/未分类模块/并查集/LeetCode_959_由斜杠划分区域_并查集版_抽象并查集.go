package main

var size int
var dx []int
var dy []int
// 初始化参数
func initParameter() {
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}
	size = 3
}

var father []int
// 初始化并查集
func initUFS(length int) {
	father = make([]int, length)
	for i := 0; i < length; i++ {
		father[i] = i
	}
}

// 合并并查集
func unionUFS(a, b int) {
	father1, father2 := findFather(a), findFather(b)
	father[father1] = father2
}

// 找爸爸
func findFather(a int) int {
	if a == father[a] {
		return a
	}
	father[a] = findFather(father[a])
	return father[a]
}

func regionsBySlashes(grid []string) int {
	if len(grid) == 0 {
		return 0
	}
	initParameter()

	N := len(grid)
	newGrid := make([][]int, N*size)
	for i := 0; i < N*size; i++ {
		newGrid[i] = make([]int, N*size)
	}
	// 这里是为了把'/'和'\'像素化，像素化到newGrid中
	for i := 0; i < N; i++ {
		for t := 0; t < N; t++ {
			if grid[i][t] == '/' {
				toPixel(newGrid, i, t, 1)
			}
			if grid[i][t] == '\\' {
				toPixel(newGrid, i, t, 2)
			}
		}
	}
	m, n := N*size, N*size
	initUFS(m * n) // 之所以这样初始化，是因为此时newGrid有m*n个点

	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if newGrid[i][t] == 1 {
				continue
			}
			for u := 0; u < len(dx); u++ {
				nx, ny := i+dx[u], t+dy[u]
				if isNeedingToUnion(newGrid, nx, ny) {
					unionUFS(i*n+t, nx*n+ny)
				}
			}
		}
	}
	// 下面是获取连通分块
	ans := 0
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if newGrid[i][t] == 1 {
				continue
			}
			if findFather(i*n+t) == i*n+t {
				ans++
			}
		}
	}
	return ans
}
func isNeedingToUnion(newGrid [][]int, x, y int) bool {
	if len(newGrid) == 0 {
		return false
	}
	m, n := len(newGrid), len(newGrid[0])
	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}
	return newGrid[x][y] == 0
}

// flag == 1 表示对'/'像素化
// flag == 2 表示对'\'像素化
func toPixel(newGrid [][]int, x, y, flag int) {
	if flag == 1 {
		nx := x * size
		ny := y*size + size - 1
		for i := 0; i < size; i++ {
			newGrid[nx+i][ny-i] = 1
		}
	} else {
		nx := x * size
		ny := y * size
		for i := 0; i < size; i++ {
			newGrid[nx+i][ny+i] = 1
		}
	}
}

/*
	总结
	1. 这题我采用了并查集解，但是写的太乱了...
	2. 思路是构造出01矩阵，再在这个矩阵中用并查集求连通分量..
*/
