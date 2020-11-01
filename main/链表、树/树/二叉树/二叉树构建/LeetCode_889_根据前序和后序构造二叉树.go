package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 通过先序+后序构造二叉树 (不唯一)
func constructFromPrePost(pre []int, post []int) *TreeNode {
	return constructFromPrePostExec(pre, post)
}
func constructFromPrePostExec(pre []int, post []int) *TreeNode {
	if len(pre) == 0 || len(post) == 0 {
		return nil
	}
	if len(pre) == 1 || len(post) == 1 {
		return &TreeNode{pre[0], nil, nil}
	}
	rootVal := pre[0]
	index := 0
	for post[index] != pre[1] {
		index++
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  constructFromPrePostExec(pre[1:index+2], post[:index+1]),
		Right: constructFromPrePostExec(pre[index+2:], post[index+1:len(post)-1]),
	}
}

/*
	总结
	1. 对于遍历序列来说，如果序列输出包括空节点，那么这个序列就能唯一的确定一棵树，而如果不包括
		空节点，那么就不能唯一确认。
	2. 以上前序，后序序列不能唯一的确认一棵树，比如下面的情况
			1              1
		     \            /
		      2          2
               \        /
				3      3
			   a         b

	    a树的前序、后序序列  和  b树的前序、后序序列  是一样的，而如果输出空节点，那就是不一样的。
*/
