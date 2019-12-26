package main

var dx []int    // x变化向量
var dy []int    // y变化向量

// 这个就是DFSCaller
func numIslands(grid [][]byte) int {
	/* 1. 确定搜索方向 */
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}

	/* 2. 对坐标进行DFS搜索 (这里开始调用DFS函数) */
	ans := 0
	for i := 0; i < len(grid); i++ {
		for t := 0; t < len(grid[i]); t++ {
			if grid[i][t] == '1'{
				ans++
			}
			DFS(grid, i, t)
		}
	}

	return ans
}

// DFS递归搜索
func DFS(grid [][]byte, x, y int) {
	/* 3. 判断进行DFS的坐标是否符合要求 */
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) || grid[x][y] == '0' {
		// 不满足时
		return
	}
	// 满足时
	/* 4. 对坐标进行标记，防止走回头路 */
	grid[x][y] = '0' // 将陆地变为海洋，防止走回头路

	/* 5.遍历所有方向 */
	for i := 0; i < len(dx); i++ {

		/* 6. 对下一个坐标进行DFS搜索 */
		nx, ny := x+dx[i], y+dy[i]
		DFS(grid, nx, ny)
	}
}
/*
	题目链接:
		https://leetcode-cn.com/problems/number-of-islands/							岛屿数量
		https://leetcode-cn.com/problems/number-of-closed-islands/submissions/		统计封闭岛屿的数目
		https://leetcode-cn.com/problems/number-of-enclaves/						飞地的数量
*/


/*
	总结
	1. 搜索一般模板:
			(1) 确定寻找方向(一般是上下左右)	---->    我们可以采用dx、dy数组
			(2) 防止走回头路					---->	 我们可以修改原矩阵的信息、或者创建一个isVisited映射
			(3) 判断该点是否可以走				---->    可以写一个judge函数，如果该点走过或者越界则返回false
*/
