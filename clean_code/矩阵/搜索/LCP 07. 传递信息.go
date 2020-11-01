package 搜索

// -------------------------------- 搜索 --------------------------------
func numWays(n int, relation [][]int, k int) int {
	canReach := get2DSlice(n, n)
	for i := 0; i < len(relation); i++ {
		canReach[relation[i][0]][relation[i][1]] = true
	}
	return getNumWays(n, canReach, 0, k)
}

func getNumWays(n int, canReach [][]bool, curNum, k int) int {
	if k == 0 {
		if curNum == n-1 {
			return 1
		}
		return 0
	}
	ways := 0
	for i := 0; i < n; i++ {
		if i == curNum {
			continue
		}
		if canReach[curNum][i] == true {
			ways += getNumWays(n, canReach, i, k-1)
		}
	}
	return ways
}

func get2DSlice(rows, column int) [][]bool {
	slice := make([][]bool, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]bool, column)
	}
	return slice
}

// -------------------------------- 动态规划 --------------------------------
func numWays(n int, relation [][]int, k int) int {
	ways := get2DIntSlice(k+1, n) // ways[i][t], 表示第 i 轮，传递到编号为 t 的方案数。
	ways[0][0] = 1
	for i := 1; i <= k; i++ {
		for t := 0; t < len(relation); t++ {
			ways[i][relation[t][1]] += ways[i-1][relation[t][0]]
		}
	}
	return ways[k][n-1]
}

func get2DIntSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

// -------------------------------- 动态规划状态压缩 --------------------------------
// 正确的状态压缩
func numWays(n int, relation [][]int, k int) int {
	ways := make([]int, n)
	ways[0] = 1
	for i := 1; i <= k; i++ {
		nextLayWays := make([]int, n)
		for t := 0; t < len(relation); t++ {
			nextLayWays[relation[t][1]] += ways[relation[t][0]]
		}
		ways = nextLayWays
	}
	return ways[n-1]
}

// 错误的状态压缩
func numWays(n int, relation [][]int, k int) int {
	ways := make([]int, n)
	ways[0] = 1
	for i := 1; i <= k; i++ {
		for t := 0; t < len(relation); t++ {
			ways[relation[t][1]] += ways[relation[t][0]]
		}
	}
	return ways[n-1]
}

/*
	题目链接: https://leetcode-cn.com/problems/chuan-di-xin-xi/
	总结:
		1. 这题我采用邻接矩阵存储边的信息。
		2. 这题官方还给出了动态规划的解法！
		3. 状态压缩时，要记得保留上层的状态，而且不要让从当层获取当层的结果。 (看上面错误的状态压缩)

*/
