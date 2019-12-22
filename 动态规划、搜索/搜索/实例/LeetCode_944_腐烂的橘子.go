package main

type Node struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dx := []int{0, 0, -1, 1}
	dy := []int{1, -1, 0, 0}

	// 先将所有腐烂的橘子坐标入队
	hasBeenInQueue := make(map[int]bool)
	queue := make([]Node, 0)
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if grid[i][t] == 2 {
				hasBeenInQueue[hash(i, t)] = true
				queue = append(queue, Node{i, t})
			}
		}
	}

	// 开始BFS，即开始腐烂感染。
	level := 0
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			size--
			top := queue[0]
			queue = queue[1:]
			for i := 0; i < len(dx); i++ {
				nx, ny := top.x+dx[i], top.y+dy[i]
				if nx < 0 || ny < 0 || nx >= m || ny >= n || hasBeenInQueue[hash(nx, ny)] || grid[nx][ny] == 0 {
					continue
				}
				grid[nx][ny] = 2
				queue = append(queue, Node{nx, ny})
				hasBeenInQueue[hash(nx, ny)] = true
			}
		}
		level++
	}

	hasRotOrrangeFlag := false // 判断是否有腐烂的橘子
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if grid[i][t] == 1 {
				return -1 // BFS还有没腐烂的橘子，按照题目要求返回 -1
			}
			if grid[i][t] == 2 {
				hasRotOrrangeFlag = true // 表示有腐烂的橘子，那最后就可以返回level - 1
			}
		}
	}
	if hasRotOrrangeFlag {
		return level - 1 // 这里表示有腐烂的橘子，按照题目要求返回传染的分钟数，在这就是level - 1
	} else {
		return 0 // 这里表示没有腐烂的橘子，按照题目要求返回 0
	}
}

func hash(x, y int) int {
	return (x << 10) | y
}

/*
	题目链接:
		https://leetcode-cn.com/problems/rotting-oranges/		腐烂的橘子
*/

/*
	总结
	1. 这题就是简单的BFS，不过有个小坑就是: 在BFS结束后，要看看橘子的状态再给出答案。
*/
