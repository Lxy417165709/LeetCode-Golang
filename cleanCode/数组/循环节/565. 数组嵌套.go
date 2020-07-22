package 循环节

var hasIndexVisit []bool

func arrayNesting(nums []int) int {
	hasIndexVisit = make([]bool, len(nums))
	maxLoopSize := 0
	for i := 0; i < len(nums); i++ {
		if hasIndexVisit[i] {
			continue
		}
		maxLoopSize = max(maxLoopSize, getLoopSizeAndMarkHasVisit(nums, i))
	}
	return maxLoopSize

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getLoopSizeAndMarkHasVisit(nums []int, begin int) int {
	hasIndexVisit[begin] = true
	slow := begin
	fast := begin
	loopSize := 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		loopSize++
		hasIndexVisit[slow] = true
		hasIndexVisit[fast] = true
		if slow == fast {
			return loopSize
		}
	}
}

/*
	题目链接:
		1. 这题还能进行优化，就是将数组元素置为负数，表示已访问过，这样能省下一些时间和空间。
			(空间复杂度优化为O(1))
*/
