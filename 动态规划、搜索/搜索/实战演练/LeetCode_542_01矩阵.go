package main

type Node struct {
	x, y int
}
/*
	给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
	两个相邻元素间的距离为 1 。
*/
func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	var dx = []int{0, 0, -1, 1}
	var dy = []int{1, -1, 0, 0}
	hasBeenInQueue := make(map[int]bool)
	m, n := len(matrix), len(matrix[0])
	queue := make([]Node, 0, m*n)
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if matrix[i][t] == 0 {
				queue = append(queue, Node{i, t})
				hasBeenInQueue[hash(i, t)] = true
			}
		}
	}
	level := 0
	for len(queue) != 0 {
		size := len(queue)
		level++
		for size != 0 {
			top := queue[0]
			queue = queue[1:]
			size --
			for i := 0; i < len(dx); i++ {
				nx, ny := top.x+dx[i], top.y+dy[i]
				hashNumber := hash(nx, ny)
				if nx < 0 || ny < 0 || nx >= m || ny >= n || hasBeenInQueue[hashNumber] {
					continue
				}
				queue = append(queue, Node{nx, ny})
				matrix[nx][ny] = level
				hasBeenInQueue[hashNumber] = true
			}
		}
	}
	return matrix
}

func hash(x, y int) int {
	return (x << 20) | y
}
/*
	题目链接:
		https://leetcode-cn.com/problems/01-matrix/			01矩阵
*/
/*
	总结
	1. 这题就是求1位于的层数，我们可以以0位第0层，之后每一层每一层的拓展，采用BFS的思路。
*/