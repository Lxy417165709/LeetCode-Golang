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

// removeDuplicates2 原地删除排序数组的重复项。 (统一写法)
func removeDuplicates2(nums []int) int {
	return removeDuplicateElements(nums, 1)
}
