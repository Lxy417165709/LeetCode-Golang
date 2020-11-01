package main

// 坐标结构体
type Node struct {
	x, y int
}

var dx []int	// x变化向量
var dy []int	// y变化向量

var hasBeenHandled map[int]bool		// 用于

// BFS调用者
func orangesRotting(grid [][]int) int {

	/* 1. 确定搜索方向dx、dy  */
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}

	hasBeenHandled = make(map[int]bool)
	goodOrangeCount := 0

	/* 2. 将需要进行BFS的坐标入队 */
	queue := make([]Node, 0)
	for i := 0; i < len(grid); i++ {
		for t := 0; t < len(grid[i]); t++ {
			if grid[i][t] == 2 {
				queue = append(queue, Node{i, t})
			}
			if grid[i][t] == 1 {
				goodOrangeCount++
			}
		}
	}

	/* 3. 调用BFS函数 */
	return BFS(grid, queue, goodOrangeCount)
}

// BFS
func BFS(grid [][]int, queue []Node, goodOrangeCount int) int {
	if goodOrangeCount == 0 {
		return 0
	}
	minutes := 0 // 相当于层数
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			size--
			top := queue[0]
			queue = queue[1:]

			/* 4. 判断进行BFS的坐标是否符合要求 */
			x, y := top.x, top.y
			if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
				continue
			}
			if hasBeenHandled[hash(x, y)] || grid[x][y] == 0 {
				continue
			}

			/* 5. 对坐标进行标记，防止走回头路 */
			hasBeenHandled[hash(x, y)] = true
			if grid[x][y] == 1 {
				goodOrangeCount--
				if goodOrangeCount == 0 {
					return minutes
				}
			}
			grid[x][y] = 2

			/* 6.遍历所有方向 */
			for i := 0; i < len(dx); i++ {
				/* 7. 将下一个坐标入队 */
				nx, ny := x+dx[i], y+dy[i]
				queue = append(queue, Node{nx, ny})
			}
		}
		minutes++
	}
	return -1
}

// 将坐标(x, y)哈希为一个数字，该数字用于唯一的标识坐标(x, y)
func hash(x, y int) int {
	// 具体的哈希取决于题目的x,y的取值范围
	return (x << 10) | y
}

/*
	题目链接:
		https://leetcode-cn.com/problems/rotting-oranges/		腐烂的橘子
*/

