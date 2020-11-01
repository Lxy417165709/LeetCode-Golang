package 获取

var minLabelsInLay []int
var maxLabelsInLay []int
var maxLayOfOwningLabel int

func widthOfBinaryTree(root *TreeNode) int {
	minLabelsInLay = make([]int, 0)
	maxLabelsInLay = make([]int, 0)
	maxLayOfOwningLabel = -1
	formLabels(root, 0, 0)
	return getMaxWidth()
}

func formLabels(root *TreeNode, nowLay int, nowLabel int) {
	if root == nil {
		return
	}
	if nowLay == maxLayOfOwningLabel+1 {
		minLabelsInLay = append(minLabelsInLay, nowLabel)
		maxLabelsInLay = append(maxLabelsInLay, nowLabel)
		maxLayOfOwningLabel++
	}
	minLabelsInLay[nowLay] = min(minLabelsInLay[nowLay], nowLabel)
	maxLabelsInLay[nowLay] = max(maxLabelsInLay[nowLay], nowLabel)

	formLabels(root.Left, nowLay+1, nowLabel*2+1)
	formLabels(root.Right, nowLay+1, nowLabel*2+2)
}

func getMaxWidth() int {
	maxWidth := 0
	for i := 0; i <= maxLayOfOwningLabel; i++ {
		maxWidth = max(maxLabelsInLay[i]-minLabelsInLay[i]+1, maxWidth)
	}
	return maxWidth
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接: https://leetcode-cn.com/problems/maximum-width-of-binary-tree/
	总结:
		1. 为了结构清晰，我创建了一个maxLabel变量 (其实可以省略这个变量，而在formLabels时就求取出maxWidth)。
			同样的，maxLayOfOwningLabel也可以省略的。
		2. formLabels的nowLay参数其实可以省略，因为它可以通过nowLabel求出。 (通过取对数运算)
*/
