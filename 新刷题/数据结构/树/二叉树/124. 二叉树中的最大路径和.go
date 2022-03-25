package 二叉树

import (
	"fmt"

	"github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"
)

// INF 无穷大。
const INF = 100000000

// cache 缓存。
var cache = make(map[string]int)

// maxPathSum 可能具有根节点的最大路径值。 (该函数其实不用缓存，因此不会重复访问)
func maxPathSum(root *TreeNode) int {
	// 1. 空返回。
	if root == nil {
		return -INF
	}

	// 2. 读取缓存。
	cacheKey := fmt.Sprintf("%s:%d", "maxPathSum", root)
	if value, ok := cache[cacheKey]; ok {
		return value
	}

	// 3. 递归。
	leftMaxPathSumWithoutRoot := maxPathSum(root.Left)
	rightMaxPathSumWithoutRoot := maxPathSum(root.Right)
	leftMaxPathSumWithRoot := root.Val + maxPathSumWithRoot(root.Left)
	rightMaxPathSumWithRoot := root.Val + maxPathSumWithRoot(root.Right)
	leftRightMaxPathSumWithRoot := leftMaxPathSumWithRoot + rightMaxPathSumWithRoot - root.Val

	// 4. 写缓存。
	cache[cacheKey] = math_util.Max(
		root.Val,
		leftMaxPathSumWithoutRoot, rightMaxPathSumWithoutRoot,
		leftMaxPathSumWithRoot, rightMaxPathSumWithRoot, leftRightMaxPathSumWithRoot,
	)

	// 5. 返回。
	return cache[cacheKey]
}

// maxPathSumWithRoot 具有根节点的最大路径值。 (该函数需要缓存，因为 maxPathSum 会调用 maxPathSum、maxPathSumWithRoot)
func maxPathSumWithRoot(root *TreeNode) int {
	// 1. 空返回。
	if root == nil {
		return -INF
	}

	// 2. 读取缓存。
	cacheKey := fmt.Sprintf("%s:%d", "maxPathSumWithRoot", root)
	if value, ok := cache[cacheKey]; ok {
		return value
	}

	// 3. 递归，写缓存。
	cache[cacheKey] = math_util.Max(maxPathSumWithRoot(root.Left), maxPathSumWithRoot(root.Right), 0) + root.Val

	// 4. 返回。
	return cache[cacheKey]
}
