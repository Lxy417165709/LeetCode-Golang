package 正确位

func findErrorNums(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for !isInRightPosition(nums, i) {
			if getRightPositionNum(nums, i) == nums[i] {
				break
			}
			toRightPosition(nums, i)
		}
	}

	for i := 0; i < len(nums); i++ {
		if !isInRightPosition(nums, i) {
			return []int{nums[i], i + 1}
		}
	}
	panic("程序出错")

}

func isInRightPosition(nums []int, index int) bool {
	return nums[index] == index+1
}

func getRightPositionNum(nums []int, index int) int {
	return nums[getRightPosition(nums, index)]
}

func toRightPosition(nums []int, index int) {
	rightPosition := getRightPosition(nums, index)
	nums[rightPosition], nums[index] = nums[index], nums[rightPosition]
}

func getRightPosition(nums []int, index int) int {
	return nums[index] - 1
}

/*
	题目链接: https://leetcode-cn.com/problems/set-mismatch/
*/
