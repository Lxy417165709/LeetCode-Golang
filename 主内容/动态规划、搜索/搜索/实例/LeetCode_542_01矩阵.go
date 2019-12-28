package main

// 坐标结构体
type Node struct {
	x, y int
}

var dx []int // x变化向量
var dy []int // y变化向量

var hasBeenHandled map[int]bool // 用于记录节点是否被处理过

// BFS调用者
func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	hasBeenHandled = make(map[int]bool)

	/* 1. 确定搜索方向dx、dy  */
	dx = []int{0, 0, -1, 1}
	dy = []int{1, -1, 0, 0}

	/* 2. 将需要进行BFS的坐标入队 */
	m, n := len(matrix), len(matrix[0])
	queue := make([]Node, 0, m*n) // 由于go语言没有队列，这里采用切片模拟队列
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if matrix[i][t] == 0 {
				queue = append(queue, Node{i, t})
			}
		}
	}

	/* 3. 执行BFS */
	BFS(matrix, queue)
	return matrix
}

// BFS
func BFS(matrix [][]int, queue []Node) {
	level := 0
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			size --
			top := queue[0]
			queue = queue[1:]
			hashNumber := hash(top.x, top.y) // 这里采用哈希，把二维坐标哈希为一个数字

			/* 4. 判断进行BFS的坐标是否符合要求 */
			if top.x < 0 || top.y < 0 || top.x >= len(matrix) || top.y >= len(matrix[top.x]) || hasBeenHandled[hashNumber] {
				// 不满足要求时
				continue
			}

			/* 5. 对坐标进行标记，防止走回头路 */
			hasBeenHandled[hashNumber] = true
			matrix[top.x][top.y] = level

			/* 6.遍历所有方向 */
			for i := 0; i < len(dx); i++ {
				/* 7. 将下一个坐标入队 */
				nx, ny := top.x+dx[i], top.y+dy[i]
				queue = append(queue, Node{nx, ny})
			}
		}
		level++
	}
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
	1. 这题就是求1位于的层数，我们可以以0为第0层，之后每一层每一层的拓展，采用BFS的思路。
*/
