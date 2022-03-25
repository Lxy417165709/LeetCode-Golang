package 二叉树

import "math"

// ---------------------------------------------- 朴素解法 ----------------------------------------------

// rightSideView 获取二叉树右视图。
func rightSideView(root *TreeNode) []int {
	// 1. 形成层级到层信息的映射。
	levelToInfoMap := make(map[int]*LevelInfo)
	formLevelToInfoMap(levelToInfoMap, root, 1)

	// 2. 形成结果集。
	rightSizeView := make([]int, 0)
	for i := 0; i < 100 && levelToInfoMap[i] != nil; i++ {
		rightSizeView = append(rightSizeView, levelToInfoMap[i].Node.Val)
	}

	// 3. 返回。
	return rightSizeView
}

// formLevelToInfoMap 形成层级到层信息的映射。
func formLevelToInfoMap(levelToInfoMap map[int]*LevelInfo, node *TreeNode, curWeight int) {
	// 1. 空返回。
	if node == nil {
		return
	}

	// 2. 获取层级、根据当前节点权重更新层信息。 (同层层级靠右的大，不同层层级高的大)
	level := int(math.Log2(float64(curWeight)))
	info := levelToInfoMap[level]
	if info == nil || curWeight > info.Weight {
		levelToInfoMap[level] = &LevelInfo{
			Node:   node,
			Weight: curWeight,
		}
	}

	// 3. 递归。
	formLevelToInfoMap(levelToInfoMap, node.Left, curWeight*2)
	formLevelToInfoMap(levelToInfoMap, node.Right, curWeight*2+1)
}

// LevelInfo 层信息。
type LevelInfo struct {
	Node   *TreeNode
	Weight int
}

// ---------------------------------------------- 朴素解法 ----------------------------------------------

// ---------------------------------------------- 遍历解法 ----------------------------------------------

// rightSideView 获取二叉树右视图。
func rightSideView(root *TreeNode) []int {
	return formRightSizeView(root, nil, 0)
}

// formRightSizeView 采用 中->右->左 遍历，保证同层遍历时，优先命中最右的节点。
func formRightSizeView(root *TreeNode, curRightSizeView []int, level int) []int {
	// 1. 空返回。
	if root == nil {
		return curRightSizeView
	}

	// 2. 命中本层的第一个节点。 (该节点为同层的最右节点)
	if len(curRightSizeView) == level {
		curRightSizeView = append(curRightSizeView, root.Val)
	}

	// 3. 递归。
	curRightSizeView = formRightSizeView(root.Right, curRightSizeView, level+1)
	curRightSizeView = formRightSizeView(root.Left, curRightSizeView, level+1)

	// 4. 返回。
	return curRightSizeView
}

// ---------------------------------------------- 遍历解法 ----------------------------------------------
