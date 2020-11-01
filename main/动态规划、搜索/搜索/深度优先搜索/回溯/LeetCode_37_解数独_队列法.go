package main

import "fmt"

var rows,cols,blocks []map[uint8]bool
type Node struct{
	x,y int
}
func solveSudoku(board [][]byte)  {
	rows = make([]map[uint8]bool,9)
	cols = make([]map[uint8]bool,9)
	blocks = make([]map[uint8]bool,9)

	for i:=0;i<9;i++{
		rows[i] = make(map[uint8]bool)
		cols[i] = make(map[uint8]bool)
		blocks[i] = make(map[uint8]bool)
	}
	queue := make([]Node,0)
	for i:=0;i<len(board);i++{
		for t:=0;t<len(board[i]);t++{
			char := board[i][t]
			mark(i,t,char,true)
			if char=='.'{
				queue = append(queue,Node{i,t})
			}
		}
	}
	solveSudokuExec(board,queue)
	// for i:=0;i<9;i++{
	//     for t:=0;t<9;t++{
	//         fmt.Print(string(board[i][t])," ")
	//     }
	//     fmt.Println()
	// }
}
func mark(i,t int,char uint8,flag bool){
	rows[i][char]=flag
	cols[t][char]=flag
	blocks[i/3*3 +t/3][char]=flag
}

func solveSudokuExec(board [][]byte,queue []Node) bool{
	if len(queue)==0{
		return true
	}

	for num := uint8(1+'0');num<=uint8(9+'0');num++{
		x,y := queue[0].x,queue[0].y
		if rows[x][num] == true || cols[y][num] == true || blocks[x/3*3+y/3][num] == true {
			continue
		}
		char := num
		mark(x, y, char, true)

		board[x][y] = num
		if solveSudokuExec(board, queue[1:]) == true {
			return true
		}
		board[x][y] = '.'
		mark(x, y, char, false)
	}
	return false
}


/*
	总结
	1. 这个版本我把需要填充的坐标加入队列中，之后再遍历这个队列，这样就可以只使用O(1)的时间找到下一个需要填充的坐标了。
*/
