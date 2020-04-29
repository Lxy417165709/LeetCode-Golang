package main

import "fmt"

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */



const INF = 1000000000

func findInMountainArray(target int, mountainArr *MountainArray) int {
	ans := solve(target, 0, mountainArr.length()-1, mountainArr)
	if ans == INF {
		return -1
	}
	return ans
}

func solve(target int, l, r int, mountainArr *MountainArray) int {
	// [l, r]区间 只有一个元素时
	if l == r && mountainArr.get(l) == target {
		return l
	}

	// [l, r]区间 有多个元素时
	if l < r {
		midIndex := (l + r) / 2
		midVal := mountainArr.get(midIndex)

		if target > midVal {
			if midIndex == 0 || midVal > mountainArr.get(midIndex-1) {
				// [2 4 6 8 10 9 7 6 2]
				// 例如此时 midVal == 8(midIndex==3)  target == 10
				// 则此时 midVal > mountainArr.get(midIndex - 1)
				// 显然我只能向右边找，因为此时左边的数都小于 target
				return solve(target, midIndex+1, r, mountainArr)
			}

			// [2 4 6 8 10 9 7 6 2]
			// 假如 midVal == 7,  target == 10
			// 则此时 midVal < mountainArr.get(midIndex - 1)
			// 显然我只能向左边找，因为此时右边的数都小于 target
			return solve(target, l, midIndex-1, mountainArr)
		}

		// 优先向左边寻找，找到直接返回，因为 左边索引 < 右边索引
		// 注意区间范围
		idx := solve(target, l, midIndex, mountainArr)
		if idx != INF {
			return idx
		}
		return solve(target, midIndex+1, r, mountainArr)
	}
	return INF
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
