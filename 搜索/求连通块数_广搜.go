package main

var dx [4]int
var dy [4]int

// 坐标结构体
type Node struct {
	x, y int
}

// BFS 解决矩阵连通块问题 (原题是求岛屿数量)
func numIslands(grid [][]byte) int {

	// 确定搜索方向
	dx = [4]int{0, 0, 1, -1}
	dy = [4]int{1, -1, 0, 0}

	ans := 0
	if len(grid) == 0 {
		return ans
	}
	m, n := len(grid), len(grid[0])

	queue := make([]Node, 0)
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if judge(grid, i, t) {
				ans++
				queue = append(queue, Node{i, t})

				// BFS
				for len(queue) != 0 {
					top := queue[0]
					queue = queue[1:]
					if judge(grid, top.x, top.y) == false {
						continue
					}
					grid[top.x][top.y] = '0'
					for j := 0; j < len(dx); j++ {
						queue = append(queue, Node{top.x + dx[j], top.y + dy[j]})
					}
				}
			}

		}
	}
	return ans
}

// 判断坐标合法性
func judge(grid [][]byte, x, y int) bool {
	if len(grid) == 0 {
		return false
	}
	m, n := len(grid), len(grid[0])

	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}

	if grid[x][y] == '0' {
		return false
	}
	return true
}

/*
	题目链接: https://leetcode-cn.com/problems/number-of-islands/
*/

/*
	总结
	1. 搜索一般模板:
			(1) 确定寻找方向(一般是上下左右)	---->    我们可以采用dx、dy数组
			(2) 防止走回头路					---->	 我们可以修改原矩阵的信息、或者创建一个isVisited映射
			(3) 判断该点是否可以走				---->    可以写一个judge函数，如果该点走过或者越界则返回false
	2. 深搜一般采用递归实现，而广搜一般采用队列 + 迭代实现。
*/
