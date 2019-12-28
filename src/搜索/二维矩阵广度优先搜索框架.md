
@[toc]

# 二维矩阵广度优先搜索 (TODO)
## 1. 框架
```go
// 坐标结构体
type Node struct {
	x, y int
}

var dx []int    // x变化向量
var dy []int    // y变化向量

// BFS调用者
func BFSCaller() {
	/* 1. 确定搜索方向dx、dy  */
	/* 2. 将需要进行BFS的坐标入队 */
	/* 3. 调用BFS函数 */
}

// BFS
func BFS(){
    level := 0  // level就是此时遍历的层次
    for len(queue) != 0 {
        size := len(queue)
        for size != 0{
            size--
            top := queue[0]
            queue = queue[1:]
            /* 4. 判断进行BFS的坐标是否符合要求 */
            /* 5. 对坐标进行标记，防止走回头路 */
            for /* 6.遍历所有方向 */ {
                /* 7. 将下一个坐标入队 */
            }
        }
        level++
    }
}
```
## 2. 实例
### 2.1 岛屿数量
[传送门](https://leetcode-cn.com/problems/number-of-islands/)
```go
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
```
### 2.2 01矩阵
[传送门](https://leetcode-cn.com/problems/01-matrix/submissions/)
```go
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
	queue := make([]Node, 0, m*n)
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

// 将坐标(x, y)哈希为一个数字，该数字用于唯一的标识坐标(x, y)
func hash(x, y int) int {
	// 具体的哈希取决于题目的x,y的取值范围
	return (x << 20) | y
}
```

### 2.3 腐烂的橘子
[传送门](https://leetcode-cn.com/problems/rotting-oranges/)
```go
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
```
## 3. 注意
- [ ] 以上只是BFS的大体框架，实际做题可能需要进行一些变化。
## 4. 练习题
- [ ] [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)
- [ ] [542. 01 矩阵](https://leetcode-cn.com/problems/01-matrix/submissions/)
- [ ] [934. 最短的桥](https://leetcode-cn.com/problems/shortest-bridge/)
- [ ] [944. 腐烂的橘子](https://leetcode-cn.com/problems/rotting-oranges/)
