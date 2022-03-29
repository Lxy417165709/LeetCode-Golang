package 一维数组

// maxSlidingWindow 获取数组在每个窗口的最大值。滑动窗口大小为k，每次向右移动一格。
func maxSlidingWindow(nums []int, k int) []int {
	// 1. 初始化。
	indexQueue := make([]int, 0) // 双端单调递减队列，存放元素索引。
	result := make([]int, 0)     // 结果集。

	// 2. 获取结果集。
	for i := 0; i < len(nums); i++ {
		// 2.1 窗口容积优先，如果窗口已满，弹出最先加入窗口的元素。
		// (不能使用 k == len(indexQueue) 判断，因为为了维护单调递减性，队列内的元素个数肯定小于等于k)
		// (队列内的最大索引在队列尾，最小在队列头，窗口当前容纳的元素个数 = 最大索引-最小索引 + 1)
		if len(indexQueue) != 0 && indexQueue[len(indexQueue)-1]-indexQueue[0]+1 == k {
			indexQueue = indexQueue[1:]
		}

		// 2.2 维护单调递减队列。
		for len(indexQueue) != 0 && nums[indexQueue[len(indexQueue)-1]] <= nums[i] {
			indexQueue = indexQueue[:len(indexQueue)-1]
		}
		indexQueue = append(indexQueue, i)

		// 2.3 记录结果集。
		if i >= k-1 {
			result = append(result, nums[indexQueue[0]])
		}
	}

	// 3. 返回。
	return result
}
