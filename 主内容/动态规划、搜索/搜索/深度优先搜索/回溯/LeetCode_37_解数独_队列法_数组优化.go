package main

import "fmt"

var rows, cols, blocks [9][10]bool	// 是一个映射，用于存放某行、某列、某块的数字使用情况

type Node struct {
	x, y int
}

// 函数入口
func solveSudoku(board [][]byte) {
	rows = [9][10]bool{}
	cols = [9][10]bool{}
	blocks = [9][10]bool{}

	queue := make([]Node, 0)
	for i := 0; i < len(board); i++ {
		for t := 0; t < len(board[i]); t++ {
			char := board[i][t]
			if char == '.' {
				queue = append(queue, Node{i, t})
			} else {
				mark(i, t, char, true)
			}
		}
	}
	solveSudokuExec(board, queue)
}
// 递归函数
func solveSudokuExec(board [][]byte, queue []Node) bool {
	// 没有填充的方格时，说明已经填充完毕了
	if len(queue) == 0 {
		return true
	}

	// 有要填充的方格时，尝试在该方格填充 '1'-'9'
	for num := uint8(1 + '0'); num <= uint8(9+'0'); num++ {
		x, y := queue[0].x, queue[0].y
		// 判断该数字是否可以填充
		if isValid(x, y, num) == false {
			continue
		}
		mark(x, y, num, true)
		board[x][y] = num
		if solveSudokuExec(board, queue[1:]) == true {
			return true
		}
		// 清除标记 -> 回溯的特点
		board[x][y] = '.'
		mark(x, y, num, false)
	}
	return false
}


// 标记该行、该列、该块，数字是否已使用 (我只是为了方便才把这个函数提取出来)
func mark(x, y int, char uint8, flag bool) {
	rows[x][char-'0'] = flag
	cols[y][char-'0'] = flag
	blocks[x/3*3+y/3][char-'0'] = flag
}

// 判断该行、该列、该块，该数字是否可以使用 (我只是为了方便才把这个函数提取出来)
func isValid(x, y int, char uint8) bool {
	// 判断该数字在该行是否已使用
	if rows[x][char-'0'] == true {
		return false
	}
	// 判断该数字在该列是否已使用
	if cols[y][char-'0'] == true {
		return false
	}
	// 判断该数字在该块是否已使用
	if blocks[x/3*3+y/3][char-'0'] == true {
		return false
	}
	return true
}


/*
	总结
	1. 在这个版本中，主要的修改是: 我把map换成了bool数组，加快查找速度。
	2. 回溯的特点就是: 处理完下一层的递归后，返回该层时，需要把下一层的递归的信息清空。
*/
