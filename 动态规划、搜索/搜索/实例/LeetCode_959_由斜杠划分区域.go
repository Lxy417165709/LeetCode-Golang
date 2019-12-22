package main

import "fmt"

var size int
var dx []int
var dy []int
// 初始化参数
func initParameter() {
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}
	size = 3
}

// 判断该点是否可以走
func judge(newGrid [][]int, x, y int) bool {
	if len(newGrid) == 0 {
		return false
	}
	m, n := len(newGrid), len(newGrid[0])
	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}
	return newGrid[x][y] == 0
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

	// 接下来就是DFS找连通块了
	ans := 0
	for i := 0; i < N*size; i++ {
		for t := 0; t < N*size; t++ {
			if newGrid[i][t] == 0 {
				ans++
				DFS(newGrid, i, t)
			}
		}
	}
	return ans

}

func DFS(newGrid [][]int, x, y int) {
	if judge(newGrid, x, y) == false {
		return
	}
	newGrid[x][y] = 1
	for i := 0; i < len(dx); i++ {
		DFS(newGrid, x+dx[i], y+dy[i])
	}
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
	1. 这题也是求连通块的题，但是很妙，需要利用像素化的思维解决。
			做法如下:  	假如我们得到的矩阵是 ["//"],
											 ["\\"]，
						那么我们可以用3*3的矩阵代替每一个小方格，
						得到另外一个矩阵 	 [001001]
											 [010010]
										     [100100]
											 [100100]
											 [010010]
										     [001001]
						之后再求这个矩阵的连通块就可以了。 (只允许上下左右4个方向移动)
						注意不能分解成2*2的，因为会有这种情况 ["//"],
															  ["//"],
						变为2*2的矩阵后是		[0101]
										        [1010]
												[0101]
										        [1010]
						此时上下左右移动得到的连通块数多于实际的连通块数。(因为有些地方本来可以到达的，但是被阻断了)
	2. 这题类似岛屿问题，就是求连通块的数目。
	3. 这题还能使用并查集解。
*/
