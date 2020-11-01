package 二叉树

import "strconv"

var board [][]string

func printTree(root *TreeNode) [][]string {
	board = getBlankBoard(root)
	boardWidth,boardHeight := len(board[0]),len(board)
	beginXCoordinate, beginYCoordinate := boardWidth/2, 0
	distanceOfNextXCoordinateToItsSymmetry := 1 << (uint8(boardHeight - 2))
	formBoard(root, beginXCoordinate, beginYCoordinate, distanceOfNextXCoordinateToItsSymmetry)
	return board
}

func getBlankBoard(root *TreeNode) [][]string {
	width, height := getBoardWidthAndHeight(root)
	board := make([][]string, height)
	for i := 0; i < len(board); i++ {
		board[i] = make([]string, width)
	}
	return board
}

func formBoard(root *TreeNode, xCoordinate, yCoordinate, distanceOfNextXCoordinateToItsSymmetry int) {
	if root == nil {
		return
	}
	board[yCoordinate][xCoordinate] = strconv.Itoa(root.Val)
	nextXCoordinateSymmetry := xCoordinate
	formBoard(
		root.Left,
		nextXCoordinateSymmetry-distanceOfNextXCoordinateToItsSymmetry,
		yCoordinate+1,
		distanceOfNextXCoordinateToItsSymmetry/2,
	)
	formBoard(
		root.Right,
		nextXCoordinateSymmetry+distanceOfNextXCoordinateToItsSymmetry,
		yCoordinate+1,
		distanceOfNextXCoordinateToItsSymmetry/2,
	)
}

func getBoardWidthAndHeight(root *TreeNode) (int, int) {
	treeHeight := getHeight(root)
	return (1 << uint8(treeHeight)) - 1, treeHeight
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接: https://leetcode-cn.com/problems/print-binary-tree/
	总结:
		1. 这题难在找规律。
*/
