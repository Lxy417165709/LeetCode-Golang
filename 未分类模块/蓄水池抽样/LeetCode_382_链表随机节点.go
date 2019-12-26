package 蓄水池抽样

import "math/rand"

type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	return this.poolAlgorithm(1)[0]
}

// 蓄水池算法，等概选择k个元素入蓄水池
func (this *Solution) poolAlgorithm(k int) []int {
	pool := make([]int, k) // 大小为k的蓄水池
	pickCount := 0         // 表示挑选了几个元素
	for cur := this.head; cur != nil; cur = cur.Next {
		random := rand.Intn(pickCount + 1)
		if random < k {
			pool[random] = cur.Val
		}
		pickCount ++
	}
	return pool
}

/*
	题目链接:
		https://leetcode-cn.com/problems/linked-list-random-node/submissions/		链表随机节点
*/
