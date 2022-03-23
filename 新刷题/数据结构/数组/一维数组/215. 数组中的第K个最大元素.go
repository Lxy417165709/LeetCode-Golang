package 一维数组

import (
	"fmt"

	"github.com/Lxy417165709/LeetCode-Golang/新刷题/util/sort_util"
)

// findKthSmallest 找出第k大。
// 注意: 当数组有重复的元素，此时存在相同的数有不同的次序。
func findKthLargest(nums []int, k int) int {
	return findKthSmallest(nums, len(nums)-k+1)
}

// findKthSmallest 找出第k小。
// 注意: 当数组有重复的元素，此时存在相同的数有不同的次序。
func findKthSmallest(nums []int, k int) int {
	// 1. 非法情况。
	if k > len(nums) {
		panic(fmt.Sprintf("长度为[%+v]的数组，不存在第[%+v]小的数。", len(nums), k))
	}

	// 2. 分区，获取基准值次序。
	index := sort_util.Partition(nums)
	curOrder := index + 1

	// 3. 递归。
	if k == curOrder {
		// 3.1 如果k等于基准值次序，此时返回。
		return nums[index]
	} else if k > curOrder {
		// 3.2 如果k大于基准值次序，说明要向数组基准值索引右边寻找。
		return findKthSmallest(nums[index+1:], k-curOrder)
	} else {
		// 3.3 如果k小于基准值次序，说明要向数组基准值索引左边寻找。
		return findKthSmallest(nums[:index], k)
	}
}

// 注意: 这个题也可以用小顶堆实现。
