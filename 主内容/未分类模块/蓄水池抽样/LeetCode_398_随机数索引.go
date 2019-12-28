package 蓄水池抽样

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Pick(target int) int {
	pool := this.poolAlgorithm(1, target)
	return pool[0]
}

// 蓄水池算法，等概选择k个元素索引入蓄水池。 (由于题目限定，此时选择的元素的值是等于target的)
func (this *Solution) poolAlgorithm(k int, target int) []int {
	pool := make([]int, k) // 大小为k的蓄水池
	pickCount := 0         // 表示挑选了几个元素
	for i := 0; i < len(this.nums); i++ {
		if this.nums[i] != target {
			continue
		}
		random := rand.Intn(pickCount + 1)
		if random < k {
			pool[random] = i
		}
		pickCount ++
	}
	return pool
}

/*
	题目链接:
		https://leetcode-cn.com/problems/random-pick-index/submissions/		随机数索引
*/
