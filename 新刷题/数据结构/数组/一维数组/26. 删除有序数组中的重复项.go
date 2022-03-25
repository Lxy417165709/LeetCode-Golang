package 一维数组

// removeDuplicates 原地删除排序数组的重复项。
func removeDuplicates(nums []int) int {
	writeIndex, readIndex := 0, 0
	for readIndex < len(nums) {
		for readIndex+1 < len(nums) && nums[readIndex+1] == nums[readIndex] {
			readIndex++
		}
		nums[writeIndex] = nums[readIndex]
		readIndex++
		writeIndex++
	}
	return writeIndex
}
