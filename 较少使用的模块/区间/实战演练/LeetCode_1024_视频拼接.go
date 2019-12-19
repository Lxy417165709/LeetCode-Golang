package main

import "sort"

func videoStitching(clips [][]int, T int) int {
	if len(clips) == 0 {
		return -1
	}
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		} else {
			return clips[i][0] < clips[j][0]
		}
	})
	left, right := clips[0][0], clips[0][1]
	if left != 0 {
		return -1
	}
	index := 0 // 表示接下来要遍历的索引
	count := 1 // 表示最少视频数
	for right < T {
		nextRight := right
		// 这里表示以right为右边界时，可以找到下一个边界。
		for index < len(clips) && clips[index][0] <= right {
			nextRight = max(nextRight, clips[index][1])
			index++
		}
		// 此时表示无法拓展
		if nextRight == right {
			return -1
		}
		right = nextRight
		count++
	}
	return count
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接:
		https://leetcode-cn.com/problems/video-stitching/submissions/		视频拼接
*/

/*
	总结
	1. 这道题的抽象模型就是: 给你一些区间，让你使用最少的区间个数，组成一个连续的目标区间。
*/
