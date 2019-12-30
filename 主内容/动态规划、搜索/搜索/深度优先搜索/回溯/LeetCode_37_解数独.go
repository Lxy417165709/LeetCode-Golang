package main

import "fmt"

var rows, cols, blocks []map[uint8]bool

func solveSudoku(board [][]byte) {
	rows = make([]map[uint8]bool, 9)
	cols = make([]map[uint8]bool, 9)
	blocks = make([]map[uint8]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[uint8]bool)
		cols[i] = make(map[uint8]bool)
		blocks[i] = make(map[uint8]bool)
	}

	release := 0
	for i := 0; i < len(board); i++ {
		for t := 0; t < len(board[i]); t++ {
			char := board[i][t]
			mark(i, t, char, true)
			if char == '.' {
				release++
			}
		}
	}
	solveSudokuExec(board, release)
	for i := 0; i < 9; i++ {
		for t := 0; t < 9; t++ {
			fmt.Print(string(board[i][t]), " ")
		}
		fmt.Println()
	}
}

// 标记
func mark(i, t int, char uint8, flag bool) {
	rows[i][char] = flag
	cols[t][char] = flag
	blocks[i/3*3+t/3][char] = flag
}

//
func solveSudokuExec(board [][]byte, release int) bool {
	if release == 0 {
		return true
	}
	x,y := 0,0
	beginInsert:
	for i := 0; i < len(board); i++ {
		for t := 0; t < len(board[i]); t++ {
			if board[i][t] == '.' {
				x,y  = i,t
				break beginInsert
			}
		}
	}
	for num := uint8(1 + '0'); num <= uint8(9+'0'); num++ {
		if rows[x][num] == true || cols[y][num] == true || blocks[x/3*3+y/3][num] == true {
			continue
		}
		char := num
		mark(x, y, char, true)

		board[x][y] = num
		if solveSudokuExec(board, release-1) == true {
			return true
		}
		board[x][y] = '.'
		mark(x, y, char, false)
	}
	return false
}

/*
	总结
	1. 理清思路再暴力，这题就不难啦~
*/
