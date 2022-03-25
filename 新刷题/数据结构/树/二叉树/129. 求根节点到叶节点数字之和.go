package 二叉树

import (
	"fmt"
	"strconv"
)

// ----------------------------------------------- 朴素法 -----------------------------------------------

// sumNumbers 求根节点到叶节点所有路径形成的数值之和。
// 思路: 获取所有数字后相加。
func sumNumbers(root *TreeNode) int {
	// 1. 获取树中的所有数。
	numberStrings := getNumberStrings(root)

	// 2. 计算结果。
	sum := 0
	for _, numberString := range numberStrings {
		num, _ := strconv.Atoi(numberString)
		sum += num
	}

	// 3. 返回。
	return sum
}

// getNumberStrings 获取数字字符串。
func getNumberStrings(root *TreeNode) []string {
	// 1. 空返回。
	if root == nil {
		return nil
	}

	// 2. 递归边界
	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}

	// 3. 递归。
	result := make([]string, 0)
	leftNumbers := getNumberStrings(root.Left)
	rightNumbers := getNumberStrings(root.Right)
	sonNumbers := append(leftNumbers, rightNumbers...)
	for _, number := range sonNumbers {
		result = append(result, fmt.Sprintf("%d", root.Val)+number)
	}

	// 4. 返回。
	return result
}

// ----------------------------------------------- 朴素法 -----------------------------------------------

// ----------------------------------------------- 精妙法 -----------------------------------------------

// sumNumbers 求根节点到叶节点所有路径形成的数值之和。
func sumNumbers(root *TreeNode) int {
	return getPathSum(root, 0)
}

// getPathSum 求根节点到叶节点所有路径形成的数值之和。
// preSum 为 根节点->本节点父节点 数字序列对应的数值 * 10。
// preSum % 10 一定为 0。
func getPathSum(root *TreeNode, preSum int) int {
	// 1. 空返回。
	if root == nil {
		return 0
	}

	// 2. 叶子节点，返回。
	if root.Left == nil && root.Right == nil {
		return preSum + root.Val
	}

	// 3. 递归。
	return getPathSum(root.Left, (preSum+root.Val)*10) + getPathSum(root.Right, (preSum+root.Val)*10)
}

// ----------------------------------------------- 精妙法 -----------------------------------------------
