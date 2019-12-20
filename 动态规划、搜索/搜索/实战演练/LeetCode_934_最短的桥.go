package main

var dx []int
var dy []int

// DFS时判断访问坐标是否合法
func judge(A [][]int, x, y int) bool {
	if len(A) == 0 {
		return false
	}
	m, n := len(A), len(A[0])
	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}
	if A[x][y] == 0 || A[x][y] == 2 {
		return false
	}
	return true
}

// 把第一个岛屿赋值为2
func DFS(A [][]int, x, y int) {
	if !judge(A, x, y) {
		return
	}
	A[x][y] = 2
	for i := 0; i < len(dx); i++ {
		DFS(A, x+dx[i], y+dy[i])
	}
}

type Node struct {
	x, y int
}

func shortestBridge(A [][]int) int {
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, -1, 1}
	if len(A) == 0 {
		return 0
	}
	m, n := len(A), len(A[0])
findFirstIsland:
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			// 找到属于第一个岛屿的一格,将第一个岛屿全标记2
			if A[i][t] == 1 {
				DFS(A, i, t)
				break findFirstIsland
			}
		}
	}

	// 将第二个岛屿的所有点入队并标记
	hasBeenInQueue := make(map[int]bool)
	queue := make([]Node, 0)
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if A[i][t] == 1 {
				queue = append(queue, Node{i, t})
				hasBeenInQueue[hash(i, t)] = true
			}
		}
	}
	level := 0
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			size--
			top := queue[0]
			queue = queue[1:]

			// 表示已经连接到第一个岛屿了
			if A[top.x][top.y] == 2 {
				return level - 1
			}
			for i := 0; i < len(dx); i++ {
				nx, ny := top.x+dx[i], top.y+dy[i]
				hashNumber := hash(nx, ny)
				if nx < 0 || ny < 0 || nx >= m || ny >= n || hasBeenInQueue[hashNumber] {
					continue
				}
				queue = append(queue, Node{nx, ny})
				hasBeenInQueue[hashNumber] = true
			}
		}
		level++
	}
	return 0
}

func hash(x, y int) int {
	return (x << 20) | y
}

/*
	题目链接:
		https://leetcode-cn.com/problems/shortest-bridge/submissions/		最短的桥
*/

/*
	总结:
	1. 这题的思路如下:
			(1) 先找到其中一个岛屿的一块，之后通过这一块，使用DFS向外拓展，将该岛屿的每一块都标记为2。
			(2) 将另外一个岛屿的所有坐标入队，进行BFS，直到遇到第一个岛屿的部分(此时这部分的标识是2了)
			(3) BFS层数-1就是桥的长度了。
	2. 这题写得真长...
*/