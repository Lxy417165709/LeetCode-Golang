package main

// 坐标结构体
type Node struct {
	x, y int
}

var dx []int
var dy []int

// BFS调用者
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	/* 1. 确定搜索方向dx、dy  */
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}

	/* 2. 将需要进行BFS的坐标入队 */
	m, n := len(grid), len(grid[0])
	ans := 0
	queue := make([]Node, 0)
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if grid[i][t] == '1' {
				ans++
				queue = append(queue, Node{i, t})

				/* 3. 执行BFS */
				BFS(grid, queue)
			}
		}
	}
	return ans
}

// BFS
func BFS(grid [][]byte, queue []Node) {
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			size --
			top := queue[0]
			queue = queue[1:]

			/* 4. 判断进行BFS的坐标是否符合要求 */
			if top.x < 0 || top.y < 0 || top.x >= len(grid) || top.y >= len(grid[top.x]) || grid[top.x][top.y] == '0' {
				continue
			}

			/* 5. 对坐标进行标记，防止走回头路 */
			grid[top.x][top.y] = '0'

			/* 6.遍历所有方向 */
			for i := 0; i < len(dx); i++ {
				/* 7. 将下一个坐标入队 */
				nx, ny := top.x+dx[i], top.y+dy[i]
				queue = append(queue, Node{nx, ny})
			}
		}
	}
}

/*
	题目链接:
		https://leetcode-cn.com/problems/number-of-islands/		求岛屿数量
*/

/*
	总结
	1. 搜索一般模板:
			(1) 确定寻找方向(一般是上下左右)	---->    我们可以采用dx、dy数组
			(2) 防止走回头路					---->	 我们可以修改原矩阵的信息、或者创建一个isVisited映射
			(3) 判断该点是否可以走				---->    可以写一个judge函数，如果该点走过或者越界则返回false
	2. 深搜一般采用递归实现，而广搜一般采用队列 + 迭代实现。
*/
