[toc]

## 二维矩阵广度优先搜索 (TODO)
### 1. 框架
```go
// 坐标结构体
type Node struct {
	x, y int
}

// 这个BFS框架无法知道此时在第几层
func BFS(){
    /* 1. 确定搜索方向 */
    /* 2. 将满足要求的坐标节点入队 */
    for len(queue) != 0 {
        // 取出队首元素并删除
        top := queue[0]     
        queue = queue[1:]
        for /* 3.遍历所有方向 */ {
            if /* 4. 判断新坐标合法性及是否已入队 */ {
                // 不满足要求时
                continue
            }
            // 满足要求时
            /* 5. 入队并标记 */
        }
    }
}

// 这个BFS框架可以知道此时在第几层
func BFS(){
    /* 1. 确定搜索方向 */
    /* 2. 将满足要求的坐标节点入队 */
    level := 0  // level就是此时遍历的层次
    for len(queue) != 0 {
        size := len(queue)
        for size != 0{
            size--
            // 取出队首元素并删除
            top := queue[0]
            queue = queue[1:]
            for /* 3. 遍历所有方向 */ {
                if /* 4. 判断新坐标合法性及是否已入队 */ {
                    // 不满足要求时
                    continue
                }
                // 满足要求时
                /* 5. 入队并标记 */
            }
        }
        level++
    }
}
```
### 2. 实例
```go
/*
	给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
	两个相邻元素间的距离为 1 。
	
	题目链接: https://leetcode-cn.com/problems/01-matrix/	01矩阵
*/

type Node struct {
	x, y int
}

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	
	/* 1. 确定搜索方向 */
	var dx = []int{0, 0, -1, 1}
	var dy = []int{1, -1, 0, 0}
	
	/* 2. 将满足要求的坐标节点入队 */
	m, n := len(matrix), len(matrix[0])
	queue := make([]Node, 0, m*n)	// 由于go语言没有队列，这里采用切片模拟队列
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if matrix[i][t] == 0 {
				queue = append(queue, Node{i, t})
				hasBeenInQueue[hash(i, t)] = true
			}
		}
	}
	
	hasBeenInQueue := make(map[int]bool)	// 用于记录节点是否入过队列
	
	level := 0
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			size --
			top := queue[0]
			queue = queue[1:]
			/* 3.遍历所有方向 */
			for i := 0; i < len(dx); i++ {
				nx, ny := top.x+dx[i], top.y+dy[i]
				hashNumber := hash(nx, ny)	// 这里采用哈希，把二维坐标哈希为一个数字
				
				/* 4. 判断新坐标合法性及是否已入队 */
				if nx < 0 || ny < 0 || nx >= m || ny >= n || hasBeenInQueue[hashNumber] {
					continue
				}
				
				/* 5. 入队并标记 */
				matrix[nx][ny] = level + 1
				queue = append(queue, Node{nx, ny})
				hasBeenInQueue[hashNumber] = true
			}
		}
		level++
	}
	return matrix
}

func hash(x, y int) int {
	return (x << 20) | y
}
```

### 3. 练习题
- [ ] [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)
- [ ] [542. 01 矩阵](https://leetcode-cn.com/problems/01-matrix/submissions/)
- [ ] [934. 最短的桥](https://leetcode-cn.com/problems/shortest-bridge/)
- [ ] [944. 腐烂的橘子](https://leetcode-cn.com/problems/rotting-oranges/)
