package 二叉树

func numTrees(n int) int {
	countOfBST := make([]int, n+1)
	countOfBST[0] = 1
	for countOfNode := 1; countOfNode <= n; countOfNode++ {
		for countOfNodeInLeftTree := 0; countOfNodeInLeftTree <= countOfNode-1; countOfNodeInLeftTree++ {
			countOfNodeInRightTree := countOfNode - countOfNodeInLeftTree - 1
			countOfBST[countOfNode] += countOfBST[countOfNodeInRightTree] * countOfBST[countOfNodeInLeftTree]
		}
	}
	return countOfBST[n]
}

/*
	题目链接: https://leetcode-cn.com/problems/unique-binary-search-trees/
	总结
		1. 这题偏向于动态规划。
*/
